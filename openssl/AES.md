https://its201.com/article/u011832525/112278849

密码学中的高级加密标准（Advanced Encryption Standard，AES），又称Rijndael加密法，是美国联邦政府采用的一种区块加密标准。
这个标准用来替代原先的DES，已经被多方分析且广为全世界所使用

AES属于对称加密算法，加解密使用同一个秘钥。
对称加密算法，一般有至少4种模式，即ECB、CBC、CFB、OFB等。

ECB模式
电子密码本模式 Electronic Code Book(ECB)。它将加密的数据分成若干组，每组的大小跟加密密钥长度相同，然后每组都用相同的密钥进行加密。

其缺点是：电子编码薄模式用一个密钥加密消息的所有块，如果原消息中重复明文块，则加密消息中的相应密文块也会重复，因此，电子编码薄模式适于加密小消息。

命令行操作
使用aes-128-ecb对hello.txt加密，128位密钥为8cc72b05705d5c46f412af8cbed55aad，密文为hello.en。

    openssl enc -e -aes-128-ecb -in hello.txt -out hello.en -K 8cc72b05705d5c46f412af8cbed55aad

使用aes-128-ecb对hello.en解密，128位密钥为8cc72b05705d5c46f412af8cbed55aad，解密后的文件为hello.de。

    openssl enc -d -aes-128-ecb -in hello.en -out hello.de -K 8cc72b05705d5c46f412af8cbed55aad


2、函数说明

生成加密/解密的Key：
int AES_set_encrypt_key(const unsigned char *userKey, const int bits, AES_KEY *key);
int AES_set_decrypt_key(const unsigned char *userKey, const int bits, AES_KEY *key);

userKey：用户指定的密码。注意：只能是16/24/32字节。如果密码字符串长度不够，可以在字符串末尾追加一些特定的字符，或者重复密码字符串，直到满足最少的长度。
bits:    密码位数。即userKey的长度 * 8，只能是128/192/256位。
key:     向外输出参数
返回值:   0 - 成功； 1 - userkey，key为空；-2 - 密钥长度不是128/192/256

AES ECB加密/解密：
void AES_ecb_encrypt(const unsigned char *in, unsigned char *out, const AES_KEY *key, const int enc);
in:	 输入数据，长度固定为16字节
out: 输出数据，长度与输入数据一致，固定为16字节
key: 使用AES_set_encrypt_key/AES_set_decrypt_key生成的Key
enc: AES_ENCRYPT 代表加密， AES_DECRYPT代表解密


3、编程实现

由于ECB模式，每次只能处理一个块的数据，即16字节，所以如果需要处理任意长度的数据，那么需要在原始数据末尾，先进行填充，使得数据长度为16的整数倍，随后再分块进行加密。
解密时，也需要分块解密，最后将解密后的数据，进行取消填充。

1）PKCS7填充方式
AES支持多种填充方式：如NoPadding、PKCS5Padding、ISO10126Padding、PKCS7Padding、ZeroPadding。密文长度与填充方式关系，可参考《AES加密模式和填充模式》。

PKCS7填充方式：
假设数据长度需要填充n(n>0)个字节才对齐，那么填充n个字节，每个字节都是n；如果数据本身就已经对齐了，则填充一块长度为块大小的数据，每个字节都是块大小。

举个例子最直观，这里以块大小为16字节，进行PKCS7填充，如下：
数据1： {
0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A}
填充后：{
0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06}
尾部填充了6个6，填充后数据长度为16。

数据2： {
0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10}
填充后：{
0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10}
尾部填充了16个16，填充后数据长度为32。



2）实现ECB模式加解密

加密过程：先通过AES_set_encrypt_key函数生成加密key，然后将输入数据进行PKCS7填充，保证得到16字节整数倍明文，ECB模式密文长度等于明文长度，
        故将输出缓冲out调整为明文大小，以保存密文数据，最后，将明文按16字节，分块进行加密。

解密过程：先通过AES_set_decrypt_key函数生成解密key，由于输入数据，即密文本身就是16字节整数倍，故直接按16字节进行分块解密，
        并将解密后的数据，进行去除填充，得到真正的明文。

经测试，本函数支持对任意长度输入数据进行加解密。


