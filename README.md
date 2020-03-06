# 备忘：

zookeeper
安装：brew install zookeeper
安装位置:  /usr/local/Cellar/zookeeper
启动：zkServer start
关闭：zkServer stop
 
kafka
安装:  brew install kafka,
安装位置：/usr/local/etc/kafka/
1. 启动：
kafka-server-start /usr/local/etc/kafka/server.properties &​
2. 关闭
kafka-server-stop

java环境变量
JAVA_HOME=/Library/Java/JavaVirtualMachines/jdk-13.0.2.jdk/Contents/Home
PATH=$JAVA_HOME/bin:$PATH:.

elasticsearch
安装路径
/Users/xushengnan02/elasticsearch-7.6.1
启动：bin/elasticsearch
访问：curl http://localhost:9200/