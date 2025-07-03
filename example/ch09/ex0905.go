package main

import (
	"errors"
	"fmt"
	"os"
)

type Status int  //liststart1

const ( // iotaについては「7.2.7 iotaと列挙型」参照
	InvalidLogin Status = iota + 1  // 不正なログイン
	NotFound // 見つからない
) //listend1

type StatusErr struct {  //liststart2
	Status  Status // 状態
	Message string // メッセージ
}

func (se StatusErr) Error() string { // StatusErrのメソッド
	return se.Message
}  //listend2

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {  //liststart3
	token, err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
							 			// 認証情報が無効
		}
	}
	data, err := getData(token, file)
	if err != nil {
		return nil, StatusErr{
			Status:  NotFound,
			Message: fmt.Sprintf("file %s not found", file),
		}
	}
	return data, nil
}  //listend3

func login(uid, pwd string) (string, error) {
	// world's worst login system
	if uid == "admin" && pwd == "admin" {
		return "user:admin", nil
	}
	return "", errors.New("bad user")
}

func getData(token, file string) ([]byte, error) {
	// world's worst access control
	if token == "user:admin" {
		switch file {
		case "secrets.txt":
			return []byte("passwords aplenty!"), nil
		case "payroll.csv":
			return []byte("everyone's salary"), nil
		}
	}
	return nil, os.ErrNotExist
}

func main() {
	data, err := LoginAndGetData("admin", "admin", "secrets.txt")
	fmt.Println(string(data), err)

	data, err = LoginAndGetData("admin", "admin", "chicken_recipe.txt")
	fmt.Println(string(data), err)

	data, err = LoginAndGetData("jon", "password", "secrets.txt")
	fmt.Println(string(data), err)
}
