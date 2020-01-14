commit := $$(git rev-list -1 HEAD)
build_date := $$(date)

ld_flags := -X 'main.GitCommit=$(commit)' -X 'main.BuildDate=$(build_date)'

export CGO_ENABLED := 1

PHONY: generate
generate:
	go generate

PHONY: install
install: generate
	go install -ldflags "$(ld_flags)"

PHONY: build
build: generate
	go build -ldflags "$(ld_flags)"

PHONY: test
test:
	go test -v
