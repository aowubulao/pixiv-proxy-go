# Pixiv-Proxy

简单的pixiv图片反代

e.g.

url: https://i.pximg.net/img-original/img/2021/03/25/19/45/12/88697292_p0.png
-> https://your.domain/img-original/img/2021/03/25/19/45/12/88697292_p0.png

## 启动

自己编译或下载编译好的对应的二进制文件

e.g.

```shell
#  Listen to 0.0.0.0:10901
./pixiv-proxy-go-linux-amd64
```

## 配置

支持：

- 端口(默认10901)
- 代理服务器(默认不使用)

```shell
# 使用配置启动
./pixiv-proxy-go-linux-amd64 port=2333 proxy=http://127.0.0.1:7890
```

## nginx配置

如果要实现开头的反代效果，需要配置服务器，这里提供nginx的配置项：

```
location / {
			proxy_set_header x-forwarded-for $remote_addr;
			proxy_pass	http://127.0.0.1:10901/proxy/;
}
```

