package main

import (
	"fmt"
	"time"

	"github.com/Krapiy/noData-chat-API/db"
	"github.com/Krapiy/noData-chat-API/db/fixtures"
	"github.com/Krapiy/noData-chat-API/ws"
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

	err = ws.StartServer()
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println("run")
		time.Sleep(time.Second * 1)
	}
}
