package models

import (
  "time"
  // "fmt"
  // "log"
  "github.com/satori/go.uuid"
  pb "user-api/rpc/track"
  "github.com/go-pg/pg"
  "user-api/internal"
)


type Track struct {
  Id uuid.UUID `sql:"type:uuid,default:uuid_generate_v4()"`
  CreatedAt time.Time `sql:"default:now()"`
	UpdatedAt time.Time

  Title string `sql:",notnull"`
  Status string  `sql:"type:track_status,notnull"`
  Enabled bool `sql:",notnull"`
  TrackNumber int32 `sql:",notnull"`
  Duration float32

  TrackGroups []uuid.UUID `sql:",type:uuid[]" pg:",array"`
  FavoriteOfUsers []uuid.UUID `sql:",type:uuid[]" pg:",array"`

  TrackServerId uuid.UUID `sql:"type:uuid,notnull"`

  CreatorId uuid.UUID `sql:"type:uuid,notnull"`
  Creator *User

  UserGroupId uuid.UUID `sql:"type:uuid,notnull"`
  UserGroup *UserGroup // track belongs to user group (the one who gets paid)

  Artists []uuid.UUID `sql:",type:uuid[]" pg:",array"` // for display purposes
  Tags []uuid.UUID `sql:",type:uuid[]" pg:",array"`

  // Composers with IPI
  // Performers with IPI
}

func (t *Track) Update(db *pg.DB, track *pb.Track) (error, string) {
  // Update tags? artists?
  // tracks can be added to a track group from dedicated endpoint
  // in TrackGroup Service AddTracksToTrackGroup
  var table string
  tx, err := db.Begin()
  if err != nil {
    return err, table
  }
  defer tx.Rollback()

  err, table = t.GetIds(track)
  if err != nil {
    return err, table
  }

  t.UpdatedAt = time.Now()
  _, pgerr := tx.Model(t).
    Column("title", "updated_at", "status", "track_number", "duration", "track_server_id").
    WherePK().
    Returning("*").
    Update()
  if pgerr != nil {
    return pgerr, "track"
  }

  return tx.Commit(), table
}

func (t *Track) Create(db *pg.DB, track *pb.Track) (error, string) {
  var table string
  tx, err := db.Begin()
  if err != nil {
    return err, table
  }
  defer tx.Rollback()

  err, table = t.GetIds(track)
  if err != nil {
    return err, table
  }

  artistIds, pgerr := GetRelatedUserGroupIds(track.Artists, tx)
  if pgerr != nil {
    return pgerr, "user_group"
  }

  // Select or create tags
  tagIds, pgerr := GetTagIds(track.Tags, tx)
  if pgerr != nil {
    return pgerr, "tag"
  }

  t.Tags = tagIds
  t.Artists = artistIds

  _, pgerr = tx.Model(t).Returning("*").Insert()

  if pgerr != nil {
    return pgerr, "track"
  }
  track.Id = t.Id.String()

  // Add track to involved user groups
  // userGroupId can be part of artistIds (artist adding his/her track)
  // or not (label adding a track for one or more artists)
  userGroupIds := internal.RemoveDuplicates(append(artistIds, t.UserGroupId))
  trackIdArr := []uuid.UUID{t.Id}
  _, pgerr = tx.Exec(`
    UPDATE user_groups
    SET tracks = (select array_agg(distinct e) from unnest(tracks || ?) e)
    WHERE id IN (?)
  `, pg.Array(trackIdArr), pg.In(userGroupIds))
  if pgerr != nil {
    return pgerr, "user_group"
  }

  return tx.Commit(), table
}

func (t *Track) Delete(db *pg.DB, track *pb.Track) (error, string) {
  // Delete from track server?
  var table string
  tx, err := db.Begin()
  if err != nil {
    return err, table
  }
  defer tx.Rollback()

  pgerr := tx.Model(t).WherePK().Select()
  if pgerr != nil {
    return pgerr, "track"
  }

  // Delete track from user_group/artists tracks array
  userGroupIds := internal.RemoveDuplicates(append(t.Artists, t.UserGroupId))
  _, pgerr = tx.Exec(`
    UPDATE user_groups
    SET tracks = array_remove(tracks, ?)
    WHERE id IN (?)
  `, t.Id, pg.In(userGroupIds))
  if pgerr != nil {
    return pgerr, "user_group"
  }

  // Delete track from track_groups tracks array
  if len(t.TrackGroups) > 0 {
    _, pgerr = tx.Exec(`
      UPDATE track_groups
      SET tracks = array_remove(tracks, ?)
      WHERE id IN (?)
    `, t.Id, pg.In(t.TrackGroups))
    if pgerr != nil {
      return pgerr, "track_group"
    }
  }

  // Delete track from user favorite_tracks array
  if len(t.FavoriteOfUsers) > 0 {
    _, pgerr = tx.Exec(`
      UPDATE users
			SET favorite_tracks = array_remove(favorite_tracks, ?)
			WHERE id IN (?)
    `, t.Id, pg.In(t.FavoriteOfUsers))
    if pgerr != nil {
      return pgerr, "user_group"
    }
  }

  pgerr = tx.Delete(t)
  if pgerr != nil {
    return pgerr, "track"
  }

  return tx.Commit(), table
}

// Get track OwnerId, UserGroupId and TrackServerId as UUID
func (t *Track) GetIds(track *pb.Track) (error, string) {
  creatorId, err := internal.GetUuidFromString(track.CreatorId)
  if err != nil {
    return err, "owner"
  }

  userGroupId, err := internal.GetUuidFromString(track.UserGroupId)
  if err != nil {
    return err, "user_group"
  }

  if track.TrackServerId != "" {
    trackServerId, err := internal.GetUuidFromString(track.TrackServerId)
    if err != nil {
      return err, "track_server"
    }
    t.TrackServerId = trackServerId
  }

  t.UserGroupId = userGroupId
  t.CreatorId = creatorId

  return nil, ""
}
