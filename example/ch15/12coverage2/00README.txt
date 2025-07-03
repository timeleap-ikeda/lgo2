次の2つのコマンドを実行してください。

$ go test -v -cover -coverprofile=c.out
$ go tool cover -html=c.out
