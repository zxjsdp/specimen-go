package config

const (
	DemoHTMLFileName = "标本录入软件数据格式.html"

	DemoSNFileName      = "流水号文件（示例）.xlsx"
	DemoOfflineFileName = "鉴定录入文件（示例）.xlsx"
)

const USAGE = `

使用方法：

1. 直接双击 specimen-go 文件以运行用户界面程序（仅在 Windows 平台下支持 GUI）；
2. 在命令行中通过参数运行（支持所有平台）：

    ./specimen-go -s 流水号数据.xlsx -d 鉴定录入文件.xlsx -o 输出文件.xlsx

`

const HelpMessage = `
使用方式：

1. 选择或输入 “录入鉴定文件名”；
2. 选择或输入 “流水号文件名”；
3. 选择或输入 “输出文件名”；
4. 点击 “开始处理” 以执行标本数据整理、网络查询以及信息聚合任务。

如果需要，可以点击 “数据格式” 按钮以查看程序的数据格式。
如果需要，可以点击 “示例文件” 按钮以生成示例文件，作为模板。
`

var SNFileDemoData = [][]string{
	{"ZY20170001", "123167", "107930", "1"},
	{"ZY20170001", "123168", "107931", "2"},
	{"ZY20170001", "123168", "107932", "3"},
	{"ZY20170002", "123170", "107933", "1"},
	{"ZY20170002", "123171", "107934", "2"},
	{"ZY20170002", "123172", "107935", "3"},
}

var OfflineDemoData = [][]string{
	{"ZY20170001", "蔓长春花", "Vinca major", "夹竹桃科", "Apocynaceae", "上海", "上海市", "杨浦区淞沪路嘉誉湾", "N31°20′30.75″", "E121°30′10.99″", "10", "20170423", "3", "半灌木", "张三", "李四", "20170426", "王五", "20171018"},
	{"ZY20170002", "红花檵木", "Loropetalum chinense", "金缕梅科", "Hamamelidaceae", "上海", "上海市", "杨浦区淞沪路嘉誉湾", "N31°20′30.76″", "E121°30′10.100″", "10", "20170423", "3", "灌木", "张三", "李四", "20170426", "王五", "20171018"},
}

