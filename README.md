# dnspod
DNSPod官方提供API的Go封装
使用时请查证官方版本，据说2019年官方推出了新版本的API并要求迁移。

## 构建

```
git clone https://github.com/op-y/dnspod.git

cd dnspod

mkdir -p bin

go build -o bin/dnspod

```

## 操作

```
启动   ：sh control start

查看PID：sh control pid

停止   ：sh control stop

重启   ：sh control restart
```
