# 目标

保存密码的时候加密。

原理是在目标机器上保存一个密码文件用来解密。

# 用法

## 设定机器密钥

1. 利用如下指令`head -c 16 /dev/random | base64`，可以获得一个base64后的随机数。
2. 将两行随机数，写入目标机器的/etc/cryptpass.key。
3. 在加密密码时需要这个文件。

## 加密密码

1. `go run epass/main.go`
2. 在一行内，输入密码。
3. 下一行的输出即为加密后的密码。
4. 可以加`-pass`参数来制定所用的密码文件，默认为/etc/cryptpass.key。
5. 原则上支持中文，实际不建议使用。

## 解密密码

1. 机器上有/etc/cryptpass.key。
2. 使用`password, err := cryptpass.DecryptPass(password)`解密。
3. 参考example/main.go。
