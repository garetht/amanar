FROM golang:1.13

WORKDIR /app

RUN go get -u github.com/kevinburke/go-bindata/...

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN make install

ENTRYPOINT ["amanar"]
