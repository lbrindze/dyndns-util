package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
)

const (
	IPIFY_URL = "https://api.ipify.org?format=text"
	IDENT_URL = "http://ident.me"
	WIMIP_URL = "https://ipv4bot.whatismyipaddress.com"
)

var ip_services = []string{
	IPIFY_URL,
	//IDENT_URL,
	WIMIP_URL,
}

func ipNeedsUpdate(publicIp net.IP, host string) bool {
	ips, err := net.LookupIP(host)
	if err != nil {
		fmt.Println(fmt.Errorf("Could not find any ip for host %s", host))
		return true
	}
	fmt.Println(ips)
	for _, ip := range ips {
		if publicIp.Equal(ip) {
			return false
		}
	}
	return true
}

func getMyIp() (string, error) {
	// pick a random ip service and get ip, parse result.
	url := ip_services[rand.Intn(len(ip_services))]
	fmt.Printf("Using URL %s\n", url)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(ip), err
}
