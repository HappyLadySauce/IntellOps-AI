# GoFLly企业版V2版本-全新升级 
# 一、安装

## 1 fresh 热更新-热加载-边开发边编译
```
go install github.com/pilu/fresh@latest
```
根目录下的 runner.conf 是fresh热编译配置文件
目前默认忽略 utils  

## 2 初始化mod
```
go mod tidy
```
## 3 热编译运行
```
 fresh 
```
热加载使用fresh或者air都可以两者区别不大，air配置比较直观，根据自己喜好选择。
注意：如果热加载太慢，请把views内的前端代码移到go项目目录外，这样就可快速热加载了，加快开发速度，移除后记得把resource\config.yaml下的vueobjroot修改成新目录路径。
# 二、打包
## 开发电脑系统为：windows
### 1 windows系统下打包在windows系统程序
```
go build main.go
```
### 2 windows系统下打包在Linux上运行（推荐）
（此时会打包成Linux上可运行的二进制文件，不带后缀名的文件）
```
SET GOOS=linux
SET GOARCH=amd64
go build
```
### 3 切换回widows系统程序
```
// 配置环境变量
SET CGO_ENABLED=1
SET GOOS=windows
SET GOARCH=amd64

// 编译命令
go build main.go

```
### 7 编译成Linux环境可执行文件

```
// 配置参数
SET CGO_ENABLED=0 
SET GOOS=linux 
SET GOARCH=amd64 
// 编译命令
go build main.go

```
## 开发电脑系统为：mac
```
// 配置参数
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
// 或者
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

```