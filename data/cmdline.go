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
      fmt.Println("\nAvailable operations: ")
      fmt.Println("1. Input blood pressure reading")
      fmt.Println("2. Input meditation log")
      fmt.Println("3. Print blood pressure readings")
      fmt.Println("4. Print meditiations logs")
      fmt.Printf("Type 'exit' or CTRL-C to stop\n\n")

      fmt.Print("Enter the operation number to continue: ")
      if option, readErr := reader.ReadString('\n'); readErr == nil {
        trimmed := strings.TrimSpace(option)
        switch trimmed {
        case "1":
          if bpErr := inputBP(reader); bpErr != nil {
            log.Fatal(bpErr)
          }
        case "2":
          if medErr := inputMeditation(reader); medErr != nil {
            log.Fatal(medErr)
          }
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

func cleanReadString(reader *bufio.Reader) string {
  val, err := reader.ReadString('\n')
  if err != nil {
    log.Fatal(err)
  }

  return val
}

func inputBP(reader *bufio.Reader) error {
  fmt.Printf("Systolic reading (mm Hg): ")
  systolic, sysErr := ValidatePressure(cleanReadString(reader))
  if sysErr != nil { return sysErr }

  fmt.Printf("Diastolic reading (mm Hg): ")
  diastolic, diaErr := ValidatePressure(cleanReadString(reader))
  if diaErr != nil { return diaErr }

  fmt.Printf("Heart rate (bpm): ")
  heartRate, hrErr := ValidateHeartRate(cleanReadString(reader))
  if hrErr != nil { return hrErr }

  fmt.Printf("Time of reading (format YYYY-MM-DD hh:mm:ss): ")
  recordedAt := strings.TrimSpace(cleanReadString(reader))

  fmt.Printf("Triple reading? (Y/N): ")
  tripleReading := ValidateTripleReading(cleanReadString(reader))

  fmt.Printf("Notes (optional): ")
  notes := ValidateNotes(cleanReadString(reader))

  result, insertErr := BPDatabase.Exec(`INSERT INTO blood_pressure_reading
    (systolic_mm_hg, diastolic_mm_hg, heart_rate_bpm, recorded_at, triple_reading, notes)
    VALUES (?, ?, ?, ?, ?, ?)`, systolic, diastolic, heartRate, recordedAt, tripleReading, notes)
  if insertErr != nil {
    return insertErr
  }

  if id, idErr := result.LastInsertId(); idErr != nil {
    return idErr
  } else {
    fmt.Printf("Successfully added new blood pressure reading, record %d\n", id)
    fmt.Printf("Sys: %d | Dia: %d | HR: %d | recordedAt: %s | tripleReading: %t | notes: %s\n", systolic, diastolic, heartRate, recordedAt, tripleReading, notes)
  }

  return nil
}

func inputMeditation(reader *bufio.Reader) error {
   fmt.Printf("Date & time of meditation (format YYYY-MM-DD hh:mm:ss): ")
   meditatedAt := strings.TrimSpace(cleanReadString(reader))

   fmt.Printf("Session rating (1-5, 5 as most relaxed): ")
   rating, ratingErr := ValidateRating(cleanReadString(reader))
   if ratingErr != nil { return ratingErr }

   fmt.Printf("Duration of meditation, in minutes: ")
   durationSec, durErr := ValidateDurationSeconds(cleanReadString(reader))
   if durErr != nil { return durErr }

   fmt.Printf("Any comments to add? (optional): ")
   comments := ValidateComments(cleanReadString(reader))

   newId, err := UpsertMeditationLog(meditatedAt, rating, durationSec, comments)
   if err != nil {
     return err
   }

   fmt.Printf("Successfully added meditation log with ID %d\n", newId)
   return nil
}

func readBloodPressure() {
  rows, readErr := BPDatabase.Query("SELECT * FROM blood_pressure_reading")
  if readErr != nil {
    log.Fatal(readErr)
  }

  defer rows.Close()

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
