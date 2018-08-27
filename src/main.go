//package main
package main

import (
	"encoding/csv"
  "bufio"
	"fmt"
	"log"
  "os"
  "strings"
  //"reflect"
)

func main() {
  filePath := os.Args[1]

  csvFile, _ := os.Open(filePath)

  record := csv.NewReader(bufio.NewReader(csvFile))


  problems, err := record.ReadAll()

  if err != nil {
    log.Fatal(err)
  }

  counter :=  0

  for _, p := range problems {
    reader := bufio.NewReader(os.Stdin)
    problem, answer := p[0], p[1]

    fmt.Print("What is", problem, " Sir? \n")

    text, _ := reader.ReadString('\n')
    userAnswer := strings.TrimSuffix(text, "\n")

    if userAnswer != answer {
      counter++
    }
  }

  fmt.Println("Got ", len(problems) - counter, "/", len(problems), " right!")
}

