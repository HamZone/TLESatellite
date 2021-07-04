# 前言
* Forked from [joshuaferrara/go-satellite](https://github.com/joshuaferrara/go-satellite)
* 原仓库协议：[BSD 2-Clause License](https://github.com/joshuaferrara/go-satellite/blob/master/LICENSE)
* 原 repo 的浮点数计算可能存在精度问题，所以新建一个用新的方式计算部分浮点数
* 在原有基础上 可能做出部分修改，HamZone.cn 相关 TLE 计算会使用此仓库，方便自主调整
* 仓库内容中文本土化
* `master` 分支保持和原仓库一致，`emin` 为默认分支

# TLESatellite
[![GoDoc](https://godoc.org/github.com/HamZone/TLESatellite?status.svg)](https://godoc.org/github.com/HamZone/TLESatellite)

```go
import "github.com/HamZone/TLESatellite"`
```

项目使用 SGP4 模型进行预测

## 用法

#### Constants


#### func  ECIToLLA

```go
func ECIToLLA(eciCoords Vector3, gmst float64) (altitude, velocity float64, ret LatLong)
```
Convert Earth Centered Inertial coordinated into equivalent latitude, longitude,
altitude and velocity. Reference: http://celestrak.com/columns/v02n03/

#### func  GSTimeFromDate

```go
func GSTimeFromDate(year, mon, day, hr, min, sec int) float64
```
Calc GST given year, month, day, hour, minute and second

#### func  JDay

```go
func JDay(year, mon, day, hr, min, sec int) float64
```
Calc julian date given year, month, day, hour, minute and second the julian date
is defined by each elapsed day since noon, jan 1, 4713 bc.

#### func  Propagate

```go
func Propagate(sat Satellite, year int, month int, day, hours, minutes, seconds int) (position, velocity Vector3)
```
Calculates position and velocity vectors for given time

#### func  ThetaG_JD

```go
func ThetaG_JD(jday float64) (ret float64)
```
Calculate GMST from Julian date. Reference: The 1992 Astronomical Almanac, page
B6.


#### type LatLong

```go
type LatLong struct {
	Latitude, Longitude float64
}
```

Holds latitude and Longitude in either degrees or radians

#### func  LatLongDeg

```go
func LatLongDeg(rad LatLong) (deg LatLong)
```
Convert LatLong in radians to LatLong in degrees

#### type LookAngles

```go
type LookAngles struct {
	Az, El, Rg float64
}
```

Holds an azimuth, elevation and range

#### func  ECIToLookAngles

```go
func ECIToLookAngles(eciSat Vector3, obsCoords LatLong, obsAlt, jday float64) (lookAngles LookAngles)
```
Calculate look angles for given satellite position and observer position obsAlt
in km Reference: http://celestrak.com/columns/v02n02/

#### type Satellite

```go
type Satellite struct {
	Line1 string
	Line2 string
}
```

Struct for holding satellite information during and before propagation

#### func  ParseTLE

```go
func ParseTLE(line1, line2, gravconst string) (sat Satellite)
```
Parses a two line element dataset into a Satellite struct

#### func  TLEToSat

```go
func TLEToSat(line1, line2 string, gravconst string) Satellite
```
Converts a two line element data set into a Satellite struct and runs sgp4init

#### type Vector3

```go
type Vector3 struct {
	X, Y, Z float64
}
```

Holds X, Y, Z position

#### func  ECIToECEF

```go
func ECIToECEF(eciCoords Vector3, gmst float64) (ecfCoords Vector3)
```
Convert Earth Centered Intertial coordinates into Earth Cenetered Earth Final
coordinates Reference: http://ccar.colorado.edu/ASEN5070/handouts/coordsys.doc

#### func  LLAToECI

```go
func LLAToECI(obsCoords LatLong, alt, jday float64) (eciObs Vector3)
```
Convert latitude, longitude and altitude into equivalent Earth Centered
Intertial coordinates Reference: The 1992 Astronomical Almanac, page K11.
