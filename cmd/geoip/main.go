package main

import (
	"demos/lib"
	"fmt"
	"log"
	"net"

	"github.com/oschwald/geoip2-golang"
)

func main() {
	db, err := geoip2.Open("resources/geoip/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	// ip := net.ParseIP("81.2.69.142")
	// for i := 0; i < 256; i++ {
	// 	for j := 0; j < 256; j++ {
	// 		for k := 0; k < 256; k++ {
	// 			for n := 0; n < 256; n++ {
	// 				ipGen := fmt.Sprintf("%d.%d.%d.%d", i, j, k, n)
	// 				record, err := db.City(net.ParseIP(ipGen))
	// 				if err != nil {
	// 					continue
	// 				}
	// 				if name, ok := record.City.Names["zh-CN"]; ok {
	// 					fmt.Printf("IP: [%s], 国家: [%v], 城市: [%v]\n", ipGen, record.Country.Names["zh-CN"], name)
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	ip := "113.89.32.129"
	record, err := db.City(net.ParseIP(ip))
	if err != nil {
		log.Fatalln(err)
	}
	if name, ok := record.City.Names["zh-CN"]; ok {
		fmt.Printf("IP: [%s], 国家: [%v], 城市: [%v], 纬度: [%v], 经度: [%v]\n", ip, record.Country.Names["zh-CN"], name, record.Location.Latitude, record.Location.Longitude)

	}

	ip = "114.89.32.130"
	record1, err := db.City(net.ParseIP(ip))
	if err != nil {
		log.Fatalln(err)
	}
	if name, ok := record1.City.Names["zh-CN"]; ok {
		fmt.Printf("IP: [%s], 国家: [%v], 城市: [%v], 纬度: [%v], 经度: [%v]\n", ip, record1.Country.Names["zh-CN"], name, record1.Location.Latitude, record1.Location.Longitude)
	}

	fmt.Printf("(%v,%v) 与 (%v,%v) 相距%v千米\n",
		record.Location.Latitude, record.Location.Longitude, record1.Location.Latitude, record1.Location.Longitude,
		lib.Distance(record.Location.Latitude, record.Location.Longitude, record1.Location.Latitude, record1.Location.Longitude, "K"))

	// if len(record.Subdivisions) > 0 {
	// 	fmt.Printf("English subdivision name: %v\n", record.Subdivisions[0].Names["en"])
	// }
	// fmt.Printf("Russian country name: %v\n", record.Country.Names["ru"])
	// fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
	// fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
	// fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)

	// fmt.Printf("%v", record)

	// Output:
	// Portuguese (BR) city name: Londres
	// English subdivision name: England
	// Russian country name: Великобритания
	// ISO country code: GB
	// Time zone: Europe/London
	// Coordinates: 51.5142, -0.0931
}
