############################################################
# Backend Build
############################################################
FROM golang:1.19-alpine as builder

ARG BUILD_TIMESTAMP
ARG VERSION

WORKDIR /app
COPY . .
RUN GOPROXY="https://goproxy.cn" go mod download

# Copy frontend build into embed directory so that we can embed
# the SPA into the Go binary via go embed.


RUN CGO_ENABLED=0 go build -o gin-rest-api-linux-x64 main.go
# Compiled backend binary is in '/app/bin/' named 'console'


############################################################
# Final Image
############################################################
FROM alpine:3.16

WORKDIR /app
COPY --from=builder /app/gin-rest-api-linux-x64 /app/gin-rest-api-linux-x64
COPY --from=builder /app/app.ini /app/app.ini

RUN  set -x \
    && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk add -U tzdata curl netcat-openbsd && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && apk del tzdata \
    && chmod +x /app/gin-rest-api-linux-x64

CMD ["/app/gin-rest-api-linux-x64"]
