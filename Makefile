# Point at your own clone of IBM-Cloud/terraform-provider-ibm, checked out at a
# reviewed release tag. Avoid shared, world-writable locations such as /tmp.
PROVIDER_SRC ?= $(HOME)/src/terraform-provider-ibm
PACKAGES ?=
PLUGIN_NAME = tflint-ruleset-ibm

.PHONY: default all extract generate fmt vet test build install regen clean extract-all generate-all coverage lint overlay hygiene

default: fmt vet test build

all: extract-all overlay generate-all fmt vet coverage test build

extract:
	go run ./extractor -provider $(PROVIDER_SRC) -out extractor/output $(if $(PACKAGES),-packages $(PACKAGES),)

extract-all:
	go run ./extractor -provider $(PROVIDER_SRC) -out extractor/output

generate:
	go run ./generator -in extractor/output -rules rules -docs docs/rules $(if $(PACKAGES),-packages $(PACKAGES),)

generate-all:
	go run ./generator -in extractor/output -rules rules -docs docs/rules

overlay:
	python3 scripts/openapi-overlay.py

regen: extract-all overlay generate-all fmt coverage

fmt:
	gofmt -w .

vet:
	go vet ./...

lint: vet hygiene
	@test -z "$$(gofmt -l .)" || (echo "gofmt needed:" && gofmt -l . && exit 1)

hygiene:
	@owner=$$(printf '%s%s' ramuk lawjju); \
	if grep -RIn --exclude-dir=.git --exclude-dir=dist --exclude-dir=extractor/output $$owner .; then \
		echo "personal GitHub owner found; use github.com/IBM/tflint-ruleset-ibm" >&2; \
		exit 1; \
	fi

coverage:
	go run ./coverage -provider $(PROVIDER_SRC) -root .

test:
	go test ./...

build:
	go build -o $(PLUGIN_NAME) .

install: build
	@echo "installing $(PLUGIN_NAME) to $(HOME)/.tflint.d/plugins (replaces any existing copy)"
	install -d -m 0755 $(HOME)/.tflint.d/plugins
	install -m 0755 ./$(PLUGIN_NAME) $(HOME)/.tflint.d/plugins/$(PLUGIN_NAME)

clean:
	rm -f $(PLUGIN_NAME)
