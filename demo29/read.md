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

