# One image per Aether SD-Core 6.1.2 NF, built from the Aether_6.1.2_TYcustom repo.
#
# Build context = repo root (not NFs/<nf> like the official Dockerfiles), so a
# go.mod "replace => ../../shared-libs/<lib>" resolves inside the build.
# Runtime layout matches the official images: the binary sits in /usr/local/bin
# and bash is installed; the chart's run script calls "<bin> -cfg ...".
#
#   docker build -f Dockerfile.nf --build-arg NF=amf --build-arg BIN=amf -t <img> <repo root>

FROM golang:1.27.1-bookworm AS builder
ARG NF
ARG BIN
WORKDIR /src
COPY NFs/${NF} NFs/${NF}
COPY shared-libs shared-libs
WORKDIR /src/NFs/${NF}
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 go build -o /out/${BIN} .

FROM alpine:3.24
ARG BIN
RUN apk add --no-cache bash
COPY --from=builder /out/${BIN} /usr/local/bin/${BIN}
