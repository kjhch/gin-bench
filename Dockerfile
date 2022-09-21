FROM golang:1.18 AS builder

WORKDIR /home/admin

ENV GOPROXY=https://goproxy.io,direct

COPY . .

RUN make all

# FROM 基础镜像
FROM centos:7

# RUN 设置 Asia/Shanghai 时区
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    yum install -y net-tools && \
    adduser -m admin

ARG APP_ENV

ENV APP_ENV=${APP_ENV:-prod}
ENV GIN_MODE=release

#USER admin

WORKDIR /home/admin

# COPY 源路径 目标路径 从镜像中 COPY
COPY --from=builder /home/admin/bin/* /bin/
COPY --from=builder /home/admin/configs/ ./configs/

CMD ["gin-bench"]