# build executable
FROM golang:1.16 as builder

WORKDIR /go/src/journey
ADD . /go/src/journey

RUN make all

# run distroless
FROM gcr.io/distroless/base-debian10

EXPOSE 3000

WORKDIR /opt/journey

COPY --from=builder /go/src/journey/www .

ENTRYPOINT [ "/opt/journey/www" ]
