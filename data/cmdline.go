package data

import (
  "fmt"
  "log"
  "bufio"
  "os"
  "strings"
)

func RunDBCLIInterface() {
  fmt.Println("Establishing DB connection...")
  Connect()

  startCmdLine()
}

func startCmdLine() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("Available operations: ")
    fmt.Println("1. Input blood pressure reading")
    fmt.Println("2. Input meditation log")
    fmt.Println("3. Print blood pressure readings")
    fmt.Println("4. Print meditiations logs")

    fmt.Print("Enter the operation number to continue: ")
    if option, readErr := reader.ReadString('\n'); readErr == nil {
      trimmed := strings.TrimSpace(option)
      switch trimmed {
      case "1":
        inputBP()
      case "2":
        inputMeditation()
      case "3":
      case "4":
      }
    } else {
      log.Fatal(readErr)
    }
}

func inputBP() { fmt.Println("To be implemented") }

func inputMeditation() { fmt.Println("To be implemented") }
