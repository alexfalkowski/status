FROM golang:1.23.5-bullseye AS build

ARG version=latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 go build -ldflags="-s -w -X 'github.com/alexfalkowski/status/cmd.Version=${version}' -extldflags '-static'" -tags netgo -a -o status main.go

FROM gcr.io/distroless/static

WORKDIR /

COPY --from=build /app/status /status
ENTRYPOINT ["/status"]
