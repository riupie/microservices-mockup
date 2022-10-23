FROM golang:1.18 as builder

WORKDIR /go/src/app
COPY . .

RUN go mod download && go vet -v && go test -v

RUN CGO_ENABLED=0 go build -o /go/bin/app/backend-go

FROM gcr.io/distroless/static-debian11

WORKDIR /app
COPY --from=builder /go/bin/app/backend-go /app/backend-go
EXPOSE 8080

CMD ["/app/backend-go"]
