> # go-zero api文件编写
>
> 2022-04-25 17:47 更新
>
> ## api文件编写
>
> ### 编写user.api文件
>
> ```go
> $ vim service/user/api/user.api
> type (
>     LoginReq {
>         Username string `json:"username"`
>         Password string `json:"password"`
>     }
> 
>     LoginReply {
>         Id           int64 `json:"id"`
>         Name         string `json:"name"`
>         Gender       string `json:"gender"`
>         AccessToken  string `json:"accessToken"`
>         AccessExpire int64 `json:"accessExpire"`
>         RefreshAfter int64 `json:"refreshAfter"`
>     }
> )
> 
> service user-api {
>     @handler login
>     post /user/login (LoginReq) returns (LoginReply)
> }
> ```
>
> ### 生成api服务
>
> - 方式一
>
> ```bash
> $ cd book/service/user/api
> $ goctl api go -api user.api -dir .
> Done.
> ```