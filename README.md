# 基于 git 的 Blog
直接读取本地的 md 文件，简单好部署。

# API
* ``` Get / ```    
* ``` GET /blog/list?page=[default 1]&count=[default 20] ```    
* ``` GET /blog/detail/:id ```
* ``` POST /blog/hook (for git hook)```

# Usage
修改 ```main.go```中的 ```Root``` 为自己的 blog 仓库.    
```go build main.go```    
然后启动就好了。    
前端页面可以自己开发或者使用程序自带的。    
