# specimen-go

## 简介

帮助用户基于植物标本采集及鉴定数据，快速整理为标准化的植物标本数据，同时可自动从网络获取并补全植物相关的属性信息（体高、胸径、茎、叶、花、果实、寄主等等）、命名人、分类（门、纲、目、科、属）等各项重要信息。

## 使用方式

最新版本下载：[specimen-go 最新版本 v1.9.4](https://github.com/zxjsdp/specimen-go/releases)

- 对于 Windows 用户，建议下载使用已经编译好的用户界面（GUI）版本，点击运行即可；
- 对于 Linux & macOS 用户，仅支持命令行版本不支持 GUI。建议自行编译，也可使用已编译好的二进制可执行文件。


## 编译方式

### 方式一：命令行版本（支持 Linux, macOS & Windows 全平台）

获取 [specimen-go](https://github.com/zxjsdp/specimen-go):

    go get github.com/zxjsdp/specimen-go
    cd $GOPATH/src/github.com/zxjsdp/specimen-go

编译为可执行文件:

    go build -o specimen-go

试运行:

    # 示例文件说明：
    # → 流水号文件: snData.xlsx.sample
    # → 物种记录及鉴定文件: offlineData.xlsx.sample
    # 使用示例文件试运行：
    ./specimen-go \
        -s samples/snData.xlsx.sample \
        -d samples/offlineData.xlsx.sample \
        -o samples/output.xlsx

查看生成的结果文件：

    open samples/output.xlsx

运行测试用例:

    go test -v ./...


### 方式二：Windows 下 GUI 版本编译方式（MinGW）

安装 [walk](https://github.com/lxn/walk):

    go get github.com/lxn/walk
    
安装 [rsrs](https://github.com/akavel/rsrc):

    go get github.com/akavel/rsrc

生成需要包含进 golang 可执行文件中的二进制 shared library:

    cd "$GOPATH/src/github.com/zxjsdp/specimen-go/gui/windows"
    rsrc -manifest specimen-go-gui.exe.manifest -o rsrc.syso -ico "../resources/icon.ico"

生成 specimen-go.exe 可执行文件:

    go build -ldflags="-H windowsgui" -o specimen-go.exe
   

