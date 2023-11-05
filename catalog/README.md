### GO コマンド

```go
//  新しいプロジェクトの作成とgo.modファイルを作る.
go mod init <project-name>
// プログラムの実行
go run main.go
// 依存関係の追加。依存関係が追加されるとgo.sumファイルが作成され.
go get github.com/hoge/fuga@latest
// protocコンパイラを使って、proto定義からデータアクセスインターフェースのコードを自動生成する。
// 生成されたファイルを使うことでgRPC通信を行うサーバやクライアントを定義できる。
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative catalogue.proto
// GOプロジェクトでgRPCを利用できるようにする。go.sumファイルが生成される。
go get google.golang.org/grpc
```
