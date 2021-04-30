# 使用OPENSSL为项目生成并签名SSL/TLS证书

- 生成CA私钥并签名证书
  - `openssl req -x509 -newkey rsa:4096 -days 365 -keyout ca-key.pem -out ca-cert.pem` 
  - `-x509` 是证书标准格式，因为使用了这个flag, 这条命令将会生成证书而不是`request`
  - `-newkey rsa:4096` 用RSA算法生成4096Bit的密钥
  - `-keyout` 密钥输出文件
  - `-out` 证书输出文件
    ```
    Enter PEM pass phrase:
    Verifying - Enter PEM pass phrase:
    ```
  - 输入上述指令后会要求输出密码。当黑客窃取证书文件后，如果不知道设置的密码，将无法解密
  - 之后`openssl`会要求一些身份信息, 也可以使用`-subj`自动添加这些信息
    - 国家代码, 省， 城市， 组织， 单位， 域名， email
  - 如何阅读加密后的文件
    - `openssl x509 -in ca-cert.pem -noout -text`

- 生成服务端私钥和证书为请求(CSR)签名
  - 与上类似，移除`x509` 因为不需要生成证书，这次需要生成CSR, 移除days,并修改文件与身份信息即可。
- 使用CA私钥给服务端的的CSR签名，返回签署好的证书
  - `openssl x509 -req -in server-req.pem -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem`
    - `-CAcreateserial` 如果CA 序列号为空，自动生成
  - `exfile` 为证书添加多域名

最后， `-nodes` 可以不需要输入密码生成密钥，证书。 仅用来测试使用

## 验证
`openssl verify -CAfile ca-cert.pem server-cert.pem`

[openssl参考文档](https://www.openssl.org/docs/manmaster/man1/req.html)