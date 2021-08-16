package database

import (
	"database/sql"
	"fmt"
	_ "fmt"
	"games-api/config"
	"games-api/gamefeed/gamefeedpb"
	_ "github.com/go-sql-driver/mysql"
)

var dbConnString = fmt.Sprintf(
	"%s:%s@tcp(%s:%d)/%s",
	config.DBUser,
	config.DBPass,
	config.DBHost,
	config.DBPort,
	config.DBName,
)

func GetGameFeeds() []gamefeedpb.GameFeed {
	db, err := sql.Open("mysql", dbConnString)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var gameFeeds []gamefeedpb.GameFeed

	dbRes, err := db.Query("SELECT id, name FROM gamefeeds")
	if err != nil {
		panic(err.Error())
	}

	for dbRes.Next() {
		var gameFeed gamefeedpb.GameFeed
		err = dbRes.Scan(&gameFeed.Id, &gameFeed.Name)
		if err != nil {
			panic(err.Error())
		}

		target := make([]gamefeedpb.GameFeed, len(gameFeeds)+1)
		copy(target, gameFeeds)
		gameFeeds = append(target, gameFeed)
	}

	return gameFeeds
}


