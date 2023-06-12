VERSION ?= v0.0.x
DIST ?= development
HTML ?= ../snyth-js/html

.PHONY: clean
.PHONY: copy
.PHONY: build-all

all: test      \
     benchmark \
     coverage

clean:
	rm -rf bin/*

update:
	go mod tidy
#	go mod vendor

update-release:
	go mod tidy
#	go mod vendor

copy: 
	cd ../../snyth-development && make release
	rm -rf html/css
	rm -rf html/fonts
	rm -rf html/images
	rm -rf html/javascript
	rm -rf html/midi
	rm -rf html/favicon.ico
	rm -rf html/index.html
	cp -r  ../../snyth-development/dist/html .

format: 
	go fmt ./...

build: format
	mkdir -p bin
#	go build -mod readonly --trimpath -ldflags="-X 'main.version=$(VERSION)'" -o bin ./...
	go build -mod readonly --trimpath -ldflags="-X 'main.version=$(VERSION)'" -o bin ./...

test: build
	go test ./...

benchmark: test
	go test -bench ./...

coverage: build
	go test -cover ./...

vet: build
	go vet ./...

build-all: 
	mkdir -p bin
	go fmt ./...
	go build --trimpath -ldflags="-X 'main.version=$(VERSION)'" -o bin ./...
	go test ./...
	go vet ./...

	mkdir -p dist/$(DIST)/windows
	mkdir -p dist/$(DIST)/darwin
	mkdir -p dist/$(DIST)/linux

	env GOOS=linux   GOARCH=amd64 GOWORK=off go build --trimpath -o dist/$(DIST)/linux   ./...
	env GOOS=darwin  GOARCH=amd64 GOWORK=off go build --trimpath -o dist/$(DIST)/darwin  ./...
	env GOOS=windows GOARCH=amd64 GOWORK=off go build --trimpath -o dist/$(DIST)/windows ./...

release: copy build-all

debug: build
	./bin/snythd --debug --html $(HTML)

run: copy build
	./bin/snythd 

deploy: copy build
	gcloud app deploy app.yaml

