package main

import (
	"fmt"
	"context"
	"net/http"
	"github.com/rs/cors"
	userServer "user-api/internal/server/user"
	userGroupServer "user-api/internal/server/usergroup"
	trackServer "user-api/internal/server/track"
	trackGroupServer "user-api/internal/server/trackgroup"
	userRpc "user-api/rpc/user"
	userGroupRpc "user-api/rpc/usergroup"
	trackRpc "user-api/rpc/track"
	trackGroupRpc "user-api/rpc/trackgroup"
	"user-api/internal/database"
)

func WithURLQuery(base http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		query := r.URL.Query()
		ctx = context.WithValue(ctx, "query", query)
		r = r.WithContext(ctx)

		base.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Printf("User Service on :8080")

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	db := database.Connect(true)

	newUserServer := userServer.NewServer(db)
	userTwirpHandler := userRpc.NewUserServiceServer(newUserServer, nil)

	newUserGroupServer := userGroupServer.NewServer(db)
	userGroupTwirpHandler := WithURLQuery(userGroupRpc.NewUserGroupServiceServer(newUserGroupServer, nil))

	newTrackServer := trackServer.NewServer(db)
	trackTwirpHandler := trackRpc.NewTrackServiceServer(newTrackServer, nil)

	newTrackGroupServer := trackGroupServer.NewServer(db)
	trackGroupTwirpHandler := trackGroupRpc.NewTrackGroupServiceServer(newTrackGroupServer, nil)

	mux := http.NewServeMux()
	mux.Handle(userRpc.UserServicePathPrefix, userTwirpHandler)
	mux.Handle(userGroupRpc.UserGroupServicePathPrefix, userGroupTwirpHandler)
	mux.Handle(trackRpc.TrackServicePathPrefix, trackTwirpHandler)
	mux.Handle(trackGroupRpc.TrackGroupServicePathPrefix, trackGroupTwirpHandler)

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST).
	handler := cors.Default().Handler(mux)

	http.ListenAndServe(":8080", handler)
	defer db.Close()
}
