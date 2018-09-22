FROM golang:1.11 AS build
WORKDIR /go/src/healthz
COPY . .
RUN go get -d -v \
    && CGO_ENABLED=0 GOOS=linux go build -v

FROM scratch
COPY --from=build /go/src/healthz/healthz /healthz
EXPOSE 8080
CMD [ "/healthz" ]
