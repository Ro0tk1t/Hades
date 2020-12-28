# Hades
(a multi service & application password crack tool)

## 开发环境
> Debian系  
> vim  
> golang1.15  

## 使用：

### 编译使用：
```bash
$ go build hades
$ ./hades --help
$ ./hades -I '127.0.0.1:23' -t 10 -T 100 --user_file /tmp/users --pass_file /tmp/pass scan
```

### 或者直接使用：
```bash
$ go run hades.go --help
$ go run hades.go -I '127.0.0.1:23' -t 10 -T 100 --user_file /tmp/users --pass_file /tmp/pass scan
```


## 扫描流程：

* 获取命令参数
* 检测输入的目标是否存活
* 根据协程数参数控制爆破协程并发
* 根据超时参数控制爆破协程自动超时
* 根据输出参数输出到json文件或者屏幕


> 可以用--ip选项检测单个目标 或者 用--ipfile检测多个目标，格式支持`ip:port|protocol` 和完整url格式

> --none参数检测空密码，--quit对同一个目标检测到一对口令后就退出检测该目标

> --threads控制对同一目标检测的并发数
