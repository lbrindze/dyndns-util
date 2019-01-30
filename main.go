package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"math/rand"
	"time"
)

var config = r53Config{Weight: 1}

func init() {
	flag.StringVar(&config.Name, "d", "", "domain name")
	flag.StringVar(&config.ZoneId, "z", "", "AWS Zone Id for domain")
	flag.Int64Var(&config.TTL, "ttl", int64(60), "TTL for DNS Cache")
}

func main() {
	flag.Parse()
	if config.Name == "" || config.ZoneId == "" {
		fmt.Println(fmt.Errorf("Incomplete arguments: d: %s, t: %s, z: %s\n", config.Name, config.ZoneId))
		flag.PrintDefaults()
		return
	}

	rand.Seed(time.Now().UTC().UnixNano())
	ip, err := getMyIp()
	fmt.Println(ip)
	if err != nil {
		fmt.Println("Failed to get current Public IP Address:", err)
	}
	config.LoadTargetIp(ip)
	if ipNeedsUpdate(config.Target, config.Name) {

		sess, err := session.NewSession()
		if err != nil {
			fmt.Println("failed to create session,", err)
			return
		}
		svc := route53.New(sess)
		updateRecord(svc, &config)
	} else {
		fmt.Println("IP Address matches record, no changes made")
	}
}
