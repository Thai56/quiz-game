//package main
package main

import (
  "encoding/csv"
  "bufio"
	"fmt"
  "log"
  "os"
  "strings"
  "flag"
  "path/filepath"
  "time"
  "io"
)


type Problem struct {
  Question string
  Answer string
}

func getProblems(problems *[]Problem, problemsDone chan bool) {
  csvPath := flag.String("csv", "problems.csv", "will be a string")
  flag.Parse()
  pwd, _ := os.Getwd()
  csvFile, _ := filepath.Abs(pwd + "/" + *csvPath)
  fileContents, _ := os.Open(csvFile)
  reader := csv.NewReader(bufio.NewReader(fileContents))

  for {
    line, error := reader.Read()
    if error == io.EOF {
      break
    } else if error != nil {
      log.Fatal(error)
    }

    *problems = append(*problems, Problem{
      Question: line[0],
      Answer: line[1],
    })
  }
  problemsDone <- true

  close(problemsDone)
}

func startQuizGame(problems []Problem, quizDone chan bool, correctCount *int) {
  for i, p := range problems {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print(i + 1, ".) What is ", p.Question, " Sir? \n")
    text, _ := reader.ReadString('\n')
    userAnswer := strings.TrimSuffix(text, "\n")
    if userAnswer == p.Answer {
      *correctCount++
    }
  }
  quizDone <- true

  close(quizDone)
}

func main() {
  timerFlag := flag.Int("timer", 15, "timer quiz assessment")

  // Gathering Problems
  var problems []Problem
  var problemsDone = make(chan bool)
  go getProblems(&problems, problemsDone)
  <-problemsDone

  // Configuring Timer
  timeInSeconds := int32(*timerFlag)
  timer1 := time.NewTimer(time.Duration(timeInSeconds) * time.Second)

  // Starting Quiz
  var quizOver = make(chan bool)
  correctCount := 0
  go startQuizGame(problems, quizOver, &correctCount)

  fmt.Println("quiz starting...")

  for {
    select {
    case <-timer1.C:
      fmt.Println("\nTime ran out!...")
      fmt.Println("Got ", correctCount, "/", len(problems), " right!")
      return
    case <-quizOver:
      fmt.Println("Congradulations! You finished the quiz within the set time!...")
      fmt.Println("Got ", correctCount, "/", len(problems), " right!")
      return
    }
  }
}
