package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func clockHandler(w http.ResponseWriter, r *http.Request) {
	// テンプレートファイルをパース
	tpl := template.Must(template.ParseFiles("./tpls/clock.html.tpl"))
	// 構造体の定義
	type DateTime struct {
		Date string
		Time string
	}
	// 構造体
	now := DateTime{
		// 1月2日 15(3)時4分5秒 2006年 という驚きのフォーマット
		// アメリカ式日付表記を123456で並べたもの
		Date: time.Now().Format("2006-01-02"),
		Time: time.Now().Format("15:04:05"),
	}
	/* マップだとこう
	// マップ [キー] バリュー
	tplMap := map[string]string{
		"Date": time.Now().Format("2006-01-02"),
		"Time": time.Now().Format("15:04:05"),
	}
	*/
	// テンプレートにマップで値を埋め込み
	if err := tpl.Execute(w, now); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// ハンドラーを登録
	http.HandleFunc("/clock", clockHandler)
	// 静的ファイルハンドラー (今回はコメントアウト)
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
	// サーバ起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}
