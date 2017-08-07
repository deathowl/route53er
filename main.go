package main

import "github.com/segmentio/go-route53"
import "github.com/mitchellh/goamz/aws"
import "os"
import "flag"
import "github.com/olekukonko/tablewriter"
import log "github.com/sirupsen/logrus"

var ZoneId string

func init() {
	ZoneId = os.Getenv("ZONE")
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	var ip = flag.String("ip", "", "IP")
	var domain = flag.String("domain", "", "domain")
	var mode = flag.String("operation", "list", "Possible operations: list, add, delete")

	auth, err := aws.EnvAuth()
	check(err)
	flag.Parse()

	dns := route53.New(auth, aws.USWest2)
	switch *mode {
	case "add":
		{
			if *ip == "" {
				panic("Mandatory parameter missing: IP")
			}
			if *domain == "" {
				panic("Mandatory parameter missing: domain")
			}

			res, err := dns.Zone(ZoneId).Add("A", *domain, *ip)
			check(err)
			if res.ChangeInfo.Status == "PENDING" {
				log.Infoln("Record :" + *domain + " added")
			}

		}
	case "list":
		{
			res, err := dns.Zone(ZoneId).Records()
			check(err)
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Name", "Type", "Target"})
			for _, Record := range res {
				if Record.Type == "A" {
					table.Append([]string{Record.Name, Record.Type, Record.Records[0]})
				}
			}
			table.Render()
		}
	case "delete":
		{
			if *domain == "" {
				panic("Mandatory parameter missing: domain")
			}
			Records, err := dns.Zone(ZoneId).RecordsByName(*domain)
			check(err)
			if len(Records) < 1 {
				log.Fatalln("No record found with name: " + *domain)
			}
			Deleted, err := dns.Zone("Z423WC4H4VGU4").Remove(Records[0].Type, Records[0].Name, Records[0].Records[0])
			if Deleted.ChangeInfo.Status == "PENDING" {
				log.Infoln("Record :" + *domain + " deleted")
			}
		}
	}

}
