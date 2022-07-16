https://www.liaoxuefeng.com/wiki/1252599548343744/1305366354722849
https://www.cnblogs.com/foxclever/p/8593072.html



存储用户的哈希口令时，要加盐存储，目的就在于抵御彩虹表攻击

Hmac算法就是一种基于密钥的消息认证码算法，它的全称是Hash-based Message Authentication Code，是一种更安全的消息摘要算法。

Hmac算法总是和某种哈希算法配合起来用的。例如，我们使用MD5算法，对应的就是HmacMD5算法，它相当于“加盐”的MD5：

HmacMD5 ≈ md5(secure_random_key, input)

HmacMD5可以看作带有一个安全的key的MD5。使用HmacMD5而不是用MD5加salt，有如下好处：

HmacMD5使用的key长度是64字节，更安全；
Hmac是标准算法，同样适用于SHA-1等其他哈希算法；
Hmac输出和原有的哈希算法长度一致。
可见，Hmac本质上就是把key混入摘要的算法。验证此哈希时，除了原始的输入数据，还要提供key。
为了保证安全，我们不会自己指定key，而是通过标准库的生成一个安全的随机的key。


