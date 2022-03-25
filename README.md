# Build docker image

docker build -t dsv-server-vpn-client .

# Run 
docker run --rm -it -p 9091:9091 \
    --add-host e-imzo.uz:127.0.0.5 \
    -v $(pwd)/dsv-server/logging.properties:/opt/app/dsv-server/logging.properties \
    -v $(pwd)/vpn-client/client-eimzo.conf:/opt/app/vpn-client/client-eimzo.conf \
    -v $(pwd)/vpn-client/client-yt.uz.yks:/opt/app/vpn-client/client-yt.uz.yks \
    -v $(pwd)/logs:/opt/app/logs \
    dsv-server-vpn-client -dsv-port 9091 -dsv-log-props /opt/app/dsv-server/logging.properties -dsv-log /opt/app/logs/dsv.log -vpn-log /opt/app/logs/vpn.log
