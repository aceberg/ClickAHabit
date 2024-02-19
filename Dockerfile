FROM golang:alpine AS builder

RUN apk add build-base
COPY . /src
RUN cd /src/cmd/ClickAHabit/ && CGO_ENABLED=0 go build -o /ClickAHabit .


FROM scratch

WORKDIR /data/ClickAHabit
WORKDIR /app

COPY --from=builder /ClickAHabit /app/

ENTRYPOINT ["./ClickAHabit"]