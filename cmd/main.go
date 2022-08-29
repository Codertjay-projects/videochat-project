package main

import (
	"log"
	server "videochat-project/internal/server"
)

func main() {
	if err := server.RUN(); err != nil {
		log.Fatalln(err.Error())
	}

}