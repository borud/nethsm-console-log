PACKAGE_NAME=nethsm-console-log
PACKAGE_VERSION=0.1.0
BINARIES := $(notdir $(shell find cmd -mindepth 1 -maxdepth 1 -type d))

.PHONY: $(BINARIES)
.PHONY: all
.PHONY: build
.PHONY: vet
.PHONY: staticcheck
.PHONY: lint
.PHONY: clean
.PHONY: install-deps

all: vet lint staticcheck build
build: $(BINARIES)

$(BINARIES):
	@echo "*** building $@"
	@cd cmd/$@ && go build -o ../../bin/$@

vet:
	@echo "*** $@"
	@go vet ./...

staticcheck:
	@echo "*** $@"
	@staticcheck ./...

lint:
	@echo "*** $@"
	@revive ./...

clean:
	@rm -rf bin

install-deps:
	@go install github.com/mgechev/revive@latest
	@go install honnef.co/go/tools/cmd/staticcheck@latest
