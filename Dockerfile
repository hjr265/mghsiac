# Build Stage
FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.21.0-alpine3.18 AS builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

# Build the Go application.
WORKDIR /mghsiac
ADD . .
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" .

# Run Stage
FROM --platform=${TARGETPLATFORM:-linux/amd64} alpine:3.18.3

RUN apk add --no-cache tzdata

# Copy the built Go application.
WORKDIR /mghsiac
COPY --from=builder /mghsiac/mghsiac /mghsiac/mghsiac

ENTRYPOINT ["./mghsiac"]
