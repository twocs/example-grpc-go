package main

import (
	"context"
	"flag"
	"log"
	"strconv"
	"strings"
	"time"

	pb "github.com/twocs/example-grpc-go/example-grpc-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
	id   = flag.Int("id", 3, "The user ID")
	ids  = flag.String("ids", "1,3", "The user IDs, separated by commas")
)

func main() {
	flag.Parse()

	// Convert the comma-separated list of IDs to a slice of integers
	idStrings := strings.Split(*ids, ",")
	ids := make([]int32, len(idStrings))
	for i, idString := range idStrings {
		id, err := strconv.Atoi(idString)
		if err != nil {
			log.Fatalf("could not convert %q to int: %v", idString, err)
			return
		}
		ids[i] = int32(id)
	}

	// set up a connection to the server
	// TODO: set up secure credentials
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	c := pb.NewGetUsersClient(conn)

	// Contact the sever and print out its response (Get single user)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := c.GetUser(ctx, &pb.ID{Id: int32(*id)})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}

	log.Printf("User: %+v", r)

	// TODO: Contact the server and print out its response (Get multiple users)
}
