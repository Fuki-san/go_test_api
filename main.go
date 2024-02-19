package main

import (
	//fmtパッケージをインポート。フォーマットを扱うための関数が含まれている。
	//"fmt"
	//httpパッケージをインポート。httpを扱うための関数が含まれている。
	"net/http"
	//ginパッケージをインポート。ginを使うための関数が含まれている。
	"github.com/gin-gonic/gin"
)
//構造体の定義。jsonタグを使ってjsonのキーを指定。
type Item struct{
	//構造はキャメルケース。jsonタグはsnakeでバッククォート
	Message string `json:"message"`
	ItemName string `json:"item_name"`
}
//構造体のスライスを作成。MessageにHello Worldを代入。
var items = []Item{
	{Message: "item received: <item1>"},
	{ItemName: "item1"},
}
//c*gin.Contextを引数に取る関数を定義。c.IndentedJSONでjsonを返す。
//http.StatusOKでステータスコードを指定。
//hellosを返す。
// func getItems(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, items)
// }
func postItems(c *gin.Context){
	var newItem Item
	//if ショートハンド() 条件式。nilはエラーがないことを示す。
	if err := c.BindJSON(&newItem); err != nil{
		return
		}

	//add new item to the slice
	items = append(items, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)
}
//main関数。ginのインスタンスを作成。
//GETメソッドで/hellosにアクセスがあった場合、getHellos関数を実行。
func main() {
	// ginのインスタンスを作成。
	r := gin.Default()
	//GETメソッドで/hellosにアクセスがあった場合、getHellos関数を実行。
	// r.GET("/items", getItems)
	r.POST("/items", postItems)
	//サーバーを起動。
	r.Run(":8000")
	}
