package satellite

import (
	"log"
	"math"
)

// this procedure converts the day of the year, epochDays, to the equivalent month day, hour, minute and second.
func days2mdhms(year int64, epochDays float64) (mon, day, hr, min, sec float64) {
	lmonth := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if year%4 == 0 {
		lmonth = [12]int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	}

	dayofyr := math.Floor(epochDays)

	i := 1.0
	inttemp := 0.0

	for dayofyr > inttemp+float64(lmonth[int(i-1)]) && i < 22 {
		inttemp = inttemp + float64(lmonth[int(i-1)])
		i += 1
	}

	mon = i
	day = dayofyr - inttemp

	temp := (epochDays - dayofyr) * 24.0
	hr = math.Floor(temp)

	temp = (temp - hr) * 60.0
	min = math.Floor(temp)

	sec = (temp - min) * 60.0

	return
}

//JDay 计算儒略日
//Calc julian date given year, month, day, hour, minute and second
//the julian date is defined by each elapsed day since noon, jan 1, 4713 bc.
func JDay(year, mon, day, hr, min, sec int) float64 {
	return (367.0*float64(year) - math.Floor((7*(float64(year)+math.Floor((float64(mon)+9)/12.0)))*0.25) + math.Floor(275*float64(mon)/9.0) + float64(day) + 1721013.5 + ((float64(sec)/60.0+float64(min))/60.0+float64(hr))/24.0)
}

//gstime 格林威治恒星时间（iau-82）
//this function finds the greenwich sidereal time (iau-82)
func gstime(jdut1 float64) (temp float64) {
	tut1 := DecimalDiv2(DecimalSub2(jdut1, 2451545.0), 36525.0)
	temp = DecimalAdd3(DecimalMul4(-6.2e-6, tut1, tut1, tut1), DecimalMul3(0.093104, tut1, tut1)+DecimalMul2(DecimalAdd2(DecimalMul2(876600.0, 3600), 8640184.812866), tut1), 67310.54841)
	temp = math.Mod((DecimalDiv2(DecimalMul2(temp, DEG2RAD), 240.0)), TWOPI)
	if temp < 0.0 {
		temp = DecimalAdd2(temp, TWOPI)
	}
	return
}

//GSTimeFromDate 计算Greenwich Mean Sidereal Time 格林威治平均恒星时
//Calc GST given year, month, day, hour, minute and second
func GSTimeFromDate(year, mon, day, hr, min, sec int) float64 {
	jDay := JDay(year, mon, day, hr, min, sec)
	return gstime(jDay)
}

//ECIToLLA 将地心惯性坐标转换为等效的纬度(latitude)、经度(longitude)、高度(altitude)和速度(velocity)
// Convert Earth Centered Inertial coordinated into equivalent latitude, longitude, altitude and velocity.
// Reference: http://celestrak.com/columns/v02n03/
func ECIToLLA(eciCoords Vector3, gmst float64) (altitude, velocity float64, ret LatLong) {
	a := 6378.137                          // Semi-major Axis
	b := 6356.7523142                      // Semi-minor Axis
	f := DecimalDiv2(DecimalSub2(a, b), a) // Flattening
	e2 := DecimalSub2(DecimalMul2(2, f), math.Pow(f, 2))
	sqx2y2 := math.Sqrt(DecimalAdd2(math.Pow(eciCoords.X, 2), math.Pow(eciCoords.Y, 2)))
	// Spherical Earth Calculations
	longitude := DecimalSub2(math.Atan2(eciCoords.Y, eciCoords.X), gmst)
	latitude := math.Atan2(eciCoords.Z, sqx2y2)
	// Oblate Earth Fix
	C := 0.0
	for i := 0; i < 20; i++ {
		C = DecimalDiv2(1, math.Sqrt(1-e2*(DecimalMul2(math.Sin(latitude), math.Sin(latitude)))))
		latitude = math.Atan2(eciCoords.Z+(DecimalMul4(a, C, e2, math.Sin(latitude))), sqx2y2)
	}
	// Calc Alt
	altitude = DecimalSub2(DecimalDiv2(sqx2y2, math.Cos(latitude)), DecimalMul2(a, C))
	// Orbital Speed ≈ sqrt(μ / r) where μ = std. gravitaional parameter
	velocity = math.Sqrt(DecimalDiv2(398600.4418, DecimalAdd2(altitude, 6378.137)))
	ret.Latitude = latitude
	ret.Longitude = longitude
	return
}

