package controller

import (
	// import gin library
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//import API packages
	"github.com/miraikeitai2020/michisirube_camp_backend/recommend/pkg/server/model"
	"github.com/miraikeitai2020/michisirube_camp_backend/recommend/pkg/recommend"

	"math"
	"strconv"
)

const (
	EQUATORIAL_RADIUS = 6378137.0            // 赤道半径 GRS80
	POLAR_RADIUS      = 6356752.314          // 極半径 GRS80
	ECCENTRICITY      = 0.081819191042815790 // 第一離心率 GRS80
)

type Coodinate struct {
	Latitude  float64
	Longitude float64
}

func main() {
	a := Coodinate{35.65500, 139.74472} // 東京
	b := Coodinate{36.10056, 140.09111} // 筑波

	hd := hubenyDistance(a, b)
	distance := strconv.FormatFloat(math.Floor(hd), 'f', 0, 64)

	fmt.Printf("距離: " + distance + " m\n")
}

func hubenyDistance(src Coodinate, dst Coodinate) float64 {
	dx := degree2radian(dst.Longitude - src.Longitude)
	dy := degree2radian(dst.Latitude - src.Latitude)
	my := degree2radian((src.Latitude + dst.Latitude) / 2)

	W := math.Sqrt(1 - (Power2(ECCENTRICITY) * Power2(math.Sin(my)))) // 卯酉線曲率半径の分母
	m_numer := EQUATORIAL_RADIUS * (1 - Power2(ECCENTRICITY))         // 子午線曲率半径の分子

	M := m_numer / math.Pow(W, 3) // 子午線曲率半径
	N := EQUATORIAL_RADIUS / W    // 卯酉線曲率半径

	d := math.Sqrt(Power2(dy*M) + Power2(dx*N*math.Cos(my)))

	return d
}

func degree2radian(x float64) float64 {
	return x * math.Pi / 180
}

func Power2(x float64) float64 {
	return math.Pow(x, 2)
}
var(
	time 	model.Request.Time
	emo		model.Request.Emotion
	state	model.Request.Location

)
type Controller struct {
	DB *gorm.DB
}


func (ctrl *Controller) Task1(context *gin.Context) {
	err := context.BindJSON(&user)
	if err != nil {
		log.Println("[ERROR] Faild Bind JSON")
		context.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}
	
	context.JSON(200, gin.H{"location": {
		"id": recommend.Recommend(state.ID), "name": recommend.Recommend(state.Name), "latitude": recommend.Recommend(state.Latitude), "longitude": recommend.Recommend(state.Longitude)},
		"time": recommend.Recommend(time.Time)})
	context.JSON(200, gin.H{"time": recommend.Recommend(time.Time)})

}
