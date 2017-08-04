package main

import "github.com/segmentio/go-route53"
import "github.com/mitchellh/goamz/aws"
import "encoding/json"
import "os"
import "flag"
import "github.com/olekukonko/tablewriter"


func check(err error) {
  if err != nil {
    panic(err)
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

  if *mode == "add" {
    if *ip == "" {
      panic("Mandatory parameter missing: IP")
    }
    if *domain == "" {
      panic("Mandatory parameter missing: domain")
    }

    res, err := dns.Zone("Z423WC4H4VGU4").Add("A", *domain, *ip)
    check(err)

    b, err := json.MarshalIndent(res, "", "  ")
    check(err)

    os.Stdout.Write(b)
  }

  if *mode == "list" {
    res, err := dns.Zone("Z423WC4H4VGU4").Records()
    check(err)

    b, err := json.MarshalIndent(res, "", "  ")
    check(err)

    os.Stdout.Write(b)
  }



}
