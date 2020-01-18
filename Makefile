commit := $$(git rev-list -1 HEAD)
build_date := $$(date)
tag := $$(git describe --tags)
main := cmd/amanar.go

ld_flags := -X 'main.GitCommit=$(commit)' -X 'main.BuildDate=$(build_date)' -X 'main.GitTag=$(tag)'


PHONY: generate
generate:
	go generate

PHONY: install
install: CGO_ENABLED := 1
install: generate
	go install -ldflags "$(ld_flags)" $(main)

PHONY: build
build: CGO_ENABLED := 1
build: generate
	go build -ldflags "$(ld_flags)" $(main)

PHONY: test
test:
	go test -v

PHONY: docker-test
docker-test: CGO_ENABLED := 0
docker-test:
	go test -v

PHONY: docker-install
docker-install: CGO_ENABLED := 0
docker-install: GOOS := linux
docker-install:
	go build -ldflags "$(ld_flags) -w -s" -a -o /bin/amanar $(main)
