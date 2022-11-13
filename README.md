# MyLog



### 功能实现

用户模块实现了登录，注册 。日志模块实现了增删改查。内部实现了mysql，redis数据处理，账户检测，跨域，token鉴权，grpc协议编程，熔断处理，响应封装，错误处理，配置及实时监听，日志处理，docker部署等功能

### 技术应用

前端：暂无前端，有兴趣的可以自己做做
后端:：mysql、redis数据库。gin、gorm，gRPC框架。Protubuf文件生成，etcd服务发现，hystrix熔断处理，viper配置管理，logurs、rotatelogs日志管理、cors跨域、jwt跨域鉴权、session存储，bcrypt加密等。

### 安装教程

#### docker下的安装
需要安装docker，
并在docker中拉取启动mysql:8.0.28:并创建数据库mylog；
拉取redis:5.0.14 

修改config下的config.yaml文件
进入文件所在的文件夹，执行docker build -f Dockerfile -t mylog .
执行 docker run -p 9003:9003 -d mylog


#### 直接安装访问
需要配置go环境，并安装mysql：8.0.28，创建数据库mylog
安装redis:5.0.14 

修改config.yaml文件
执行命令go mod tidy下载go依赖包
go run main.go执行


### 使用说明

需要在windows下或者docker下安装mysql，redis
如果主机名不是localhost，端口号想修改，需要修改config/config.yaml

swagger导入依赖后，swag init初始化 进入路由界面即可http://localhost:9003/swagger/index.html
授权token格式 Bearer+空格+登录返回reponse数据中token中的字符串

### 参与贡献

Fork 本仓库
新建 Feat_xxx 分支
提交代码
新建 Pull Request

### 项目实例


