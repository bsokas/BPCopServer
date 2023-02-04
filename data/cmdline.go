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

    for {
      fmt.Println()
      fmt.Println("Available operations: ")
      fmt.Println("1. Input blood pressure reading")
      fmt.Println("2. Input meditation log")
      fmt.Println("3. Print blood pressure readings")
      fmt.Println("4. Print meditiations logs")
      fmt.Println("Type 'exit' or CTRL-C to stop")
      fmt.Println()

      fmt.Print("Enter the operation number to continue: ")
      if option, readErr := reader.ReadString('\n'); readErr == nil {
        trimmed := strings.TrimSpace(option)
        switch trimmed {
        case "1":
          inputBP()
        case "2":
          inputMeditation()
        case "3":
          readBloodPressure()
        case "4":
          readMeditation()
        case "exit":
          os.Exit(0)
        }
      } else {
        log.Fatal(readErr)
      }
    }

}

func inputBP() { fmt.Println("To be implemented") }

func inputMeditation() { fmt.Println("To be implemented") }

func readBloodPressure() {
  rows, readErr := BPDatabase.Query("SELECT * FROM blood_pressure_reading")
  if readErr != nil {
    log.Fatal(readErr)
  }

  defer rows.Close()

  // This only works for blood_pressure_reading rows
  for rows.Next() {
    var reading BloodPressureReading
    if readErr := rows.Scan(&reading.ID, &reading.SystolicMMHg, &reading.DiastolicMMHg, &reading.HeartRateBpm, &reading.CreatedAt, &reading.RecordedAt, &reading.TripleReading, &reading.Notes); readErr != nil {
      log.Fatal(readErr)
    }

    fmt.Printf("%+v\n", reading)
  }
}

func readMeditation() {
  rows, readErr := BPDatabase.Query("SELECT * FROM meditation_log")
  if readErr != nil {
    log.Fatal(readErr)
  }

  for rows.Next() {
    var medLog MeditationLog
    if readErr := rows.Scan(&medLog.ID, &medLog.CreatedAt, &medLog.MeditatedAt, &medLog.Rating, &medLog.Comments); readErr != nil {
      log.Fatal(readErr)
    }

    fmt.Printf("%+v\n", medLog)
  }
}