const DemoHTMLContent = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Demo</title>
    <style type="text/css">
        .tg  {border-collapse:collapse;border-spacing:0;border-color:#ccc;}
        .tg td{font-family:Arial, sans-serif;font-size:14px;padding:10px 5px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:#ccc;color:#333;background-color:#fff;}
        .tg th{font-family:Arial, sans-serif;font-size:14px;font-weight:normal;padding:10px 5px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:#ccc;color:#333;background-color:#f0f0f0;}
        .tg .tg-h31u{font-family:Arial, Helvetica, sans-serif !important;;vertical-align:top}
        .tg .tg-yw4l{vertical-align:top}
    </style>
</head>
<body>
<h2>鉴定录入文件示例</h2>
<table class="tg">
    <tr>
        <th class="tg-yw4l">物种编号</th>
        <th class="tg-yw4l">中文名</th>
        <th class="tg-yw4l">种名（拉丁）</th>
        <th class="tg-yw4l">科名</th>
        <th class="tg-yw4l">科名（拉丁）</th>
        <th class="tg-yw4l">省</th>
        <th class="tg-yw4l">市</th>
        <th class="tg-yw4l">具体小地名</th>
        <th class="tg-yw4l">纬</th>
        <th class="tg-yw4l">东经</th>
        <th class="tg-yw4l">海拔</th>
        <th class="tg-yw4l">日期</th>
        <th class="tg-yw4l">份数</th>
        <th class="tg-yw4l">草灌</th>
        <th class="tg-yw4l">采集人</th>
        <th class="tg-yw4l">鉴定人</th>
        <th class="tg-yw4l">鉴定日期</th>
        <th class="tg-yw4l">录入人</th>
        <th class="tg-yw4l">录入日期</th>
    </tr>
    <tr>
        <td class="tg-yw4l">ZY20170001</td>
        <td class="tg-yw4l">蔓长春花</td>
        <td class="tg-yw4l">Vinca major</td>
        <td class="tg-yw4l">夹竹桃科</td>
        <td class="tg-h31u">Apocynaceae</td>
        <td class="tg-yw4l">上海</td>
        <td class="tg-yw4l">上海市</td>
        <td class="tg-yw4l">杨浦区淞沪路嘉誉湾</td>
        <td class="tg-yw4l">N31°20′30.75″</td>
        <td class="tg-yw4l">E121°30′10.99″</td>
        <td class="tg-yw4l">10</td>
        <td class="tg-yw4l">20170423</td>
        <td class="tg-yw4l">3</td>
        <td class="tg-yw4l">半灌木</td>
        <td class="tg-yw4l">张三</td>
        <td class="tg-yw4l">李四</td>
        <td class="tg-yw4l">20170426</td>
        <td class="tg-yw4l">王五</td>
        <td class="tg-yw4l">20171018</td>
    </tr>
    <tr>
        <td class="tg-yw4l">ZY20170002</td>
        <td class="tg-yw4l">红花檵木</td>
        <td class="tg-yw4l">Loropetalum chinense</td>
        <td class="tg-yw4l">金缕梅科</td>
        <td class="tg-yw4l">Hamamelidaceae</td>
        <td class="tg-yw4l">上海</td>
        <td class="tg-yw4l">上海市</td>
        <td class="tg-yw4l">杨浦区淞沪路嘉誉湾</td>
        <td class="tg-yw4l">N31°20′30.76″</td>
        <td class="tg-yw4l">E121°30′10.100″</td>
        <td class="tg-yw4l">10</td>
        <td class="tg-yw4l">20170423</td>
        <td class="tg-yw4l">3</td>
        <td class="tg-yw4l">灌木</td>
        <td class="tg-yw4l">张三</td>
        <td class="tg-yw4l">李四</td>
        <td class="tg-yw4l">20170426</td>
        <td class="tg-yw4l">王五</td>
        <td class="tg-yw4l">20171018</td>
    </tr>
</table>


<h2>流水号文件示例</h2>

<table class="tg">
    <tr>
        <th class="tg-yw4l">物种编号</th>
        <th class="tg-yw4l">流水号</th>
        <th class="tg-yw4l">条形码</th>
        <th class="tg-yw4l">同一物种的个体编号（1、2、3、...）</th>
    </tr>
    <tr>
        <td class="tg-yw4l">ZY20170001</td>
        <td class="tg-yw4l">123167</td>
        <td class="tg-yw4l">107930</td>
        <td class="tg-yw4l">1</td>
    </tr>
    <tr>
        <td class="tg-yw4l">ZY20170001</td>
        <td class="tg-yw4l">123168</td>
        <td class="tg-yw4l">107931</td>
        <td class="tg-yw4l">2</td>
    </tr>
    <tr>
        <td class="tg-yw4l">ZY20170001</td>
        <td class="tg-yw4l">123168</td>
        <td class="tg-yw4l">107932</td>
        <td class="tg-yw4l">3</td>
    </tr>
    <tr>
        <td class="tg-yw4l">ZY20170002</td>
        <td class="tg-yw4l">123170</td>
        <td class="tg-yw4l">107933</td>
        <td class="tg-yw4l">1</td>
    </tr>
    <tr>
        <td class="tg-yw4l">ZY20170002</td>
        <td class="tg-yw4l">123171</td>
        <td class="tg-yw4l">107934</td>
        <td class="tg-yw4l">2</td>
    </tr>
    <tr>
        <td class="tg-yw4l">ZY20170002</td>
        <td class="tg-yw4l">123172</td>
        <td class="tg-yw4l">107935</td>
        <td class="tg-yw4l">3</td>
    </tr>
</table>

</body>
</html>
`
