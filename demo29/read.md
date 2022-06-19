gin框架中集成pprof

这里说的是以http的方式集成，如果是非http的方式，那你在任何地方添加代码都可以。

在gin框架中集成pprof首先要引入包_ "net/http/pprof"。

之后有两种方式，第一种是新监听另一个端口作为pprof http

    go func() {
        log.Println(http.ListenAndServe(":6060", nil))
    }()

第二种是共用服务的端口，这里需要修改路由注册方式

    package main
    
    import (
    "net/http"
    
        "github.com/gin-contrib/pprof"
        "github.com/gin-gonic/gin"
    )
    
    func main() {
    app := gin.Default()
    
        pprof.Register(app) // 性能
    
        app.GET("/test", func(c *gin.Context) {
            c.String(http.StatusOK, "test")
        })
        app.Run(":3000")
    }


分析性能
https://blog.csdn.net/u012189747/article/details/122704004

获取.pb.gz文件
go tool pprof http://localhost:6060/debug/pprof/profile

浏览器上查看图片化信息
注意：.pb.gz文件路径可根据上个命令得到。

输入命令：
go tool pprof -http=:8080 /Users/zhuyun/pprof/pprof.samples.cpu.001.pb.gz

自动打开浏览器，可看到图片化信息，




https://www.jianshu.com/p/4e4ff6be6af9

https://studygolang.com/articles/20529

https://www.bbsmax.com/A/6pdDqWByzw/

https://eddycjy.gitbook.io/golang/di-9-ke-gong-ju/go-tool-pprof



