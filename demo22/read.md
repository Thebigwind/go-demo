在 Go 语言中，你一定要使用 time.Time 和 time.Duration 两个类型：

在命令行上，flag 通过 time.ParseDuration 支持了 time.Duration
JSon 中的 encoding/json 中也可以把time.Time 编码成 RFC 3339 的格式
数据库使用的 database/sql 也支持把 DATATIME 或 TIMESTAMP 类型转成 time.Time
YAML你可以使用 gopkg.in/yaml.v2 也支持 time.Time 、time.Duration 和 RFC 3339 格式
如果你要和第三方交互，实在没有办法，也请使用 RFC 3339 的格式。




性能提示

1.如果需要把数字转字符串，使用 strconv.Itoa() 会比 fmt.Sprintf() 要快一倍左右
2.尽可能地避免把String转成[]Byte 。这个转换会导致性能下降。
3.如果在for-loop里对某个slice 使用 append()请先把 slice的容量很扩充到位，这样可以避免内存重新分享以及系统自动按2的N次方幂进行扩展但又用不到，从而浪费内存。
4.使用StringBuffer 或是StringBuild 来拼接字符串，会比使用 + 或 += 性能高三到四个数量级。
5.尽可能的使用并发的 go routine，然后使用 sync.WaitGroup 来同步分片操作
6.避免在热代码中进行内存分配，这样会导致gc很忙。尽可能的使用 sync.Pool 来重用对象
7.使用 lock-free的操作，避免使用 mutex，尽可能使用 sync/Atomic包。 （关于无锁编程的相关话题，可参看《无锁队列实现》或《无锁Hashmap实现》）
8.使用 I/O缓冲，I/O是个非常非常慢的操作，使用 bufio.NewWrite() 和 bufio.NewReader() 可以带来更高的性能。
9.对于在for-loop里的固定的正则表达式，一定要使用 regexp.Compile() 编译正则表达式。性能会得升两个数量级。
10.如果你需要更高性能的协议，你要考虑使用 protobuf 或 msgp 而不是JSON，因为JSON的序列化和反序列化里使用了反射
11.你在使用map的时候，使用整型的key会比字符串的要快，因为整型比较比字符串比较要快。

