package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func getLoc(ip string) {
	type geolocation struct {
		CountryName string `json:"country_name"`
		State       string `json:"state_prov"`
		IP          string `json:"ip"`
		ISP         string `json:"isp"`
	}
	req, _ := http.NewRequest("GET", "https://api.ipgeolocation.io/ipgeo", nil)
	req.Header.Set("Referer", "https://ipgeolocation.io/")
	q := req.URL.Query()
	q.Add("include", "hostname")
	q.Add("ip", ip)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	geoInfo := &geolocation{}
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&geoInfo)
	fmt.Println("Country Name:\t", geoInfo.CountryName)
	fmt.Println("State:\t\t", geoInfo.State)
	fmt.Println("IP:\t\t", geoInfo.IP)
	fmt.Println("ISP:\t\t", geoInfo.ISP)
}

func main() {
	ip := os.Args[len(os.Args)-1]
	if os.Args[0] == ip {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("IP: ")
		ip, _ = reader.ReadString('\n')
	} else if len(os.Args) > 1 {
		ips := os.Args[1:]
		for _, i := range ips {
			getLoc(i)
			fmt.Println("")
		}
		os.Exit(0)
	}
	getLoc(ip)
}
