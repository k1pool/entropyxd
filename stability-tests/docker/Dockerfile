ARG ENXPAD_IMAGE
ARG ENXPAMINER_IMAGE

FROM ${ENXPAD_IMAGE} as entropyxd
FROM ${ENXPAMINER_IMAGE} as entropyxminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/k1pool/entropyxd

WORKDIR /go/src/github.com/k1pool/entropyxd

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=entropyxd /app/ /app/
COPY --from=entropyxminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/k1pool/entropyxd/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]