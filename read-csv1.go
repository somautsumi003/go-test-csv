package sample

import (
    "encoding/csv"
    "fmt"
    "log"
    "strings"
)

func Sample() {
    // テスト用文字列
    str := "test,テスト\nHello,こんにちは"

    // CSVのReaderを用意
    r := csv.NewReader(strings.NewReader(str))

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
