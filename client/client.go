package main

import (
	"context"
	"fmt"
	"games-api/config"
	"games-api/gamefeed/gamefeedpb"
	"google.golang.org/grpc"
	"log"
)

var cltConnStr = fmt.Sprintf("%s:%d", config.ClientHost, config.ClientPort)

func main(){
	fmt.Println("Hello my master! I'm your slave. I live to serve you. What's your wishes? ... ")

	conn, err := grpc.Dial(cltConnStr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}

	defer conn.Close()

	client := gamefeedpb.NewGameFeedServiceClient(conn)

	doUnary(client)
}

func doUnary(client gamefeedpb.GameFeedServiceClient) {
	req := &gamefeedpb.GameFeedRequest{
		GameFeedInput: &gamefeedpb.GameFeedInput{
			Platform: "web",
			Template: "default",
			Brand: "casino.com",
			State: "user_state",
			Currency: "eur",
		},
	}

	res, err := client.GameFeed(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GamesAPI: %v", err)
	}

	log.Printf("%v", res.GameFeed)
}
