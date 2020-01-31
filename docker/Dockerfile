# build用兼開発用
FROM golang:1.13-alpine as build

WORKDIR /go/app

COPY . .

RUN apk add --no-cache git bash && \
 go build -o app && \
 go get gopkg.in/urfave/cli.v2@master && \
 go get github.com/oxequa/realize

# 本番用
FROM alpine

ENV APP_HOME /app
WORKDIR $APP_HOME

COPY --from=build /go/app .

RUN addgroup go \
  && adduser -D -G go go \
  && chown -R go:go $APP_HOME

CMD ["./app"]
