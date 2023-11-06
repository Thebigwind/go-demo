import os

#print("return code:", return_code)
tmp = os.popen('./zdlzctl cipher aes --encode 123456').readlines()
print("tmp:")
print(tmp[0])