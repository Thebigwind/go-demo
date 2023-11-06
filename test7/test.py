#pip install cryptography

from cryptography.hazmat.primitives import padding
from cryptography.hazmat.primitives.kdf.pbkdf2 import PBKDF2HMAC
from cryptography.hazmat.primitives import hashes
from cryptography.hazmat.primitives import serialization
from cryptography.hazmat.primitives.asymmetric import rsa
from cryptography.hazmat.backends import default_backend
from cryptography.hazmat.primitives import padding
from cryptography.hazmat.primitives.ciphers import Cipher, algorithms, modes

# 生成AES密钥
def generate_aes_key(password, salt):
    kdf = PBKDF2HMAC(
        algorithm=hashes.SHA256(),
        iterations=100000,
        salt=salt,
        length=32,
        backend=default_backend()
    )
    return kdf.derive(password.encode())

# 加密
def aes_encrypt(key, plaintext):
    iv = b'\x00' * 16  # 初始化向量，随机生成更安全
    cipher = Cipher(algorithms.AES(key), modes.CFB(iv), backend=default_backend())
    encryptor = cipher.encryptor()
    ciphertext = encryptor.update(plaintext) + encryptor.finalize()
    return iv + ciphertext

# 解密
def aes_decrypt(key, ciphertext):
    iv = ciphertext[:16]
    ciphertext = ciphertext[16:]
    cipher = Cipher(algorithms.AES(key), modes.CFB(iv), backend=default_backend())
    decryptor = cipher.decryptor()
    plaintext = decryptor.update(ciphertext) + decryptor.finalize()
    return plaintext

if __name__ == "__main__":
    password = b"supersecretpassword"
    salt = b"salt"
    plaintext = b"Hello, AES encryption and decryption!"

    key = generate_aes_key(password, salt)
    encrypted_data = aes_encrypt(key, plaintext)
    decrypted_data = aes_decrypt(key, encrypted_data)

    print("Plaintext:", plaintext)
    print("Encrypted Data:", encrypted_data)
    print("Decrypted Data:", decrypted_data)
