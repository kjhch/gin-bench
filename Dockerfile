FROM golang:1.18 AS builder

WORKDIR /home/admin

ENV GOPROXY=https://goproxy.io,direct

COPY . .

RUN make all

# FROM 基于 alpine:latest
FROM alpine:latest

# RUN 设置 Asia/Shanghai 时区
RUN echo -e http://mirrors.ustc.edu.cn/alpine/v3.13/main/ > /etc/apk/repositories && \
    apk --no-cache add tzdata  && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    adduser -D admin

ENV APP_ENV=test

USER admin

WORKDIR /home/admin

# COPY 源路径 目标路径 从镜像中 COPY
COPY --from=builder /home/admin/bin/* /bin/
COPY --from=builder /home/admin/configs/ .

CMD ["gin-bench"]