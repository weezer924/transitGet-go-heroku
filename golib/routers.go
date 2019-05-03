package golib

import (
	"net/http"
)

type Route struct {
	Name        string           // リスクエストを受け取る関数の名称
	Method      string           // HTTPリクエストのメソッド名
	Pattern     string           // URLパターン
	HandlerFunc http.HandlerFunc // リクエストを受け取るハンドラ関数
}

type Routes []Route

//*****************************************************
func NewRouter() {

	// URLパターン毎のハンドラ関数を登録
	for _, route := range routes {
		var handler http.HandlerFunc
		handler = route.HandlerFunc

		http.HandleFunc(route.Pattern, handler)
	}
}

//*****************************************************
var routes = Routes{
	Route{
		"TransitGet",
		"GET",
		"/v1/transit",
		TransitGet,
	},
}
