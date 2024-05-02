DIST      ?= development
SERIALNO  ?= 405419896
NEWDEVICE ?= 102030405
CARD      ?= 10058400
PIN       ?= 7531
PASSCODE  ?= 654321
DOOR      ?= 3
DEBUG     ?= --debug
DOCKER    ?= ghcr.io/uhppoted/simulator:latest

WORKDIR=/Users/tonyseebregts/Development/uhppote/uhppoted/uhppote-simulator/workdir

.DEFAULT_GOAL := test
.PHONY: clean
.PHONY: update
.PHONY: update-release
.PHONY: docker-run

all: test      \
	 benchmark \
     coverage

clean:
	go clean
	rm -rf bin

update:
	go get -u github.com/uhppoted/uhppote-core@master
	go get -u github.com/uhppoted/uhppoted-lib@main

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
	rm -f dist/development.tar.gz
	gh release create "$(VERSION)" "./dist/uhppote-simulator_$(VERSION).tar.gz"  "./dist/uhppote-simulator_$(VERSION).zip" --draft --prerelease --title "$(VERSION)-beta" --notes-file release-notes.md

debug:
	# curl -X POST "http://127.0.0.1:8765/uhppote/simulator/706050403/swipe" -H "accept: application/json" -H "Content-Type: application/json" -d '{"door":1, "card-number":10058400,"direction":1}'
	go test -v ./...

delve: build
#	dlv test github.com/uhppoted/uhppote-simulator/simulator/UT0311L04 -- run TestCheckTimeProfileInTimeSegmentWithOffset
	dlv debug github.com/uhppoted/uhppote-simulator/cmd/uhppote-simulator -- --bind 0.0.0.0:60000 --rest 0.0.0.0:8000 --devices "../runtime/simulation/devices"

godoc:
	godoc -http=:80	-index_interval=60s

version: build
	./bin/uhppote-simulator version

run: build
	./bin/uhppote-simulator --debug --bind 0.0.0.0:60000 --rest 0.0.0.0:8000 --devices "../runtime/simulation/devices"

list-devices:
	curl -X GET "http://127.0.0.1:8000/uhppote/simulator" -H "accept: application/json"

create-device:
	curl -X POST "http://127.0.0.1:8000/uhppote/simulator" -H "accept: */*" -H "Content-Type: application/json" -d '{"device-id":$(NEWDEVICE),"device-type":"UT0311-L04","compressed":false}'

delete-device:
	curl -X DELETE "http://127.0.0.1:8000/uhppote/simulator/$(NEWDEVICE)" -H "accept: */*"

swipe:
	curl -X POST "http://127.0.0.1:8000/uhppote/simulator/$(SERIALNO)/swipe" -H "accept: application/json" -H "Content-Type: application/json" -d '{"door":$(DOOR),"card-number":$(CARD),"direction":1,"PIN":$(PIN)}'

swipe-in:
	curl -X POST "http://127.0.0.1:8000/uhppote/simulator/$(SERIALNO)/swipe" -H "accept: application/json" -H "Content-Type: application/json" -d '{"door":$(DOOR),"card-number":$(CARD),"direction":1,"PIN":1234}'

swipe-out:
	curl -X POST "http://127.0.0.1:8000/uhppote/simulator/$(SERIALNO)/swipe" -H "accept: application/json" -H "Content-Type: application/json" -d '{"door":$(DOOR),"card-number":$(CARD),"direction":2,"PIN":1234}'

passcode:
	curl -X POST "http://127.0.0.1:8000/uhppote/simulator/$(SERIALNO)/code" -H "accept: application/json" -H "Content-Type: application/json" -d '{"door":$(DOOR),"passcode":$(PASSCODE)}'

open:
	curl -X POST "http://127.0.0.1:8000/uhppote/simulator/$(SERIALNO)/door/$(DOOR)" -H "accept: application/json" -H "Content-Type: application/json" -d '{"action":"open"}'

