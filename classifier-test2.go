package team

import(
  "github.com/zuu-development/papillon/src/lib/frame"
  "encoding/csv"
  "log"
  "strings"
)


type Team struct {
  Text []string
}


func New(c *frame.PapillonContext) *User {
	return &User{
		base: model.NewBase(c),
	}
}

func Reader() {
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
      return v
  }
}
