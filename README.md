# specimen-go

## 简介

用户根据植物标本采集信息快速整理为固定格式，并可以自动从网络或者植物标本相关的性状描述。运行程序后可以得到标准化的植物标本整理资料。

## 运行方式

下载已经编译好的可执行文件，点击运行即可。最新版本下载：[1.2.1](https://github.com/zxjsdp/specimen-go/releases)。

也可以在命令行中直接传入参数执行。

## 编译方式

获取方式

    go get github.com/zxjsdp/specimen-go

编译为可执行文件

    go build -o specimen-go
    
Windows 下 GUI 版本编译方式（MinGW）

    go get github.com/lxn/walk
    cd "$GOPATH/src/github.com/zxjsdp/specimen-go/gui/windows"
    go build -ldflags="-H windowsgui" -o specimen-go.exe
    .\specimen-go.exe
   
运行测试用例

    go test -v ./...

