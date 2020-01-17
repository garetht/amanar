FROM golang:1.13 as build

WORKDIR /app

RUN go get -u github.com/kevinburke/go-bindata/...

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN make docker-install

FROM scratch

COPY --from=build /bin/amanar /bin/amanar

ENTRYPOINT ["/bin/amanar"]
