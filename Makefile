dist : clean deps build-static docker

clean:
	rm route53er
deps:
	dep ensure

build-static: 
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' .

build:
	go build .

docker:
	docker build . -t deathowl/route53er
	
.PHONY : dist
