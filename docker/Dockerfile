# -- multistage docker build: stage #1: build stage
FROM golang:1.23-alpine AS build

RUN mkdir -p /go/src/github.com/entropyxnet/entropyxd

WORKDIR /go/src/github.com/entropyxnet/entropyxd

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .


# Cache entropyxd dependencies
RUN go mod download

COPY . .

RUN go build $FLAGS -o entropyxd .

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /app

RUN apk add --no-cache ca-certificates tini

COPY --from=build /go/src/github.com/k1pool/entropyxd/entropyxd /app/
COPY --from=build /go/src/github.com/k1pool/entropyxd/infrastructure/config/sample-entropyxd.conf /app/

USER nobody
ENTRYPOINT [ "/sbin/tini", "--" ]
