#first stage - builder
FROM golang:1.11.0-stretch as builder
WORKDIR /dir
COPY app.go .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
#second stage
FROM scratch
ENV NEW_RELIC_INSIGHTS_INSERT_API_KEY='<NEW-RELIC-INSIGHTS-INSERT-API-KEY>'
WORKDIR /root/
COPY --from=builder /dir/app .
CMD ["./app"]