FROM golang:buster AS builder

WORKDIR /build

# Copy `go.mod` for definitions and `go.sum` to invalidate the next layer in case of a change in the dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w \
  -X 'go.jlucktay.dev/stack/pkg/version.builtBy=Docker' \
  -X 'go.jlucktay.dev/stack/pkg/version.date=$(TZ=UTC date '+%Y-%m-%dT%H:%M:%SZ')'"

# TODO(jlucktay): incorporate remaining ldflags?
#   -X 'go.jlucktay.dev/stack/pkg/version.version={{ .Version }}'
#   -X 'go.jlucktay.dev/stack/pkg/version.commit={{ .ShortCommit }}'

FROM scratch

VOLUME /workdir
WORKDIR /workdir

COPY --from=builder /build/stack /

ENTRYPOINT [ "/stack" ]
CMD [ "--help" ]
