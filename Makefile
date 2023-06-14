VERSION ?= v0.0.x
DIST ?= cloudflare
HTML ?= ./snyth-js/html

.PHONY: clean
.PHONY: copy
.PHONY: build-all

all: test      \
     benchmark \
     coverage

clean:
	rm -rf dist/*

update:

deploy: 
	cd ./snyth-js && make release
	rm -rf dist/cloudflare
	cp -r  ./snyth-js/dist/html dist/cloudflare
	cp -r  ./cloudflare/*       dist/cloudflare

