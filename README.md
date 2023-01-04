# steggo

StegGo want to be StegCloak‘s go version

使用密码安全地把信息隐藏在不可见字符中。它可用于安全地为字符串、网页上的不可见脚本、社交媒体上的文本或任何其他隐蔽通信添加水印。

## 安装

```shell
go get github.com/asjdf/steggo 
```

## 使用

```go
import "github.com/asjdf/steggo"


func main() {
    plain := []byte("steggoo")
    encode, err := Encode(plain)
    decode, err := Decode(encode)
}
```