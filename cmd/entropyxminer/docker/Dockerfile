# -- multistage docker build: stage #1: build stage
FROM golang:1.23-alpine AS build

RUN mkdir -p /go/src/github.com/kaspanet/entropyxd

WORKDIR /go/src/github.com/kaspanet/entropyxd

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

WORKDIR /go/src/github.com/k1pool/entropyxd/cmd/entropyxminer
RUN GOOS=linux go build -a -installsuffix cgo -o entropyxminer .

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /app

RUN apk add --no-cache ca-certificates tini

COPY --from=build /go/src/github.com/k1pool/entropyxd/cmd/entropyxminer/entropyxminer /app/

USER nobody
ENTRYPOINT [ "/sbin/tini", "--" ]
