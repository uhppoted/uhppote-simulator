VERSION    = v0.5.1x
LDFLAGS    = -ldflags "-X uhppote.VERSION=$(VERSION)" 

SERIALNO  ?= 405419896
NEWDEVICE ?= 102030405
CARD      ?= 65538
DOOR      ?= 3
DEBUG     ?= --debug

all: test      \
	 benchmark \
     coverage

clean:
	go clean
	rm -rf bin

format: 
	go fmt ./...

build: format
	mkdir -p bin
	go build -o bin ./...

test: build
	go test ./...

vet: build
	go vet ./...

lint: build
	golint ./...

benchmark: build
	go test -bench ./...

coverage: build
	go test -cover ./...

debug: build
	go test ./...

release: test vet
	mkdir -p dist/$(DIST)/windows
	mkdir -p dist/$(DIST)/darwin
	mkdir -p dist/$(DIST)/linux
	mkdir -p dist/$(DIST)/arm7
	env GOOS=linux   GOARCH=amd64       go build -o dist/$(DIST)/linux/uhppote-cli       ./...
	env GOOS=linux   GOARCH=arm GOARM=7 go build -o dist/$(DIST)/arm7/uhppote-cli        ./...
	env GOOS=darwin  GOARCH=amd64       go build -o dist/$(DIST)/darwin/uhppote-cli      ./..
	env GOOS=windows GOARCH=amd64       go build -o dist/$(DIST)/windows/uhppote-cli.exe ./...

release-tar: release
	tar --directory=dist --exclude=".DS_Store" -cvzf dist/$(DIST).tar.gz $(DIST)

new-device: build
	./bin/uhppote-simulator --debug --devices "../runtime/simulation/devices" new-device $(NEWDEVICE)

list-devices:
	curl -X GET "http://127.0.0.1:8000/uhppote/simulator" -H "accept: application/json"

create-device:
	curl -X POST "http://127.0.0.1:8000/uhppote/simulator" -H "accept: */*" -H "Content-Type: application/json" -d "{\"device-id\":$(NEWDEVICE),\"device-type\":\"UT0311-L04\",\"compressed\":false}"

delete-device:
	curl -X DELETE "http://127.0.0.1:8000/uhppote/simulator/$(NEWDEVICE)" -H "accept: */*"

swipe:
	curl -X POST "http://127.0.0.1:8000/uhppote/simulator/$(SERIALNO)/swipe" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"door\":$(DOOR),\"card-number\":$(CARD)}"
