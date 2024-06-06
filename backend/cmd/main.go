package main

import (
	"backend/internal/server"
	"backend/utils/util"
)

func main() {
	err := server.NewServer(":8080")
	utils.CheckError(err)
}
