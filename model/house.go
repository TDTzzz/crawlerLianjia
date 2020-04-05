package model

import "time"

type HouseInfo struct {
	HouseType string
	Area      string
	Toward    string
	Level     string
	Floor     string
	BuildYear string
	BuildType string
	Villa     string
}

type HouseDetail struct {
	Id          int
	Region      string
	SubRegion   string
	Name        string
	Community   string
	CommunityId int
	TotalPrice  float64
	UnitPrice   float64
	Area        float64
	HouseType   string
	Toward      string
	Floor       string
	BuildYear   string
	BuildType   string
	Villa       string
	Date        time.Time
}
