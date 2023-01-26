package main

import (
  "fmt"
  "github.com/bsokas/BPCopServer/data"
)

func main() {
	fmt.Println("Booting BPCopServer")
  fmt.Println("Attempting database connection.......")

  data.Connect()
}
