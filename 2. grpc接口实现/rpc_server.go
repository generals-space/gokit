package main 

import (
	"google.golang.org/grpc"
	"context"
)
// UserManagerServiceServer ...
type UserManagerServiceServer struct{}

// GetUser ...
func (server *UserManagerServiceServer)GetUser(ctx context.Context, req GetUserRequest)(res GetUserResponse, err error){
	return
}
func main(){
	
}