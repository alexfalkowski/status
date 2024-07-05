FROM golang:1.22.5-bullseye AS build

ARG version=latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 go build -ldflags="-s -w -X 'github.com/alexfalkowski/status/cmd.Version=${version}'" -a -o status main.go

FROM gcr.io/distroless/base-debian12

WORKDIR /

COPY --from=build /app/status /status
ENTRYPOINT ["/status"]
