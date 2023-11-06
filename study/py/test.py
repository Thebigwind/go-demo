import os
def test_exec():
        command = "ssh 10.10.10.174 ls /data/qskm/deploy/rotation_key_threshold.key"
        print(command)
        result = os.system(command)
        # 检查执行结果
        if result == 0:
            print("命令成功执行")
        else:
            print("命令执行失败")

test_exec()
