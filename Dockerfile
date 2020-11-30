# Build server
###############################################
FROM golang:1.14-alpine as server-build-env

WORKDIR /api

RUN apk update && apk add --no-cache gcc musl-dev git bash

COPY ./api .

RUN go mod download

RUN go build -ldflags '-w -s' -a -o ./bin/api ./cmd

# Build app
###############################################
FROM node:14.11-alpine3.12 as app-build-env

WORKDIR /app

RUN apk update && apk add --no-cache bash

COPY ./app .

RUN yarn

RUN yarn build

# Deploy
###############################################
FROM alpine:3.12
RUN apk update && apk add --no-cache bash

COPY --from=server-build-env /api/bin/api /srv/http/api
COPY --from=app-build-env /app/build /srv/www/app

EXPOSE 80 80

CMD ["/srv/http/api"]