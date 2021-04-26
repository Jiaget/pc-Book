# key is of invalid type
今天写了一个用JWT的对称算法生成token 的demo。

我拿到一个 `key is of invalid type` 的错误。

一开始还好奇，go的map会有这种报错？（因为是demo，用户信息写在内存而不是数据库）

ok 百度一下，原来这是JWT的经典报错，但还是不知所云。github的 issue 虽然也报这个错，但那是`[]byte`类型问题，和我无关。

于是乎在我的`jwt_manager.go`中追报错的根源，最终发现是因为我没有提供一个私钥。

Emm....

没错，算法用错了 `SigningMethodHS256` 才是 `SigningMethodHMAC` 的对称密钥算法， 但是我用的是`SigningMethodES256`。它们只差了一个字母。

但这种设计确实会令人非常迷惑，并且报的错也让我很难排查出错误的根源。

JWT是真的难用。。。。