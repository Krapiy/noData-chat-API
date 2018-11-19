package main

import (
	"github.com/Krapiy/noData-chat-API/db"
	"github.com/Krapiy/noData-chat-API/db/fixtures"
)

func main() {
	client, err := db.New("root:@tcp(mysql:3306)/noData-chat-API?timeout=1s&readTimeout=10s&multiStatements=true")
	if err != nil {
		panic(err)
	}

	err = fixtures.Load(client)
	if err != nil {
		panic(err)
	}
}
