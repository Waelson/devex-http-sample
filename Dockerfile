FROM golang:1.26-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY *.go ./
RUN CGO_ENABLED=0 go build -trimpath -o /out/http-sample .

FROM alpine:3.23
RUN addgroup -S app && adduser -S -G app app
COPY --from=build /out/http-sample /usr/local/bin/http-sample
USER app
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/http-sample"]
