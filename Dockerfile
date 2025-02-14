# syntax=docker/dockerfile:1

ARG ALPINE_VERSION=3.21

FROM --platform=$BUILDPLATFORM golang:1.24-alpine AS base
WORKDIR /src
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=bind,source=go.mod,target=go.mod \
    --mount=type=bind,source=go.sum,target=go.sum \
    go mod download

FROM base AS build
ARG TARGETOS
ARG TARGETARCH
ENV GOOS=$TARGETOS
ENV GOARCH=$TARGETARCH
ARG VERSION="0.0.0+unknown"
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=bind,target=. \
    go build -ldflags="-X 'main.Version=${VERSION}'" -o /bin/app ./cmd 

FROM base AS test
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=bind,target=. \
    go test -v -coverprofile=/tmp/coverage.txt ./... > /tmp/result.txt; \
    [[ $? -eq 0 ]] || { cat /tmp/result.txt; exit 1; }

FROM scratch AS export-test
COPY --from=test /tmp/coverage.txt /
COPY --from=test /tmp/result.txt /

FROM alpine:${ALPINE_VERSION} AS image
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app"]

FROM scratch AS binaries
COPY --from=build /bin/app /
