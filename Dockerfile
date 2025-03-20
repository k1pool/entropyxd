# -- multistage docker build: stage #1: build stage
FROM golang:1.22.0-alpine AS build

RUN mkdir -p /go/src/github.com/k1pool/entropyxd

WORKDIR /go/src/github.com/k1pool/entropyxd

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .


# Cache entropyxd dependencies
RUN go mod download

COPY . .
RUN mkdir -p /entropyx/bin/
RUN go build $FLAGS -o /entropyx/bin/ ./cmd/...

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /root/

RUN apk add --no-cache ca-certificates tini

COPY --from=build /entropyx/bin/* /usr/bin/

ENTRYPOINT [ "/usr/bin/entropyxd" ]
