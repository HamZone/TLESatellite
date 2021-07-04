package satellite

import (
	"log"
	"math"
)

//GravConst 重力模型的一些变量
// Holds variables that are dependent upon selected gravity model
type GravConst struct {
	mu, radiusearthkm, xke, tumin, j2, j3, j4, j3oj2 float64
}

//getGravConst 重力模型 wgs72 模型似乎是卫星中最常用的跟踪 SGP4 算法中需要使用的参数
// Returns a GravConst with correct information on requested model provided through the name parameter
func getGravConst(name string) (grav GravConst) {
	switch name {
	case "wgs72old":
		grav.mu = 398600.79964
		grav.radiusearthkm = 6378.135
		grav.xke = 0.0743669161
		grav.tumin = DecimalDiv2(1.0, grav.xke)
		grav.j2 = 0.001082616
		grav.j3 = -0.00000253881
		grav.j4 = -0.00000165597
		grav.j3oj2 = DecimalDiv2(grav.j3, grav.j2)
	case "wgs72":
		grav.mu = 398600.8
		grav.radiusearthkm = 6378.135
		grav.xke = DecimalDiv2(60.0, math.Sqrt(DecimalDiv2(DecimalMul3(grav.radiusearthkm, grav.radiusearthkm, grav.radiusearthkm), grav.mu)))
		grav.tumin = DecimalDiv2(1.0, grav.xke)
		grav.j2 = 0.001082616
		grav.j3 = -0.00000253881
		grav.j4 = -0.00000165597
		grav.j3oj2 = DecimalDiv2(grav.j3, grav.j2)
	case "wgs84":
		grav.mu = 398600.5
		grav.radiusearthkm = 6378.137
		grav.xke = DecimalDiv2(60.0, math.Sqrt(DecimalDiv2(DecimalMul3(grav.radiusearthkm, grav.radiusearthkm, grav.radiusearthkm), grav.mu)))
		grav.tumin = DecimalDiv2(1.0, grav.xke)
		grav.j2 = 0.00108262998905
		grav.j3 = -0.00000253215306
		grav.j4 = -0.00000161098761
		grav.j3oj2 = DecimalDiv2(grav.j3, grav.j2)
	default:
		log.Fatal(name, "is not a valid gravity model")
	}

	return
}

// Not the movie
