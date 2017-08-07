Manage route53 hosted zones with a single command.
Required: 
	go, dep, optionally docker

run in container:
* docker run -e ZONE=your_zone_id -e AWS_ACCESS_KEY_ID=ACCESSKEY -e AWS_SECRET_ACCESS_KEY=SECRETACCESKEY -v '/etc/ssl/certs/ca-certificates.crt:/etc/ssl/certs/ca-certificates.crt' deathowl/route53er /route53er

run on your host:
* dep ensure
* go build .
* ZONE=ZONE_ID AWS_ACCESS_KEY_ID=ACCES_KEY_ID AWS_SECRET_ACCESS_KEY=ACCESS_KEY ./route53er --operation=list

Supported operations:
* list
* add : required parameters --ip:Target for the A record --domain: domain
* delete: required parameters: --domain