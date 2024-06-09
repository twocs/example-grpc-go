package main

import (
	"context"
	"flag"
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	pb "github.com/twocs/example-grpc-go/example-grpc-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr  = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
	id    = flag.Int("id", 3, "The user ID")
	ids   = flag.String("ids", "1,3", "The user IDs, separated by commas")
	query = flag.String("query", "Bob", "The search query")
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

	// Contact the server and print out its response (Get single user)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get a single user
	getUser(ctx, c, id)

	// Contact the server and print out its response (Get multiple users)
	listUsers(ctx, c, ids)

	// Contact the server and print out its response (Search for users by first name)
	searchUsers(ctx, c, *query)
}

// getUser gets a single user from the server.
func getUser(ctx context.Context, c pb.GetUsersClient, id *int) {
	log.Println("GetUser")
	r1, err := c.GetUser(ctx, &pb.ID{Id: int32(*id)})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}

	log.Println(r1)
}

// listUsers gets a list of users from the server.
func listUsers(ctx context.Context, c pb.GetUsersClient, ids []int32) {
	log.Println("ListUsers")
	r2, err := c.ListUsers(ctx, &pb.IDs{Id: ids})
	if err != nil {
		log.Fatalf("could not list users: %v", err)
	}

	for {
		user, err := r2.Recv()
		if err == io.EOF {
			// read done
			break
		}

		if err != nil {
			log.Fatalf("error in ListUsers - %v", err)
		}

		log.Println(user)
	}
}

// searchUsers searches for users by first name.
func searchUsers(ctx context.Context, c pb.GetUsersClient, query string) {
	log.Println("SearchUsers")
	r3, err := c.SearchUsers(ctx, &pb.Query{Query: query})
	if err != nil {
		log.Fatalf("could not search users: %v", err)
	}

	for {
		user, err := r3.Recv()
		if err == io.EOF {
			// read done
			break
		}

		if err != nil {
			log.Fatalf("error in SearchUsers - %v", err)
		}

		log.Println(user)
	}
}
