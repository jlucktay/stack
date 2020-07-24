FROM golang:buster AS builder

WORKDIR /build

# Copy `go.mod` for definitions and `go.sum` to invalidate the next layer in case of a change in the dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

COPY . .

RUN GOOS=linux go build -tags 'osusergo netgo'

FROM scratch

WORKDIR /

COPY --from=builder /build/stack .

ENTRYPOINT [ "/stack" ]
