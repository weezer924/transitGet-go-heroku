package golib

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

//*****************************************************
func FindTransitdata(StartLoc int, EndLoc int) TransitRoutes {

	// データベース接続
	db, err := openCon()
	err = db.Ping()
	if err != nil {
		log.Printf("データベースの接続に失敗しました : %s", err.Error())
		os.Exit(1)
	}
	log.Printf("データベース 接続 \n")

	// データベース検索
	tranroutes, err := selectWithQuery(db, StartLoc, EndLoc)
	if err != nil {
		log.Printf("クエリーの実行に失敗しました : %s", err.Error())
		os.Exit(1)
	}

	defer db.Close()
	log.Printf("データベース切断 \n")

	return tranroutes
}

//*****************************************************
func openCon() (*sql.DB, error) {
	return sql.Open("postgres", os.Getenv("DATABASE_URL"))
}

//*****************************************************
func selectWithQuery(db *sql.DB, StartLoc int, EndLoc int) (TransitRoutes, error) {

	query := `select rt.identity,
	                 ptp.start_location,
	                 ptp.start_point_name,
	                 ptp.start_datetime,
                     ptp.end_location,
                     ptp.end_point_name,
                     ptp.end_datetime,
                     ptp.fare,
                     ptp.parent
                from route as rt, route_point_to_point as rpp, point_to_point as ptp
               where rt.start_location  = $1  and
                     rt.end_location    = $2  and
                     rt.identity        = rpp.route and
                     rpp.point_to_point = ptp.identity
               order by rpp.sort_order`

	rows, err := db.Query(query, StartLoc, EndLoc)

	var tranroutes TransitRoutes
	if err != nil {
		return tranroutes, err
	}

	for rows.Next() {
		var tr TransitRoute
		if err = rows.Scan(
			&tr.ID, &tr.StartLoc, &tr.StartPoint, &tr.StartTime, &tr.EndLoc,
			&tr.EndPoint, &tr.EndTime, &tr.FARE, &tr.PARENT); err != nil {
			return tranroutes, err
		}
		tranroutes = append(tranroutes, tr)
	}
	rows.Close()
	return tranroutes, err
}
