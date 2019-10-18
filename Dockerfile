FROM golang:1.12.7-alpine3.10 as build-step

RUN apk add --update --no-cache ca-certificates git make

WORKDIR /go-app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN make devel-deps
RUN make bin/app

#runtime image
FROM alpine
COPY --from=build-step /go-app/bin/app /app
ENTRYPOINT ["./app"]

EXPOSE 8080
