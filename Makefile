VENODR_TOOL = dep
VENODR_CMD = $(if $(shell which $(VENODR_TOOL)),,$(error Please install $(VENODR_TOOL))) $(VENODR_TOOL)

default: build

build:  ## Build gaos
	go build -v ./cmd/gaos

test:  ## Test appleopensource package
	go test -v -race ./

vendor/list:  ## List vendor packages
	$(VENODR_CMD) status

vendor/clean:  ## clean vendor package files
	@find vendor -type f \( -name '*_test.go' -o -name '.gitignore' -o -name '*appveyor.yml' -o -name '.travis.yml' -o -name 'circle.yml' \) -print -exec rm -f {} ";"
	@find vendor -type d \( -name 'testdata' \) -print | xargs rm -rf

help:  ## Print this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[33m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

clean:
	$(RM) ./gaos
