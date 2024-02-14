VERSION ?= v0.1.x
DIST ?= development
HTML ?= ./snyth-js/html

.PHONY: clean
.PHONY: copy
.PHONY: build-all
.PHONY: release

all: test      \
     benchmark \
     coverage

clean:
	rm -rf dist/*

update:

build-all:
	cd snythd && make release

release: 
	@echo "... releasing snyth_$(VERSION)"
	cd snythd && make release VERSION=$(VERSION) DIST="snythd"
	rm -rf dist/$(DIST)
	mkdir -p dist/$(DIST)
	cp -r  ./snythd/dist/snythd dist/$(DIST)/
	cp -r  ./snyth-js/dist/html dist/$(DIST)/cloudflare
	cp -r  ./cloudflare/*       dist/$(DIST)/cloudflare
	tar --directory=dist --exclude=".DS_Store" -cvzf dist/$(DIST).tar.gz $(DIST)
	cd dist; zip --recurse-paths $(DIST).zip $(DIST)

debug:
	python3 -m http.server 9000 --directory ./dist/development/cloudflare
