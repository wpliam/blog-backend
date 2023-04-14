package main

import (
	"blog-backend/server"
)

func main() {
	s := server.NewServer()
	s.Run()
}
