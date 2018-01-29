package main

import (
  "encoding/csv"
  "fmt"
  "io"
  "strings"
)

func main(){
  lines := []string{
  "名前,部署",
  "山田,マーケ",
  "田中,人事",
  }
  csvStr := strings.Join(lines, "\n")
  r := csv.NewReader(strings.NewReader(csvStr))
  for {
    record, err := r.Read()
    if err == io.EOF {
      break
    }
    if err != nil {
      fmt.Println("Read error: ", err )
      break
    }
    fmt.Printf("record0[%s] record1[%s]\n", record[0], record[1])
  }
}
