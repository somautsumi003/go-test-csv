package main

import(
  "encoding/csv"
  "fmt"
  "log"
  "strings"
)


// type Team struct {
//   Text []string
// }
func main() {
  Classify("test,テスト\nHello,こんにちは")
}


func Classify (c string)  {

  // CSVのReaderを用意
  r := csv.NewReader(strings.NewReader(c))

  // デリミタ(TSVなら\t, CSVなら,)設定
  r.Comma = ','

  // 全部読みだす
  records, err := r.ReadAll()
  if err != nil {
      log.Fatal(err)
  }

  // 各行でループ
  for _, v := range records {
      // 1列目
      // 2列目
      fmt.Println(v)
  }
}
