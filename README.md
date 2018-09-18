# go-microservice
A code base from [blog](http://callistaenterprise.se/blogg/teknik/2017/02/17/go-blog-series-part1/)

# build
> cd accountservice
> export GOOS=linux
> go build -o accountservice

# test
> go test ./...

or run single package tests:
> go test github.com\navono\go-microservice\accountservice\service

## goConvey 前端
下载安装：
> go get -u github.com/smartystreets/goconvey

在项目根目录运行：
> goconvey.exe

在浏览器打开监听的`url`即可查看当前的所有测试的运行状况。

# docker
Note about proxy in docker build process.
