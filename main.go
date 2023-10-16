package main

import (
	"github.com/MahiroWatanabe/practiceapi/db"
	"github.com/MahiroWatanabe/practiceapi/server"
)

func main() {
	db.Init()
	server.Init()
}
