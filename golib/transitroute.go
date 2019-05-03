package golib

type TransitRoute struct {
	ID         int    `json:"identity"`
	StartLoc   int    `json:"start_location"`
	StartPoint string `json:"start_point_name"`
	StartTime  string `json:"start_datetime"`
	EndLoc     int    `json:"end_location"`
	EndPoint   string `json:"end_point_name"`
	EndTime    string `json:"end_datetime"`
	FARE       int    `json:"fare"`
	PARENT     int    `json:"parent"`
}

type TransitRoutes []TransitRoute
