package main

import (
	"fmt"
	"net/http"
	"github.com/rs/cors"
	userServer "user-api/internal/server/user"
	userGroupServer "user-api/internal/server/usergroup"
	userRpc "user-api/rpc/user"
	userGroupRpc "user-api/rpc/usergroup"
	"user-api/internal/database"
)

func main() {
	fmt.Printf("User Service on :8080")

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	db := database.Connect(false)

	newUserServer := userServer.NewServer(db)
	userTwirpHandler := userRpc.NewUserServiceServer(newUserServer, nil)

	newUserGroupServer := userGroupServer.NewServer(db)
	userGroupTwirpHandler := userGroupRpc.NewUserGroupServiceServer(newUserGroupServer, nil)

	mux := http.NewServeMux()
	mux.Handle(userRpc.UserServicePathPrefix, userTwirpHandler)
	mux.Handle(userGroupRpc.UserGroupServicePathPrefix, userGroupTwirpHandler)

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST).
	handler := cors.Default().Handler(mux)

	http.ListenAndServe(":8080", handler)
	defer db.Close()
}
