package main

import (
	"net/http"
	"os"
	"fmt"
	"bufio"
	"net"
	"encoding/json"
	"strings"
)

type GeoIP struct {
	Ip          string  `json:"ip"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name""`
	RegionCode  string  `json:"region_code"`
	RegionName  string  `json:"region_name"`
	City        string  `json:"city"`
	Zipcode     string  `json:"zipcode"`
	Lat         float32 `json:"latitude"`
	Lon         float32 `json:"longitude"`
	MetroCode   int     `json:"metro_code"`
}

var (
	err      error
	geo      GeoIP
	response *http.Response
	ip		 string
)

func getGeoData(ip string) GeoIP {
	response, err = http.Get("https://freegeoip.net/json/" + ip)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	json.NewDecoder(response.Body).Decode(&geo)

	return geo
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your ip: ")
	ip, err := reader.ReadString('\n')
	ip = strings.TrimRight(ip, "\n")
	if nil != err {
		fmt.Println(err)
	}

	if nil == net.ParseIP(ip) {
		fmt.Println("Input incorrect IP-address, has been used you public IP")
		ip = ""
	}

	geo := getGeoData(ip)

	fmt.Println(ip, " based in ", geo.CountryName)
}
