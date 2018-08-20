package tagserver

import (
  "context"
  // "fmt"
  "github.com/go-pg/pg"
  "github.com/twitchtv/twirp"
  // "github.com/satori/go.uuid"

  pb "user-api/rpc/tag"
  // "user-api/internal"
  "user-api/internal/database/models"
)

type Server struct {
	db *pg.DB
}

func NewServer(db *pg.DB) *Server {
	return &Server{db: db}
}

func (s *Server) SearchGenres(ctx context.Context, q *pb.Query) (*pb.SearchResults, error) {
  if len(q.Query) < 3 {
    return nil, twirp.InvalidArgumentError("query", "must be a valid search query")
  }
  tags, twerr := models.SearchTags(q.Query, "genre", s.db)
  if twerr != nil {
    return nil, twerr
  }

  // Build query string "tagId1|tagId2|..."
  var tagIds string
  for _, tag := range tags {
    tagIds = tagIds + "|" + tag.Id.String()
  }

  trackGroupSearchResults, twerr := models.SearchTrackGroups(tagIds, s.db)
  if twerr != nil {
    return nil, twerr
  }

  trackSearchResults, twerr := models.SearchTracks(tagIds, s.db)
  if twerr != nil {
    return nil, twerr
  }

  return &pb.SearchResults{
    Playlists: trackGroupSearchResults.Playlists,
    Albums: trackGroupSearchResults.Albums,
    Tracks: trackSearchResults.Tracks,
  }, nil
}
