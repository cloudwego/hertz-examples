## Offer_Tiktok

### How to run

#### 1. Local environment

[Download ffmpeg package](https://ffmpeg.org/download.html) && **add ffmpeg to system path or user path**
```shell
docker-compose up
go build -o offer_tiktok && ./offer_tiktok
```

#### 2 Docker environment
```shell
docker build -t tiktok:latest -f ./docker-build/Dockerfile .

cd docker-build 
docker-compose up -d
```