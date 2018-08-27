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
  //"reflect"
  "path/filepath"
)

func main() {
  csvPath := flag.String("csv", "problems.csv", "will be a string")

  // TODO: Use the timer flag below to create a timer
  timer := flag.Int("timer", 30, "timer quiz assessment")

  // TODO: Find a way to refactor below

  flag.Parse()

  fmt.Println("timer ", *timer)

  pwd, _ := os.Getwd()

  csvFile, _ := filepath.Abs(pwd + "/" + *csvPath)

  fileContents, _ := os.Open(csvFile)

  record := csv.NewReader(bufio.NewReader(fileContents))

  problems, err := record.ReadAll()

  fmt.Println("problems ", problems)

  if err != nil {
    log.Fatal(err)
  }

  counter :=  0

  for i, p := range problems {
    reader := bufio.NewReader(os.Stdin)
    problem, answer := p[0], p[1]

    fmt.Print(i + 1, ".) What is ", problem, " Sir? \n")

    text, _ := reader.ReadString('\n')
    userAnswer := strings.TrimSuffix(text, "\n")

    if userAnswer != answer {
      counter++
    }
  }

  correctCount := (len(problems) - counter);
  fmt.Println("Got ", correctCount, "/", len(problems), " right!")
  /*
   *TODO: Part Two:Adapt your program from part 1 to add a timer. The default time limit should be 30 seconds, but should also be customizable via a flag.

   Your quiz should stop as soon as the time limit has exceeded. That is, you shouldn’t wait for the user to answer one final questions but should ideally stop the quiz entirely even if you are currently waiting on an answer from the end user.


   */


}

