FROM golang:1.16 AS builder

# Set some shell options for using pipes and such.
SHELL [ "/bin/bash", "-euo", "pipefail", "-c" ]

# Install/update the common CA certificates package now, and blag it later.
RUN apt-get update \
  && DEBIAN_FRONTEND=noninteractive apt-get install --assume-yes --no-install-recommends ca-certificates \
  && apt-get autoremove --assume-yes \
  && apt-get clean \
  && rm --force --recursive /root/.cache \
  && rm --force --recursive /var/lib/apt/lists/*

# Don't call any C code; the 'scratch' base image used later won't have any libraries to reference.
ENV CGO_ENABLED=0

# Use Go modules
ENV GO111MODULE=on

WORKDIR /go/src/go.jlucktay.dev/stack

# Add the sources.
COPY . .

# Compile! With the '--mount' flags below, Go's build cache is kept between builds.
# https://github.com/golang/go/issues/27719#issuecomment-514747274
RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  go build -ldflags="-X 'go.jlucktay.dev/version.builtBy=Docker'" -trimpath -v -o /bin/stack

FROM scratch AS runner

# Bring common CA certificates and binary over.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /bin/stack /bin/stack

VOLUME /workdir
WORKDIR /workdir

ENTRYPOINT [ "/bin/stack" ]
CMD [ "--help" ]
