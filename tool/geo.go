package tool

import (
	"math"
)

// Geo 单例geo并导出
var Geo = &geo{}

type geo struct {
}

// GetDistance 获取两个坐标点的距离，单位m
func (*geo) GetDistance(lat1, lng1, lat2, lng2 float64) int64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515
	return int64(dist * 1.609344 * 1000)
}
