[TOC]

# 教程文档地址
[https://doc.unitedrhino.com/pages/5235c3/](https://doc.unitedrhino.com/pages/5235c3/)

# 本地安装goctl【非官方，请使用i-Things/go-zero】

1. 本地将 `go-zero 项目克隆下来：  `git clone git@github.com:i-Things/go-zero.git`
2. 到目录 `go-zero\tools\goctl 下 执行命令： `go install`
3. 后续执行下面的各种goctl命令即可

# 环境初始化

`protoc/protoc-gen-go/protoc-gen-grpc-go` 依赖可以通过下列命令 一键安装

```shell
$ goctl env check --install --verbose --force
```

# 服务新增方案

## rpc服务
```
goctl rpc new lightsvr  --style=goZero -m
```
## api服务
```
goctl api new lightsvr  --style=goZero 
```

# 库表新增方案

在每个服务的 `internal/repo/relationDB` 目录下有example.go 
1. 借助 `https://sql2gorm.mccode.info/` 生成对应的模型 放到 `internal/repo/relationDB/modle.go` 中
2. 复制 `internal/repo/relationDB/example.go` 到对应目录下,并修改表名
3. 将example.go中的Example替换为表名
4. 定制修改对应函数即可

# 代码生成

```shell
#cd apisvr && goctl api go -api http/api.api  -dir ./  --style=goZero && cd ..
goctl api go -api http/lightsvr.api  -dir ./  --style=goZero -ws  && goctl api swagger -filename swagger.json -api http/lightsvr.api -dir ./http&&  goctl api access  -api http/lightsvr.api -dir ./http 
goctl api access  -api http/lightsvr.api -dir ./http 
goctl api swagger -filename swagger.json -api http/lightsvr.api -dir ./http
```
