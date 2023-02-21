# douyin

## **本项目是在window系统上，使用WSL2进行linus环境下的开发**
忘记wsl密码：  
https://blog.csdn.net/weixin_44237337/article/details/119221921?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522167481407316800182760405%2522%252C%2522scm%2522%253A%252220140713.130102334..%2522%257D&request_id=167481407316800182760405&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~sobaiduend~default-2-119221921-null-null.142^v71^pc_new_rank,201^v4^add_ask&utm_term=wsl%E5%BF%98%E8%AE%B0%E5%AF%86%E7%A0%81&spm=1018.2226.3001.4187

# 1. go环境配置
`vim ~/.zshrc`
```
export GOROOT=/usr/local/src/go
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin
export PATH=$GOPATH/bin:$PATH
```
运行`source ~/.zshrc`  
创建终端都需要重新运行
`go env`设置

```
go env -w GO111MODULE=on // 开启go mod
go env -w GOPROXY=https://goproxy.cn,direct // 切换下载代理
```
注意：权限问题，需要root权限
`chmod -R 777 /home/guo/go`
这个是我的项目配置路径
# 2. docker使用
`docker-compose up`  
docker端口占用问题：  
https://blog.csdn.net/u012558210/article/details/127999746?ops_request_misc=&request_id=&biz_id=102&utm_term=Error%20response%20from%20daemon:%20Po&utm_medium=distribute.pc_search_result.none-task-blog-2~all~sobaiduweb~default-0-127999746.nonecase&spm=1018.2226.3001.4187  
```
net stop winnat 
net start winnat
```
其中，使用这个命令可能导致某些命令无法使用，如`git push`，`go mod`等问题，建议最后在使用docker容器
# 3. kitex目录生成
字节example：  
根目录下:
`kitex -module github.com/cloudwego/kitex-examples/bizdemo/easy_note idl/user.thrift`  

cd cmd/user:
`kitex -module github.com/cloudwego/kitex-examples/bizdemo/easy_note -service userdemo -use github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen ../../idl/user.thrift`  
如果需要idl文件格式检验
`go install github.com/cloudwego/thrift-gen-validator@latest`




