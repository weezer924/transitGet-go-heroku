package golib

import (
	"encoding/json"
	"net/http"
	"strconv"
)

//*****************************************************
func TransitGet(w http.ResponseWriter, r *http.Request) {

	// リクエストパラメータの取得
	q := r.URL.Query()
	stloc, _ := strconv.Atoi(q.Get("start_location"))
	edloc, _ := strconv.Atoi(q.Get("end_location"))

	// 乗換情報検索
	ts := FindTransitdata(stloc, edloc)

	if ts != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(ts); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
