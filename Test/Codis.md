#   codis redis集群

github开源地址

* [Codis官方开源地址](https://github.com/CodisLabs/codis)

* [Codis官方中文文档](https://github.com/CodisLabs/codis/blob/release3.2/doc/tutorial_zh.md)


### 准备工作
下载Codis 最新安装包  codis3.2.2-go1.9.2-linux.tar.gz

下载组件Zookeeper 最新安装包  zookeeper-3.4.9.tar.gz

下载 Java JDK  jdk-8u111-linux-x64.tar.gz


### 1.安装Java JDK

这里放在了Home 目录 ,路径在环境变量里别配错了
解压 jdk-8u111-linux-x64.tar.gz 到 /home/thk/  


安装 jdk

配置 环境变量

    vi ~/.bashrc    

    export JAVA_HOME=/home/thk/jdk1.8.0_111/
    export JRE_HOME=${JAVA_HOME}/jre
    export CLASSPATH=.:${JAVA_HOME}/lib:${JRE_HOME}/lib
    export PATH=${JAVA_HOME}/bin:$PATH

最后生效
    source /etc/profile

测试是否成功
    
    
    thk@hcy-System-Product-Name:~$ java -version
    java version "1.8.0_111"
    Java(TM) SE Runtime Environment (build 1.8.0_111-b14)
    Java HotSpot(TM) 64-Bit Server VM (build 25.111-b14, mixed mode)
    thk@hcy-System-Product-Name:~$ 



### 2.安装 zookeeper

解压 zookeeper-3.4.9.tar.gz 到 /usr/local/zookeeper-3.4.9
    
    mv zookeeper-3.4.9.tar.gz /usr/local
    
    cd /usr/local
    
    tar -zxvf zookeeper-3.4.9.tar.gz
    
    cd zookeeper-3.4.9/

进入 zookeeper-3.4.9/

     cd conf/ 
     
     cp   zoo_sample.cfg zoo.cfg
     sudo vim zoo.cfg

设置这两行：(分别创建 data 目录、logs目录)

    dataDir=/usr/local/zookeeper-3.4.9/data
    dataLogDir=/usr/local/zookeeper-3.4.9/logs

进入 /bin 目录启动zookeeper

    ./zkServer.sh start

使用命令jps 能够看到进程 ，即可

    jps
    29618 QuorumPeerMain

使用 ./zkServer.sh status 查看zookeeper状态



### 3.生成Codis配置文件

解压 codis3.2.2-go1.9.2-linux.tar.gz 安装包
    
进入目录
创建存放配置文件和logs的输出目录

    mkdir logs 
    mkdir conf
    mkdir redis-conf

##### 3.1 生成 dashboard配置文件

   生成dashboard.toml 配置文件
   
    ./codis-dashboard --default-config | tee conf/dashboard.toml
   
   进入dashboard.toml 设置访问密码这里和Redis的密码一样
   
   product_auth = "123456"
   
##### 3.2 生成 proxy  配置文件
   生成proxy.tom 配置文件
    
    ./codis-proxy --default-config | tee conf/proxy.toml
    
   设置密码
    
    product_auth = "123456"
    session_auth = "123456"
    
##### 3.3 生成codis.json配置
    
    ./codis-admin --dashboard-list --zookeeper=127.0.0.1:2181 | tee conf/codis.json
    
### 4.启动集群

>启动1 codis-dashboard

nohup ./codis-dashboard --ncpu=4 --config=conf/dashboard.toml     --log=logs/dashboard.log --log-level=WARN &

查看http://127.0.0.1:18080/topom 看到json已启动

>启动2：

nohup ./codis-proxy --ncpu=4 --config=conf/proxy.toml --log=logs/proxy.log --log-level=WARN &

>启动3  (第一次启动需要绑定，不用再次绑定了。)

./codis-admin --dashboard=127.0.0.1:18080 --create-proxy -x 127.0.0.1:11080  

   
启动默认的Redis 库（6379端口的那个）
./codis-server
这里使用自定义配置文件启动

在redis-conf文件中分别编辑 

启动redis master 库 slave 配置文件6380.conf 6381.conf  20186.conf 20187.conf 20188.conf 20189.conf


注意修改配置文件里的几个地方

分别为绑定的端口、pid文件、日志输出、数据库名字

* port 6380
* pidfile "/tmp/redis_6380.pid"
* logfile "/tmp/redis_6380.log"
* dbfilename "dump6380.rdb"

最后这里在配置文件设置密码
在 # requirepass foobared 添加设置密码

requirepass "123456"

master 的配置文件

    nohup ./codis-server ./redis-conf/6380.conf &> ./logs/redis_6380.log &
    nohup ./codis-server ./redis-conf/6381.conf &> ./logs/redis_6381.log &

slave 配置文件

6380的副本

    nohup ./codis-server ./redis-conf/20186.conf &> ./logs/redis_20186.log &
    nohup ./codis-server ./redis-conf/20187.conf &> ./logs/redis_20187.log &

6381的副本

    nohup ./codis-server ./redis-conf/20188.conf &> ./logs/redis_20188.log &
    nohup ./codis-server ./redis-conf/20189.conf &> ./logs/redis_20189.log &


这时候访问
http://127.0.0.1:18090/
进入管理端

左侧能看到默认的 codis -demo
右侧管理整个集群设置


至此完毕

### 5. FAQ 请查看官方文档
    退出重启
    
    ./codis-admin --dashboard=127.0.0.1:18080 --remove-proxy --addr=127.0.0.1:11080 --force
    ./codis-admin --dashboard="127.0.0.1:18080" --shutdown
    ./codis-admin --remove-lock --product=codis-demo --zookeeper=127.0.0.1:2181