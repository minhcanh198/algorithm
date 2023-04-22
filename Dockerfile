FROM golang:1.18-buster as builder

RUN go version

WORKDIR /code

# add these two lines
ADD go.mod go.sum /code/
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /usr/bin/app cmd/main.go
COPY .env /usr/bin/


FROM golang:1.18-buster
RUN rm /etc/localtime
RUN ln -s /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime

WORKDIR /

COPY --from=builder /usr/bin/.env /

COPY --from=builder /usr/bin/app /usr/bin/app
ENTRYPOINT ["/usr/bin/app"]
