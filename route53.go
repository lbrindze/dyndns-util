package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
	"net"
	"time"
)

type AwsDnsConfig interface {
	LoadTargetIp(string)
}

type r53Config struct {
	Name   string
	ZoneId string
	TTL    int64
	Weight int64
	Target net.IP
}

func (r *r53Config) LoadTargetIp(ip string) {
	r.Target = net.ParseIP(ip)
}

func updateRecord(svc *route53.Route53, config *r53Config) {
	commentString := fmt.Sprintf("DynDNS Update on %s.", time.Now().Format("2006-01-02 3:4:5"))

	params := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{ // Required
			Changes: []*route53.Change{ // Required
				{ // Required
					Action: aws.String("UPSERT"), // Required
					ResourceRecordSet: &route53.ResourceRecordSet{ // Required
						Name: aws.String(config.Name), // Required
						Type: aws.String("A"),         // Required
						ResourceRecords: []*route53.ResourceRecord{
							{ // Required
								Value: aws.String(config.Target.String()), // Required
							},
						},
						TTL: aws.Int64(config.TTL),
					},
				},
			},
			Comment: aws.String(commentString),
		},
		HostedZoneId: aws.String(config.ZoneId), // Required
	}
	resp, err := svc.ChangeResourceRecordSets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println("Change Response:")
	fmt.Println(resp)
}
