# Initialize

Unzip `dsv-server.zip` and `vpn-client.zip`

So, directory structure should look like as below:
```
...
./dsv-server
./dsv-server/keys
./dsv-server/keys/truststore.jks
./dsv-server/dsv-server.jar
./dsv-server/lib/*
...
./vpn-client
./vpn-client/client-eimzo.conf
./vpn-client/truststore.jks
./vpn-client/vpn-client.jar
./vpn-client/lib/*
...
./go.mod
./main.go
./Dockerfile
...
```

Copy `client-xxxxxxx.yks` file into the directory `./vpn-client/` and make appropriate changes in the file `./vpn-client/client-eimzo.conf`

# Build docker image

```
docker build -t dsv-server-vpn-client .
```

# Run

```
docker run --rm -it -p 9091:9091 \
    --add-host e-imzo.uz:127.0.0.5 \
    -v $(pwd)/vpn-client/client-eimzo.conf:/opt/app/vpn-client/client-eimzo.conf \
    -v $(pwd)/vpn-client/client-xxxxxxx.yks:/opt/app/vpn-client/client-xxxxxxx.yks \
    -v $(pwd)/logs:/opt/app/logs \
    dsv-server-vpn-client -dsv-port 9091 -dsv-log /opt/app/logs/dsv.log -vpn-log /opt/app/logs/vpn.log
```

## Run as daemon

```
docker run -d -it -p 9091:9091 \
    --name dsv-server-vpn-client \
    --restart unless-stopped \
    --add-host e-imzo.uz:127.0.0.5 \
    -v $(pwd)/vpn-client/client-eimzo.conf:/opt/app/vpn-client/client-eimzo.conf \
    -v $(pwd)/vpn-client/client-xxxxxxx.yks:/opt/app/vpn-client/client-xxxxxxx.yks \
    -v $(pwd)/logs:/opt/app/logs \
    dsv-server-vpn-client -dsv-port 9091 -dsv-log /opt/app/logs/dsv.log -vpn-log /opt/app/logs/vpn.log
```