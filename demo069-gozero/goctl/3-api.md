# goctl api参数说明
    goctl api [go/java/ts] [-api user/user.api] [-dir ./src]

api 后面接生成的语言，现支持go/java/typescript -api 自定义api所在路径 -dir 自定义生成目录

如需自定义模板，运行如下命令生成api gateway模板：

    goctl api go template

生成的模板放在$HOME/.goctl目录下，根据需要自行修改模板，下次运行goctl生成代码时会优先采用模板文件的内容

## API 语法说明
    info(
    title: "doc title"
    desc: "doc description first part，doc description second part"
    version: "1.0"
    )
    
    type User {
    name string `json:"user"` // 用户姓名
    }
    
    type Student {
    name string `json:"name"` // 学生姓名
    }
    
    type Teacher {
    }
    
    type (
    Address {
    city string `json:"city"`
    }
    
        InnerType {
            image string `json:"image"`
        }
    
        CreateRequest {
            InnerType
            name string `form:"name"`
            age int `form:"age,optional"`
            address []Address `json:"address,optional"`
        }
    
        GetRequest {
            name string `path:"name"`
            age int `form:"age,optional"`
        }
    
        GetResponse {
            code int `json:"code"`
            desc string `json:"desc,omitempty"`
            address Address `json:"address"`
            Service int `json:"service"`
        }
    )
    
    @server(
    group: user
    )
    service user-api {
    @doc(
    summary: "user title"
    )
    @handler GetUserHandler
    get /api/user/:name (GetRequest) returns (GetResponse)
    
        @handler CreateUserHandler
        post /api/users/create (CreateRequest)
    }
    
    @server(
    jwt: Auth
    group: profile
    )
    service user-api {
    @doc(
    summary: "user title"
    )
    @handler GetProfileHandler
    get /api/profile/:name (GetRequest) returns (GetResponse)
    
        @handler CreateProfileHandler
        post /api/profile/create (CreateRequest)
    }
    
    service user-api {
    @doc(
    summary: "desc in one line"
    )
    @handler PingHandler
    head /api/ping ()
    }

### 总结如下：

#### 1.info部分：描述了api基本信息，比如Auth，api是哪个用途。
#### 2.type部分：type类型声明和golang语法兼容（不支持type alias）。
#### 3.service部分：
    service代表一组服务，一个服务可以由多组名称相同的service组成，可以针对每一组service配置jwt和auth认证。
    通过group属性可以指定service生成所在子目录，支持多级子目录a/b/c。
    service里面包含api路由，比如上面第一组service的第一个路由，doc用来描述此路由的用途，GetProfileHandler表示处理这个路由的handler， get /api/profile/:name(getRequest) returns(getResponse) 中get代表api的请求方式（get/post/put/delete）, /api/profile/:name 描述了路由path，:name通过 请求getRequest里面的属性赋值，getResponse为返回的结构体，这两个类型都定义在2描述的类型中。
    server 标签支持配置middleware，示例如下：
        @server(
            middleware: AuthUser
        )

    添加完middleware后需要设置ServiceContext 中middleware变量的值，middleware实现可以参考测试用例 TestWithMiddleware 或者 TestMultiMiddlewares。

    handler 支持缩写，实例如下：
    @handler CreateProfileHandler
    post /api/profile/create(CreateRequest)

#### 4.支持在info下面和type顶部import外部api文件，import语法：import "xxxx.api"



根据定义好的api文件生成golang代码
命令如下： goctl api go -api user/user.api -dir user

    .
    ├── internal
    │   ├── config
    │   │   └── config.go
    │   ├── handler
    │   │   ├── pinghandler.go
    │   │   ├── profile
    │   │   │   ├── createprofilehandler.go
    │   │   │   └── getprofilehandler.go
    │   │   ├── routes.go
    │   │   └── user
    │   │       ├── createuserhandler.go
    │   │       └── getuserhandler.go
    │   ├── logic
    │   │   ├── pinglogic.go
    │   │   ├── profile
    │   │   │   ├── createprofilelogic.go
    │   │   │   └── getprofilelogic.go
    │   │   └── user
    │   │       ├── createuserlogic.go
    │   │       └── getuserlogic.go
    │   ├── svc
    │   │   └── servicecontext.go
    │   └── types
    │       └── types.go
    └── user.go

生成的代码可以直接跑，有几个地方需要改：

    在servicecontext.go里面增加需要传递给logic的一些资源，比如mysql, redis，rpc等
    在定义的get/post/put/delete等请求的handler和logic里增加处理业务逻辑的代码


