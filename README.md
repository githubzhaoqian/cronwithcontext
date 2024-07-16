# CronWithContext

CronWithContext By github.com/robfig/cron/v3@v3.0.1 extension of job added custom context, better implementation of the middleware.!

CronWithContext 由github.com/robfig/cron/v3@v3.0.1 扩展而来，对job增加了自定义context,能更好的实现中间件。

To download the specific tagged release, run:

	go get github.com/githubzhaoqian/cronwithcontex/v3@v3.0.2

Import it in your program as:

	import "github.com/githubzhaoqian/cronwithcontex/v3"

It requires Go 1.11 or later due to usage of Go Modules.

基本使用和cron/v3库一样，只是增加了一个context参数，可以传递一些自定义参数。
