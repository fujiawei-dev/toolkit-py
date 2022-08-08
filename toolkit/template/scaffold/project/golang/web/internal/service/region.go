package service

import (
	"errors"

	"github.com/guonaihong/gout"
	"github.com/tidwall/gjson"
)

const key = "8e3a36c2520ca68c2c8ae1b202213490"

// GetForwardGeocodingInformation 获取正向地理编码信息
func GetForwardGeocodingInformation(address string) (location string, err error) {
	var information string

	err = gout.
		GET("https://restapi.amap.com/v3/geocode/geo").
		// Debug(true).
		SetQuery(gout.H{
			"output":  "json",
			"address": address,
			"key":     key,
		}).
		BindBody(&information).
		Do()

	if err != nil {
		return
	}

	status := gjson.Get(information, "status").String()
	if status != "1" {
		err = errors.New(gjson.Get(information, "info").String())
		return
	}

	geocodes := gjson.Get(information, "geocodes")
	location = geocodes.Get("0.location").String()

	return
}

// GetReverseGeocodingInformation 获取逆向地理编码信息
func GetReverseGeocodingInformation(location string) (address string, err error) {
	var information string

	err = gout.
		GET("https://restapi.amap.com/v3/geocode/regeo").
		// Debug(true).
		SetQuery(gout.H{
			"output":   "json",
			"location": location,
			"key":      key,
			"radius":   "1000",
		}).
		BindBody(&information).
		Do()

	if err != nil {
		return
	}

	status := gjson.Get(information, "status").String()
	if status != "1" {
		err = errors.New(gjson.Get(information, "info").String())
		return
	}

	address = gjson.Get(information, "regeocode.formatted_address").String()

	return
}
