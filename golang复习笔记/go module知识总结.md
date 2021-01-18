**go module知识总结**

参考：https://eddycjy.com/posts/go/gin/2018-02-10-install/

>```go
>//go get
>//拉取最新的版本(优先择取 tag)：go get golang.org/x/text@latest
>//拉取 master 分支的最新 commit：go get golang.org/x/text@master
>//拉取 tag 为 v0.3.2 的 commit：go get golang.org/x/text@v0.3.2
>//拉取 hash 为 342b231 的 commit，最终会被转换为 v0.3.2：go get golang.org/x/text@342b2e
>
>//go mod
>//go mod download 下载 go.mod 文件中指明的所有依赖
>//go mod tidy 整理现有的依赖
>//go mod graph 查看现有的依赖结构
>//go mod init 生成 go.mod 文件 (Go 1.13 中唯一一个可以生成 go.mod 文件的子命令)
>//go mod edit 编辑 go.mod 文件
>//go mod verify 校验一个模块是否被篡改过
>
>//go.sum
>//go.sum 文件详细罗列了当前项目直接或间接依赖的所有模块版本，并写明了那些模块版本的 SHA-256 哈希值以备 Go 在今后的操作中保证项目所依赖的那些模块版本不会被篡改
>
>//go.mod 动词
>//module：用于定义当前项目的模块路径。
>//go：用于设置预期的 Go 版本。
>//require：用于设置一个特定的模块版本。
>//exclude：用于从使用中排除一个特定的模块版本。
>//replace：用于将一个模块版本替换为另外一个模块版本。
>
>
>```
>
>
>
>