package main

import (
	"blog-backend/server"
	_ "github.com/wpliap/common-wrap"
)

func main() {
	s := server.NewServer()
	s.Run()
}
