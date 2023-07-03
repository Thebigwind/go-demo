// hello.h
void SayHello(const char* s);

typedef struct CK_INFO_ {
 char sn[32];  //唯一标识
}CK_INFO;


/**
 * @brief 获取当前驱动在线的设备
 *
 * @param ck_info session句柄数组
 * @param size shd数组的长度
 * @return int 0为成功，非0为失败
 */
int CardOrUKeyEnum( CK_INFO** ck_info, unsigned int size);





typedef struct users_info_
{
    char name[32];  //唯一标识
}users_info;

/**
 * @brief 获取已建的用户名列表
 * @param users 用户列表
 * @param users_len 用户列表的长度
 * @return int 0为成功，非0为失败
 */
int GetUsers( users_info** users, unsigned int* users_len);