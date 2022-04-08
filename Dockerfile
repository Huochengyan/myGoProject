FROM alpine
RUN apk add --no-cache tzdata \
&& ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
&& echo "Asia/Shanghai" > /etc/timezone
RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true

WORKDIR /usr/mygoproject
COPY conf ./conf/
#COPY resources ./resources/
COPY mygoproject.bin .

EXPOSE 8600
CMD ["./mygoproject.bin"]

