## tiktok_demo

### How to run

#### 1. Local environment

[Download ffmpeg package](https://ffmpeg.org/download.html) && **add ffmpeg to system path or user path**
```shell
# Install other services
docker-compose up

go build -o offer_tiktok && ./offer_tiktok
```

#### 2 Docker environment
```shell
# Compile in docker image && Packaged with ffmpeg
docker build -t tiktok:latest -f ./docker-build/Dockerfile .

# Start all service
cd docker-build 
docker-compose up -d
```