package main

import (
	"github.com/Krapiy/noData-chat-API/db"
)

func main() {
	_, err := db.New("root:@tcp(mysql:3306)/noData-chat-API?timeout=1s&readTimeout=10s&multiStatements=true")
	if err != nil {
		panic(err)
	}
}
