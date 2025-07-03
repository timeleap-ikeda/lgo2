// この例はエラーになります

package main

import(
	"fmt"
	"strings"
)

// インタフェースOrderable（順序付け可能）の定義  //liststart1
type Orderable interface {
	Order(any) int  // このメソッドが定義されていればOrderableになる
	// メソッドOrderは次のような値を戻す：
	//  ・Orderableが渡された値よりも小さい場合、0よりも小さい値を返す
	//  ・Orderableが渡された値よりも大きい場合、0よりも大きい値を返す
	//  ・2つの値が等しい時0を返す
} //listend1

type Tree struct { //liststart2
	val         Orderable
	left, right *Tree
}

// Treeに付随するメソッドInsertの定義。ポインタレシーバーになっている
// 二分木tにvalに指定された要素（ノード）を挿入
func (t *Tree) Insert(val Orderable) *Tree {
	if t == nil { //二分木がnilだったら
		return &Tree{val: val} //valを唯一の値とする二分木を返す
		//left, right はゼロ値になるのでnil
	}

	switch comp := val.Order(t.val); { //比較（COMParison）の結果
	case comp < 0: //valのほうが自分(t)よりも小さい
		t.left = t.left.Insert(val)
	case comp > 0: //valのほうが自分(t)よりも大きい
		t.right = t.right.Insert(val)
	}
	
	return t  // valを挿入した結果の二分木を返す
	// 同じ場合は何もしない
} //listend2


// 型OrderableIntの定義  //liststart3
type OrderableInt int


// OrderableInt型（の変数）をレシーバーとするメソッドOrderの定義
// これが定義されることでOrderableIntはOrderableになる
// （インタフェースOrderableを満たす）
func (oi OrderableInt) Order(val any) int {
	return int(oi - val.(OrderableInt))
	// 自分自身（oi）とvalをintにして比較し、結果を返す
} //listend3


// 型OrderableStringの定義  //liststart4
type OrderableString string

// 付随する（この型をレシーバーとする）唯一のメソッドの定義
// この型OrderableStringもOrderableになる
func (os OrderableString) Order(val any) int {
	return strings.Compare(string(os), string(val.(OrderableString)))
	// 自分自身（os）とvalをstringにして比較し、結果を返す
} //listend4


func main() {  //liststart5
	var it *Tree
	it = it.Insert(OrderableInt(5))
	it = it.Insert(OrderableString("nope"))  //listend5
	fmt.Println(it.val)
} 
