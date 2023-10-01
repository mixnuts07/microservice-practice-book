// 実行可能なプログラムのエントリーポイントがmainパッケージになる.つまりmainパッケージにはmain()関数が必要
package main

import "fmt"

// 書籍情報の構造体
type Book struct {
	ID int
	Title string
	Author string
	Price int
}

// DBを使わずに書籍情報を返せるように構造体にデータを保持
var (
	book1 = Book{
		ID: 1,
		Title: "Go言語プログラミング",
		Author: "山田太郎",
		Price: 3600,
	}
	book2 = Book{
		ID: 2,
		Title: "Go言語でつくるインタプリタ",
		Author: "山田太郎2",
		Price: 36000,
	}
	books = []Book{book1, book2}
)


func getBooks(i int32) Book {
	return books[i-1]
}

func main() {
	fmt.Println("Hello, World!")
}