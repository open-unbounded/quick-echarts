# Quick-ECharts

一个简单高效的ECharts图表快速生成工具，基于Go语言和Gin框架开发。

## 功能特点

- 快速生成ECharts图表，无需编写HTML和JavaScript代码
- 通过简单的HTTP请求即可渲染复杂的图表
- 支持所有ECharts图表类型（折线图、柱状图、饼图等）
- 轻量级服务，易于部署和使用
- 支持自定义图表配置选项

## 安装

### 从源码构建

```bash
# 克隆仓库
git clone https://github.com/open-unbounded/quick-echarts.git
cd quick-echarts

# 构建项目
go build -o quick-echarts
```

### 使用预编译二进制文件

可以使用build.sh脚本为不同平台构建二进制文件：

```bash
./build.sh
```

## 使用方法

### 启动服务

```bash
# 默认在0.0.0.0:8080端口启动
./quick-echarts

# 指定端口启动
./quick-echarts --addr=127.0.0.1:3000
```

### 生成图表

向`/chart`接口发送GET请求，通过`c`参数传递ECharts配置选项（JSON格式）：

```
http://localhost:8080/chart?c={"xAxis":{"type":"category","data":["Mon","Tue","Wed","Thu","Fri","Sat","Sun"]},"yAxis":{"type":"value"},"series":[{"data":[150,230,224,218,135,147,260],"type":"line"}]}
```

### 示例

1. 简单折线图

```
http://localhost:8080/chart?c={"xAxis":{"type":"category","data":["Mon","Tue","Wed","Thu","Fri","Sat","Sun"]},"yAxis":{"type":"value"},"series":[{"data":[150,230,224,218,135,147,260],"type":"line"}]}
```
![example1.png](img/example1.png)

2. 柱状图

```
http://localhost:8080/chart?c={"xAxis":{"type":"category","data":["Mon","Tue","Wed","Thu","Fri","Sat","Sun"]},"yAxis":{"type":"value"},"series":[{"data":[120,200,150,80,70,110,130],"type":"bar"}]}
```
![example2.png](img/example2.png)
## 依赖

- [Gin Web Framework](https://github.com/gin-gonic/gin) - HTTP Web框架
- [ECharts](https://echarts.apache.org/) - 图表渲染库

## 许可证

本项目采用 [Apache License 2.0](LICENSE) 许可证。