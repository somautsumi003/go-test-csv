package main
import (
  "encoding/csv"
  "io"
  "os"
  "fmt"
)

func main() {
  // csvオープン
  file, err := os.Open("test.csv")

	if err != nil {
		panic(err)
	}
	defer file.Close()
	// csv読み出し
	reader := csv.NewReader(file)

	for {
		record, err := reader.Read() // 1行読み出す  →スライスになる
		if err == io.EOF {
			break
		} else if err != nil {
  		panic(err)
  	}

  	fmt.Printf("%#v", record)
  }
}
