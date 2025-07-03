// この例はエラーになります

package main

import(
	"fmt"
)

// インタフェースOrderable（順序付け可能）の定義  //liststart1
type Orderable interface {
	Order(any) int  // このメソッドが定義されていればOrderableになる
	// メソッドOrderは次のような値を戻す：
	//  ・Orderable（レシーバー）が渡された値（メソッドの引数any）よりも小さい:
	//      0よりも小さい値を返す
	//  ・Orderableが渡された値よりも大きい: 0よりも大きい値を返す
	//  ・2つの値が等しい: 0を返す
} //listend1

type Tree struct { //liststart2
	val         Orderable
	left, right *Tree
}

// Treeに付随するメソッドInsertの定義。ポインタレシーバーになっている
// 二分木tにvalに指定された要素（ノード）を挿入
func (t *Tree) Insert(val Orderable) *Tree {
	if t == nil { // 二分木がnilだったら
		return &Tree{val: val} // valを唯一の値とする二分木を返す
		// left, right はゼロ値になるのでnil
	}

	switch comp := val.Order(t.val); { // 比較（COMParison）の結果
	case comp < 0: // valのほうが自分(t)よりも小さい
		t.left = t.left.Insert(val) // 左の子として挿入
	case comp > 0: //valのほうが自分(t)よりも大きい
		t.right = t.right.Insert(val) // 右の子として挿入
	}
	
	return t  // valを挿入した結果の二分木を返す
	// tとvalが同じ順序のときは何もせずに同じ二分木を返す
} //listend2


// 型OrderableIntの定義  //liststart3
type OrderableInt int

// OrderableInt型（の変数）をレシーバーとするメソッドOrderの定義
// これが定義されることでOrderableIntはOrderableになる
// （インタフェースOrderableを満たす）
func (oi OrderableInt) Order(val any) int {
	return int(oi - val.(OrderableInt))
	// 自分自身（oi）とvalをOrderableIntにして比較し、結果を返す
	// val.(OrderableInt)は「7.11.1 型アサーション」参照
	// 同じOrderbleIntにしないと「-」ができない
	// 「return int(oi) - val.(int) 」でもOK
} //listend3


func main() { //liststart5
	var it *Tree
	it = it.Insert(OrderableInt(5))
	it = it.Insert(OrderableInt(3))
	fmt.Println(it.val) // 5  // topのノードのvalの値を表示  //listend5
} 
