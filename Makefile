DIST      ?= development
SERIALNO  ?= 405419896
NEWDEVICE ?= 102030405
CARD      ?= 10058400
PIN       ?= 7531
DOOR      ?= 3
DEBUG     ?= --debug

.DEFAULT_GOAL := test
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
	go get -u github.com/uhppoted/uhppoted-lib@master

update-release:
	go get -u github.com/uhppoted/uhppote-core
	go get -u github.com/uhppoted/uhppoted-lib

format: 
	go fmt ./...

build: format
	mkdir -p bin
	go build -trimpath -o bin ./...

test: build
	go test ./...

benchmark: build
	go test -bench ./...

coverage: build
	go test -cover ./...

vet: build
	go vet ./...

lint: build
	env GOOS=darwin  GOARCH=amd64 staticcheck ./...
	env GOOS=linux   GOARCH=amd64 staticcheck ./...
	env GOOS=windows GOARCH=amd64 staticcheck ./...

vuln:
	govulncheck ./...

build-all: test vet lint
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
	cd dist; zip --recurse-paths $(DIST).zip $(DIST)

publish: release
	echo "Releasing version $(VERSION)"
	rm dist/development.tar.gz
	gh release create "$(VERSION)" "./dist/uhppote-simulator_$(VERSION).tar.gz"  "./dist/uhppote-simulator_$(VERSION).zip" --draft --prerelease --title "$(VERSION)-beta" --notes-file release-notes.md

debug: build
	go test -v ./simulator/UT0311L04/... -run TestCheckTimeProfile

delve: build
	dlv test github.com/uhppoted/uhppote-simulator/simulator/UT0311L04 -- run TestCheckTimeProfileInTimeSegmentWithOffset

godoc:
	godoc -http=:80	-index_interval=60s

version: build
	./bin/uhppote-simulator version

run: build
	# ./bin/uhppote-simulator --debug --bind 0.0.0.0:60000 --rest 0.0.0.0:8000 --devices "../runtime/simulation/devices"
	./bin/uhppote-simulator --bind 0.0.0.0:60000 --rest 0.0.0.0:8000 --devices "../runtime/simulation/devices"

new-device: build
	./bin/uhppote-simulator --debug --devices "../runtime/simulation/devices" new-device $(NEWDEVICE)

list-devices:
	curl -X GET "http://127.0.0.1:8000/uhppote/simulator" -H "accept: application/json"

create-device:
	curl -X POST "http://127.0.0.1:8000/uhppote/simulator" -H "accept: */*" -H "Content-Type: application/json" -d "{\"device-id\":$(NEWDEVICE),\"device-type\":\"UT0311-L04\",\"compressed\":false}"

delete-device:
	curl -X DELETE "http://127.0.0.1:8000/uhppote/simulator/$(NEWDEVICE)" -H "accept: */*"

swipe:
	curl -X POST "http://127.0.0.1:8000/uhppote/simulator/$(SERIALNO)/swipe" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"door\":$(DOOR),\"card-number\":$(CARD),\"direction\":1,\"PIN\":$(PIN)}"

swipe-in:
	curl -X POST "http://127.0.0.1:8000/uhppote/simulator/$(SERIALNO)/swipe" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"door\":$(DOOR),\"card-number\":$(CARD),\"direction\":1,\"PIN\":1234}"

swipe-out:
	curl -X POST "http://127.0.0.1:8000/uhppote/simulator/$(SERIALNO)/swipe" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"door\":$(DOOR),\"card-number\":$(CARD),\"direction\":2,\"PIN\":1234}"

open:
	# curl -X POST "http://127.0.0.1:8000/uhppote/simulator/$(SERIALNO)/door/$(DOOR)" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"action\":\"open\",\"duration\":10}"
	curl -X POST "http://127.0.0.1:8000/uhppote/simulator/$(SERIALNO)/door/$(DOOR)" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"action\":\"open\"}"

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


