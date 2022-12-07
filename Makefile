DIST      ?= development
SERIALNO  ?= 405419896
NEWDEVICE ?= 102030405
CARD      ?= 8165538
DOOR      ?= 3
DEBUG     ?= --debug

.PHONY: clean
.PHONY: update
.PHONY: update-release

all: test      \
	 benchmark \
     coverage

clean:
	go clean
	rm -rf bin

update:
	go get -u github.com/uhppoted/uhppote-core@master

update-release:
	go get -u github.com/uhppoted/uhppote-core

format: 
	go fmt ./...

build: format
	mkdir -p bin
	go build -trimpath -o bin ./...

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

build-all: test vet
	mkdir -p dist/$(DIST)/windows
	mkdir -p dist/$(DIST)/darwin
	mkdir -p dist/$(DIST)/linux
	mkdir -p dist/$(DIST)/arm
	mkdir -p dist/$(DIST)/arm7
	env GOOS=linux   GOARCH=amd64         GOWORK=off go build -trimpath -o dist/$(DIST)/linux   ./...
	env GOOS=linux   GOARCH=arm64         GOWORK=off go build -trimpath -o dist/$(DIST)/arm     ./...
	env GOOS=linux   GOARCH=arm   GOARM=7 GOWORK=off go build -trimpath -o dist/$(DIST)/arm7    ./...
	env GOOS=darwin  GOARCH=amd64         GOWORK=off go build -trimpath -o dist/$(DIST)/darwin  ./...
	env GOOS=windows GOARCH=amd64         GOWORK=off go build -trimpath -o dist/$(DIST)/windows ./...

release: update-release build-all
	find . -name ".DS_Store" -delete
	tar --directory=dist --exclude=".DS_Store" -cvzf dist/$(DIST).tar.gz $(DIST)

debug: build
	go test -v ./simulator/UT0311L04/... -run TestCheckTimeProfile

delve: build
	dlv test github.com/uhppoted/uhppote-simulator/simulator/UT0311L04 -- run TestCheckTimeProfileInTimeSegmentWithOffset

godoc:
	godoc -http=:80	-index_interval=60s

version: build
	./bin/uhppote-simulator version

run: build
	./bin/uhppote-simulator --debug --bind 0.0.0.0:60000 --rest 0.0.0.0:8000 --devices "../runtime/simulation/devices"

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

open:
	curl -X POST "http://127.0.0.1:8000/uhppote/simulator/$(SERIALNO)/door/$(DOOR)" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"action\":\"open\",\"duration\":10}"

close:
	curl -X POST "http://127.0.0.1:8000/uhppote/simulator/$(SERIALNO)/door/$(DOOR)" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"action\":\"close\"}"

button:
	curl -X POST "http://127.0.0.1:8000/uhppote/simulator/$(SERIALNO)/door/$(DOOR)" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"action\":\"button\", \"duration\":10}"

# v06.62 events
v6.62-swipe:
	curl -X POST "http://127.0.0.1:8000/uhppote/simulator/201020304/swipe" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"door\":$(DOOR),\"card-number\":$(CARD)}"

v6.62-open:
	curl -X POST "http://127.0.0.1:8000/uhppote/simulator/201020304/door/1" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"action\":\"open\",\"duration\":10}"

v6.62-close:
	curl -X POST "http://127.0.0.1:8000/uhppote/simulator/201020304/door/1" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"action\":\"close\"}"

v6.62-button:
	curl -X POST "http://127.0.0.1:8000/uhppote/simulator/201020304/door/1" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"action\":\"button\", \"duration\":10}"


