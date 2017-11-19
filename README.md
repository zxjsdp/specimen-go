# specimen-go

## 简介

用户根据植物标本采集信息快速整理为固定格式，并可以自动从网络或者植物标本相关的性状描述。运行程序后可以得到标准化的植物标本整理资料。

## 运行方式

下载已经编译好的用户界面版本，点击运行即可。最新版本下载：[specimen-go 最新 V1.8.0 版本](https://github.com/zxjsdp/specimen-go/releases)。

也可以使用命令行版本，通过传入参数执行。

## 编译方式

### 命令行版本

获取 [specimen-go](https://github.com/zxjsdp/specimen-go)

    go get github.com/zxjsdp/specimen-go

编译为可执行文件

    go build -o specimen-go
    
运行测试用例

    go test -v ./...

### Windows 下 GUI 版本编译方式（MinGW）

安装 [walk](https://github.com/lxn/walk)

    go get github.com/lxn/walk
    
安装 [rsrs](https://github.com/akavel/rsrc)

    go get github.com/akavel/rsrc

生成需要包含进 golang 可执行文件中的二进制 shared library

    cd "$GOPATH/src/github.com/zxjsdp/specimen-go/gui/windows"
    rsrc -manifest specimen-go-gui.exe.manifest -o rsrc.syso -ico "../resources/icon.ico"

生成 specimen-go.exe 可执行文件

    go build -ldflags="-H windowsgui" -o specimen-go.exe
   

