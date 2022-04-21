#!/usr/bin/env bash
rm *.bin
#静态编译
go build --ldflags "-extldflags -static" -o mygoproject.bin main.go

#docker build
sudo docker build -t mygoproject .

#查看本地docker 镜像
#sudo docker images

#打标签
sudo docker tag mygoproject  registry.cn-beijing.aliyuncs.com/huochengyan/hcy:1.0

## 登录镜像
sudo docker login --username=929914120@qq.com registry.cn-beijing.aliyuncs.com

#上传到docker服务期
sudo docker push registry.cn-beijing.aliyuncs.com/huochengyan/hcy:1.0
#服务上获取最新的镜像（服务器上执行）
# sudo docker stop  mygoproject
# sudo docker rm mygoproject
# sudo docker pull registry.cn-beijing.aliyuncs.com/huochengyan/hcy:1.0

#服务期上启动（服务器上执行,指定配置文件执行）
# sudo docker run -i -t -v ~/applogs/mygoproject/logs:/usr/mygoproject/logs -v ~/applogs/mygoproject/conf:/usr/mygoproject/conf  -p 8600:8600 --restart=always -d --name=mygoproject registry.cn-beijing.aliyuncs.com/huochengyan/hcy:1.0