//LatLongDeg 将以弧度表示的LatLong转换为以度表示的LatLong
// Convert LatLong in radians to LatLong in degrees
func LatLongDeg(rad LatLong) (deg LatLong) {
	deg.Longitude = math.Mod(DecimalMul2(DecimalDiv2(rad.Longitude, math.Pi), 180), 360)
	if deg.Longitude > 180 {
		deg.Longitude = 360 - deg.Longitude
	} else if deg.Longitude < -180 {
		deg.Longitude = 360 + deg.Longitude
	}

	if rad.Latitude < (-math.Pi/2) || rad.Latitude > math.Pi/2 {
		log.Fatal("Latitude not within bounds -pi/2 to +pi/2")
	}
	deg.Latitude = (rad.Latitude / math.Pi * 180)
	return
}

// Calculate GMST from Julian date.
// Reference: The 1992 Astronomical Almanac, page B6.
func ThetaG_JD(jday float64) (ret float64) {
	_, UT := math.Modf(jday + 0.5)
	jday = jday - UT
	TU := (jday - 2451545.0) / 36525.0
	GMST := 24110.54841 + TU*(8640184.812866+TU*(0.093104-TU*6.2e-6))
	GMST = math.Mod(GMST+86400.0*1.00273790934*UT, 86400.0)
	ret = 2 * math.Pi * GMST / 86400.0
	return
}

// Convert latitude, longitude and altitude(km) into equivalent Earth Centered Intertial coordinates(km)
// Reference: The 1992 Astronomical Almanac, page K11.
func LLAToECI(obsCoords LatLong, alt, jday float64) (eciObs Vector3) {
	re := 6378.137
	theta := math.Mod(ThetaG_JD(jday)+obsCoords.Longitude, TWOPI)
	r := (re + alt) * math.Cos(obsCoords.Latitude)
	eciObs.X = r * math.Cos(theta)
	eciObs.Y = r * math.Sin(theta)
	eciObs.Z = (re + alt) * math.Sin(obsCoords.Latitude)
	return
}

//ECIToECEF 将地心间坐标转换为地心最终坐标 Earth Centered Intertial
//Convert Earth Centered Intertial coordinates into Earth Cenetered Earth Final coordinates
//Reference: http://ccar.colorado.edu/ASEN5070/handouts/coordsys.doc
func ECIToECEF(eciCoords Vector3, gmst float64) (ecfCoords Vector3) {
	ecfCoords.X = DecimalAdd2(DecimalMul2(eciCoords.X, math.Cos(gmst)), DecimalMul2(eciCoords.Y, math.Sin(gmst)))
	ecfCoords.Y = DecimalAdd2(DecimalMul2(eciCoords.X, (-math.Sin(gmst))), DecimalMul2(eciCoords.Y, math.Cos(gmst)))
	ecfCoords.Z = eciCoords.Z
	return
}

//ECIToLookAngles 计算给定卫星位置和观察者位置的视角
// Calculate look angles for given satellite position and observer position
// obsAlt in km
// Reference: http://celestrak.com/columns/v02n02/
func ECIToLookAngles(eciSat Vector3, obsCoords LatLong, obsAlt, jday float64) (lookAngles LookAngles) {
	theta := math.Mod(ThetaG_JD(jday)+obsCoords.Longitude, 2*math.Pi)
	obsPos := LLAToECI(obsCoords, obsAlt, jday)

	rx := eciSat.X - obsPos.X
	ry := eciSat.Y - obsPos.Y
	rz := eciSat.Z - obsPos.Z

	top_s := math.Sin(obsCoords.Latitude)*math.Cos(theta)*rx + math.Sin(obsCoords.Latitude)*math.Sin(theta)*ry - math.Cos(obsCoords.Latitude)*rz
	top_e := -math.Sin(theta)*rx + math.Cos(theta)*ry
	top_z := math.Cos(obsCoords.Latitude)*math.Cos(theta)*rx + math.Cos(obsCoords.Latitude)*math.Sin(theta)*ry + math.Sin(obsCoords.Latitude)*rz

	lookAngles.Az = math.Atan(-top_e / top_s)
	if top_s > 0 {
		lookAngles.Az = lookAngles.Az + math.Pi
	}
	if lookAngles.Az < 0 {
		lookAngles.Az = lookAngles.Az + 2*math.Pi
	}
	lookAngles.Rg = math.Sqrt(rx*rx + ry*ry + rz*rz)
	lookAngles.El = math.Asin(top_z / lookAngles.Rg)

	return
}
