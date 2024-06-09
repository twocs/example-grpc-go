package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/twocs/example-grpc-go/database"
	pb "github.com/twocs/example-grpc-go/example-grpc-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// Default port is 50051, or the port is specified by the user.
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()

	// listen for incoming requests
	// create a new gRPC server
	// register the user service with the gRPC server
	// start the gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGetUsersServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// server implements the user service.
type server struct {
	pb.UnimplementedGetUsersServer
}

// GetUserByID returns a user based on the user ID.
func (s *server) GetUser(ctx context.Context, req *pb.ID) (*pb.User, error) {
	log.Printf("GetUserByID: %v", req.GetId())
	user, err := database.GetUserByID(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	return (*User)(user).PBUser(), nil
}

type User database.User

func (user *User) PBUser() *pb.User {
	return &pb.User{
		ID:      user.ID,
		Fname:   user.Fname,
		City:    user.City,
		Phone:   user.Phone,
		Height:  user.Height,
		Married: user.Married,
	}
}

// GetUsersByIds returns a list of users based on a list of user IDs.
// If a user ID is not found, it is skipped.
func (s *server) ListUsers(ids *pb.IDs, lus pb.GetUsers_ListUsersServer) error {
	for _, id := range ids.Id {
		user, err := database.GetUserByID(id)
		switch {
		case errors.Is(err, database.ErrUserNotFound):
			continue
		case err != nil:
			return status.Errorf(codes.NotFound, "user not found: %v", err)
		default:
			if err := lus.Send((*User)(user).PBUser()); err != nil {
				return status.Errorf(codes.Internal, "failed to send user: %v", err)
			}
		}
	}
	return nil
}

// TODO:
// Implement the SearchUsers method.
func (s *server) SearchUsers(query *pb.Query, sus pb.GetUsers_SearchUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchUsers not implemented")
}