close:
	curl -X POST "http://127.0.0.1:8000/uhppote/simulator/$(SERIALNO)/door/$(DOOR)" -H "accept: application/json" -H "Content-Type: application/json" -d '{"action":"close"}'

button:
	curl -X POST "http://127.0.0.1:8000/uhppote/simulator/$(SERIALNO)/door/$(DOOR)" -H "accept: application/json" -H "Content-Type: application/json" -d '{"action":"button", "duration":10}'

# v06.62 events
v6.62-swipe:
	curl -X POST "http://127.0.0.1:8000/uhppote/simulator/201020304/swipe" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"door\":$(DOOR),\"card-number\":$(CARD)}"

v6.62-open:
	curl -X POST "http://127.0.0.1:8000/uhppote/simulator/201020304/door/1" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"action\":\"open\",\"duration\":10}"

v6.62-close:
	curl -X POST "http://127.0.0.1:8000/uhppote/simulator/201020304/door/1" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"action\":\"close\"}"

v6.62-button:
	curl -X POST "http://127.0.0.1:8000/uhppote/simulator/201020304/door/1" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"action\":\"button\", \"duration\":10}"

swagger: 
	docker run --detach --publish 80:8080 --rm swaggerapi/swagger-editor 
	open http://127.0.0.1:80

rest-swipe:
	python3 scripts/REST.py swipe     --controller 405419896 --door 1 --card 10058400

rest-swipe-in:
	python3 scripts/REST.py swipe-in  --controller 405419896 --door 1 --card 10058400

rest-swipe-out:
	python3 scripts/REST.py swipe-out --controller 405419896 --door 1 --card 10058400

rest-passcode:
	python3 scripts/REST.py passcode  --controller 405419896 --door 1 --code 13571

rest-button:
	python3 scripts/REST.py button    --controller 405419896 --door 1 --duration 30

rest-open:
	python3 scripts/REST.py open      --controller 405419896 --door 1

rest-close:
	python3 scripts/REST.py close     --controller 405419896 --door 1

rest-create-controller:
	python3 scripts/REST.py create-controller --controller 123456789 --type UT0311-L04 --compressed false
	python3 scripts/REST.py list-controllers

rest-delete-controller:
	python3 scripts/REST.py delete-controller --controller 123456789
	python3 scripts/REST.py list-controllers

rest-list-controllers:
	python3 scripts/REST.py list-controllers

docker: docker-dev docker-ghcr
	cd docker && find . -name .DS_Store -delete && rm -f compose.zip && zip --recurse-paths compose.zip compose

docker-dev: build
	rm -rf dist/docker/dev/*
	mkdir -p dist/docker/dev
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -o dist/docker/dev ./...
	cp docker/dev/Dockerfile     dist/docker/dev
	cp docker/dev/405419896.json dist/docker/dev
	cp docker/dev/303986753.json dist/docker/dev
	cp docker/dev/201020304.json dist/docker/dev
	cd dist/docker/dev && docker build --no-cache -f Dockerfile -t uhppoted/simulator-dev .

docker-ghcr: build
	rm -rf dist/docker/ghcr/*
	mkdir -p dist/docker/ghcr
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -o dist/docker/ghcr ./...
	cp docker/ghcr/Dockerfile     dist/docker/ghcr
	cp docker/ghcr/405419896.json dist/docker/ghcr
	cd dist/docker/ghcr && docker build --no-cache -f Dockerfile -t $(DOCKER) .

docker-run-dev:
	docker run --publish 8000:8000 --publish 60000:60000/udp --name simulator --rm uhppoted/simulator-dev

docker-run-ghcr:
	docker run --publish 8000:8000 --publish 60000:60000/udp --name simulator \
	           --mount source=uhppoted-simulator,target=/usr/local/etc/uhppoted \
	           --rm ghcr.io/uhppoted/simulator

docker-compose:
	cd docker/compose && docker compose up

docker-clean:
	docker image     prune -f
	docker container prune -f
