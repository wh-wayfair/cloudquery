# Test unit
.PHONY: test
test:
	go test -race -timeout 3m ./...

# Install tools
.PHONY: install-tools
install-tools:
	@echo Installing tools from tools/tool.go
	@cat tools/tool.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml

.PHONY: gen-mocks
gen-mocks: install-tools
	rm -rf ./client/mocks/*
	go generate ./client/...

.PHONY: gen-docs
gen-docs:
	rm -rf ../../../website/tables/slack
	go run main.go doc ../../../website/tables/slack
	sed 's_(\(.*\))_(../../../../../website/tables/slack/\1)_' ../../../website/tables/slack/README.md > ./docs/tables/README.md
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/slack/README.md
	mv ../../../website/tables/slack/README.md ../../../website/pages/docs/plugins/sources/slack/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/slack/*.md
	rm -rf ../../../website/tables/slack/*.bak

# All gen targets
.PHONY: gen
gen: gen-mocks gen-docs
