package main

import (
	"context"
	"fmt"
	"games-api/cache"
	"games-api/config"
	"games-api/database"
	"games-api/gamefeed/gamefeedpb"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type server struct{}

var srvConnStr = fmt.Sprintf("%s:%d", config.ServerHost, config.ServerPort)

func (*server) GameFeed(ctx context.Context, req *gamefeedpb.GameFeedRequest) (*gamefeedpb.GameFeedResponse, error) {
	//firstName := req.GetGameFeedInput().GetFirstName()
	gameFeedsRaw := database.GetGameFeeds()

	var gameFeeds []*gamefeedpb.GameFeed
	for i := 0; i < len(gameFeedsRaw); i++ {
		gameFeeds = append(gameFeeds, &gameFeedsRaw[i])
	}

	return &gamefeedpb.GameFeedResponse{
		GameFeed: gameFeeds,
	}, nil
}

func startServer(lis net.Listener, out chan bool) {
	defer close(out)

	newServer := grpc.NewServer()
	gamefeedpb.RegisterGameFeedServiceServer(newServer, &server{})

	if err := newServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	out <- true
}

func initServer(){
	lis, err := net.Listen("tcp", srvConnStr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	for {
		channel := make(chan bool)
		startServer(lis, channel)
	}
}

func main() {
	fmt.Println("Hello my master! I'm your slave Ricardo. I live to serve you. Command me and I'll took my life! ... ")

	initServer()
}

func redisUsage(){
	// Redis Connect Usage
	redisdb, err := cache.Connect(config.RedisHost, config.RedisPort)
	if err != nil {
		log.Fatalf("Failed to connect to redis: %s", err.Error())
	}

	// Redis Exists Usage
	isKeyExists, err := redisdb.Exists("test")
	if err != nil{
		log.Fatalf("Failed to check existence of the key in Redis: %s", err.Error())
	}

	// Redis Delete Usage
	if isKeyExists > 0{
		delRes, err := redisdb.Del("test")
		if err != nil{
			log.Fatalf("Failed to delete by from from Redis: %s", err.Error())
		}

		fmt.Println(delRes)
	}

	// Redis Set Usage
	res, err := redisdb.Set("test","12345", 0)
	if err != nil{
		log.Fatalf("Failed to set new key/value in Redis: %s", err.Error())
	}

	// Redis Get Usage
	test, err := redisdb.Get("test")
	if err != nil{
		log.Fatalf("Failed to get by key from Redis: %s", err.Error())
	}

	fmt.Println(res)
	fmt.Println(test)
	os.Exit(1)
}