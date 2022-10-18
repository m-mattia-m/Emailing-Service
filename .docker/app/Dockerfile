FROM golang:1.19-alpine AS builder

WORKDIR /build

COPY . /build/

RUN go get

ENV CGO_ENABLED=0
RUN go build -o dist/emailing-service

# https://github.com/GoogleContainerTools/distroless/blob/main/base/README.md
FROM gcr.io/distroless/static-debian11 AS image

ARG USER=nonroot

USER ${USER}

WORKDIR /app

COPY --from=builder --chown=${USER}:${USER} /build/dist/emailing-service /app/emailing-service

ENTRYPOINT ["/app/emailing-service"]
