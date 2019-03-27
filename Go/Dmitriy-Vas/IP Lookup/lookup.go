package IP_Lookup

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type IP struct {
	IP string
	HostName string
	City string
	Region string
	Country string
	Loc string
}

type Country struct {
	Name string
	Capital string
	Region string
	Population int
	NativeName string
}

func IPInfo(s string) IP {
	data := get("https://ipinfo.io/" + s + "/json")

	var ip IP
	err := json.Unmarshal(data, &ip)
	logError(err)

	return ip
}

func CountryInfo(s string) Country {
	ip := IPInfo(s)

	data := get("https://restcountries.eu/rest/v2/alpha/" + ip.Country)

	var country Country
	err := json.Unmarshal(data, &country)
	logError(err)

	return country
}

func logError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func get(s string) []byte {
	resp, err := http.Get(s)
	logError(err)

	data, err := ioutil.ReadAll(resp.Body)
	logError(err)
	resp.Body.Close()

	return data
}
