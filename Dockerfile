FROM golang:alpine AS builder

RUN apk add build-base
COPY . /src
RUN cd /src/cmd/CheckList/ && CGO_ENABLED=0 go build -o /CheckList .


FROM scratch

WORKDIR /data/CheckList
WORKDIR /app

COPY --from=builder /CheckList /app/

ENTRYPOINT ["./CheckList"]