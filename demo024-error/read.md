这里有个好消息是，这样的代码不必再写了，有一个第三方的错误库（github.com/pkg/errors），对于这个库，我无论到哪都能看到他的存在，所以，这个基本上来说就是事实上的标准了。代码示例如下：

`
    import "github.com/pkg/errors"
    
    
    //错误包装
    if err != nil {
    return errors.Wrap(err, "read failed")
    }
    
    // Cause接口
    switch err := errors.Cause(err).(type) {
    case *MyError:
    // handle specifically
    default:
    // unknown error
    }
`