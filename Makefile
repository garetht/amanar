commit := $$(git rev-list -1 HEAD)
build_date := $$(date)

ld_flags := -X 'main.GitCommit=$(commit)' -X 'main.BuildDate=$(build_date)'

export CGO_ENABLED := 1

PHONY: install
install:
	go install -ldflags "$(ld_flags)"

PHONY: build
build:
	go build -ldflags "$(ld_flags)"
