FROM golang:1.18 as builder
RUN mkdir -p /dsv-service-launcher
WORKDIR /dsv-service-launcher
COPY go.mod /dsv-service-launcher/
COPY main.go /dsv-service-launcher/
RUN go mod tidy
RUN go build -o service

FROM openjdk:8u322-oraclelinux8
RUN mkdir -p /opt/app/dsv-server && mkdir -p /opt/app/vpn-client
COPY --from=builder /dsv-service-launcher/service /opt/app/
COPY /dsv-server/dsv-server.jar /opt/app/dsv-server/
COPY /dsv-server/lib /opt/app/dsv-server/lib
COPY /dsv-server/keys /opt/app/dsv-server/keys
COPY /vpn-client/vpn-client.jar /opt/app/vpn-client/
COPY /vpn-client/truststore.jks /opt/app/vpn-client/
COPY /vpn-client/lib /opt/app/vpn-client/lib
WORKDIR /opt/app

ENTRYPOINT [ "/opt/app/service" ]