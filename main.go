package main

import (
	"log"

	"ads_server/internal/ads"
	"github.com/oschwald/geoip2-golang"
)

func main() {
	geoIp, err := geoip2.Open("GeoLite2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	server := ads.NewServer(geoIp)

	log.Println("Starting server")
	log.Fatal(server.Listen())
}
