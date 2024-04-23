# Billing Mobile API

Dibuat menggunakan bahasa pemrograman Go. Digunakan untuk interkoneksi dengan aplikasi billing mobile.

## Development

```bash
compiledaemon --command="./go-billing-mobile-api"
```

## Deployment

```bash
docker build -t go-billing-mobile-api .
docker container create --name go-billing-mobile-api_1 -p 7720:7720 go-billing-mobile-api
docker start go-billing-mobile-api_1
```

Rebuild

```bash
sudo docker stop go-billing-mobile-api_1 && sudo docker rm go-billing-mobile-api_1 && sudo docker rmi go-billing-mobile-api && sudo docker build -t go-billing-mobile-api . && sudo docker container create --name go-billing-mobile-api_1 -p 7720:7720 go-billing-mobile-api && sudo docker start go-billing-mobile-api_1
```

## Lisensi

Hak Cipta dilindungi dan milik @atozpw
