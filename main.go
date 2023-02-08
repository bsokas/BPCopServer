package main

import (
	"fmt"

	"github.com/bsokas/BPCopServer/data"
	"github.com/bsokas/BPCopServer/webserver"
)

func main() {
	fmt.Println("Booting BPCopServer")
	fmt.Println("Attempting database connection.......")

	// data.Connect()
  go webserver.Start()
	data.RunDBCLIInterface()
}
