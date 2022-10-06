FROM --platform=linux/amd64 golang:1.18.1-alpine3.15 AS builder

WORKDIR "/app"

RUN apk add --no-cache git ca-certicates

ENV go env -W GO111MODULE=on

COPY . .

RUN go build ./... & go build

FROM --platform=linux/amd64 alpine

RUN apk update upgrade
RUN apk --no-cache add ca-certificates bash

WORKDIR /root/

COPY --from=builder /app .

RUN chmod +x go-delivery-service-kiwibot

EXPOSE $PORT

ENTRYPOINT ["./go-delivery-service-kiwibot"]