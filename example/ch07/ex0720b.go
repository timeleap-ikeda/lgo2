/*

   /hello  のほかに  /goodbye も有効にしたバージョン
   
1. $ go run ex0720b.go

2. 続いてブラウザで次のURLなどを表示（user_id= のあとに1, 2, 3などを指定）
	例：
		http://localhost:8080/hello?user_id=1
		http://localhost:8080/goodbye?user_id=2
*/

package main

import (
	"errors"
	"fmt"
	"net/http"
)

//ログを出力  //liststart1
func LogOutput(message string) {
	fmt.Println(message)
}  //listend1

// 文字列に対して文字列を記憶するマップ  //liststart2
type SimpleDataStore struct {  
	userData map[string]string
}

// SimpleDataStoreに付随するメソッド。userIDを受け取り、名前を返す
func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}  //listend2

// 新しいSimpleDataStoreを返すファクトリ関数  //liststart3
func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "花子",
			"2": "太郎",
			"3": "パット",
		},
	}
}  //listend3

// インタフェースDataStoreの定義  //liststart4
// SimpleDataStoreを抽象化したもの
type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

// インタフェースLoggerの定義
// LogOutputを抽象化したもの
type Logger interface {
	Log(message string)
}  //listend4

// 型LoggerAdapterを定義  //liststart5
type LoggerAdapter func(message string)

// LoggerAdapterに付随するメソッド
func (lg LoggerAdapter) Log(message string) {
	lg(message)  // 受け取ったメッセージmessageを引数として自分を呼び出す
}  //listend5

type SimpleLogic struct {  //liststart6
	l  Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("in SayHello for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("未登録のユーザーです")
	}
	return name + "さん、こんにちは", nil
} //listend6

func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	sl.l.Log("in SayGoodbye for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("未登録のユーザーです")
	}
	return name + "さん、さようなら" , nil
}

	//liststart7
func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}  //listend7

type Logic interface {
	SayHello(userID string) (string, error)
	SayGoodbye(userID string) (string, error)
}

type Controller struct {
	l     Logger
	logic Logic
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In SayHello")
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

func (c Controller) SayGoodbye(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In SayGoodbye")
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayGoodbye(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}


func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}

func main() {
	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	http.HandleFunc("/hello", c.SayHello)
	http.HandleFunc("/goodbye", c.SayGoodbye)
	http.ListenAndServe(":8080", nil)
}
