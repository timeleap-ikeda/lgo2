// 別の例題
package main

import(
	"fmt"
	"net/http"
	"time"
	"context"
	"encoding/json"
)

func main() {
	client := http.Client{
		Timeout: 10 * time.Second, // タイムリミットの設定  10秒
	}

	url := "https://www.dictjuggler.net/api/yakugo-api/?word=friend&apiId=js1daylab" // 辞書データをJSONで返す。訳者が運営するサイト
	// 簡単なAPI解説： https://www.dictjuggler.net/api
	
	req, err := http.NewRequestWithContext(
		context.Background(), // 14章参照。当面はこのまま利用
		http.MethodGet, // HTTPのメソッド"GET"に対応する定数
		url,
		nil) // io.Readerでボディを指定（この場合はなし）
	if err != nil {
		panic(err)
	}

	req.Header.Add("X-My-Client", "Learning Go")
	// リクエストのヘッダに追加。この場合はなくても結果は同じ...
	res, err := client.Do(req) // 準備したリクエストを送信し結果をもらう
	if err != nil {
		panic(err)
	}

	
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("unexpected status: got %v", res.Status))
	}
	fmt.Println(res.Header.Get("Content-Type"))
	var data struct {
		ReturnCode    int      `json:"returnCode"`
		HeadWord      string   `json:"headWord"`
		YakugoList    []string `json:"yakugoList"`
		Idioms        []string `json:"Idioms"`
	}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", data.ReturnCode) // %+vでフィールド名付き表示
	fmt.Printf("%+v\n", data.HeadWord)
	fmt.Printf("%+v\n", data.YakugoList)
	fmt.Printf("%+v\n", data.Idioms)
}
