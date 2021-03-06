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

## AutoPass和SafePass

* AutoPass会自动尝试解密密码。如果无法解密，会返回原始密码。如果能解密，会返回解密后密码。
* SafePass需要在加密后密码前面加前缀.[~。
  * 如果没有前缀，返回原始密码。
  * 如果有前缀，去掉前缀后解密。如果无法解密，返回原始密码。
  * 如果一切顺利，返回解密后密码。
  * .[~前缀不会在辅助工具中自动加上，需要手工添加。
* 两者都会缓存解密结果，对于同样的密码解密不会重复运算。
