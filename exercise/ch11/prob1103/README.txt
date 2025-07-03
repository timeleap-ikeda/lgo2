# 11章 練習問題3 

ARM64版Windows用には次のコマンドを実行します。

$ GOOS=windows GOARCH=arm64 go build

prob1103.exeというファイルができます。

次のコマンドで確認できます（実行するにはWindowsマシンにコピーして実行してください）

$ file prob1103.exe


Linux AMD64用には次を実行します。

$ GOOS=linux GOARCH=amd64 go build

prob1103というファイルができます。
