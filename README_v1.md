# etcd-manage

本项目受到[e3w](https://github.com/soyking/e3w)启发

项目旨在于方便开发者简单明了的查看、添加、修改和删除存储于etcd中的配置。

相比于e3w本程序运行简单，已将ui静态文件打包到了可执行文件中只需在终端运行此程序即可自动打开浏览器访问子集的etcd。

### 与e3w好的地方

1. 支持多etcd服务，更适用于开发者日常开发场景，一般都会有开发环境、测试环境和线上环境的不同切换etcd服务端。
2. 支持Basic Auth用户验证登录，可以通过配置文件添加用户并设置用户角色实现用户可访问etcd server的空。



## Demo

demo地址： [Demo](http://140.143.234.132:10280/ui/#/keys)

用户名：admin

密码：123456


## 使用

可以直接通过docker-compose将服务运行在docker中，可以参考 [https://github.com/shiguanghuxian/docker-compose/tree/master/etcd33](https://github.com/shiguanghuxian/docker-compose/tree/master/etcd33)直接使用。

配置文件介绍 [bin/config/cfg.toml](bin/config/cfg.toml)

配置完后就可以直接在终端运行 `etcd-manage` 


## 备注

如果需要切换用户需要彻底退出浏览器，重新打开管理页面。


## 编译和运行
由于已经将前端静态文件打包到了[tpls/tpls.go](tpls/tpls.go)，如果不需要自定义ui可以不忽略 *前端代码编译* 部分

***服务端编译：***

> 1. 编译到当前平台
> 	
> 	make build
> 2. 编译到其它平台
> 
> 	make linux\_build 或 make windows\_build
> 
> 3. 直接运行
> 
> 	make run
> 
> 4. 安装
> 
> 	make install
> 
> 5. 编译前端代码
> 
> 	make build\_web
> 
> 6. 编译docker
> 
> 	make docker\_build
> 
> 	docker-compose up
> 



***前端代码编译：***

>
> 1. 安装node 
> 
> 		下载地址 [https://nodejs.org/](https://nodejs.org/)
>
> 2. 安装cnpm加速node依赖安装速度
>
> 		npm install -g cnpm  --registry=https://registry.npm.taobao.org
> 
> 3. 安装vue-cli3
> 
> 		cnpm install -g @vue/cli
> 
> 4. 安装依赖，编译项目
>
>  		cd $GOPATH/github.com/shiguanghuxian/etcd-manage/static
> 
> 		cnpm install && npm run build
> 

## 运行效果
首页
![](./images/abc.png)

列表形式显示
![](./images/list.png)

添加
![](./images/add.png)

删除
![](./images/del.png)

查看&修改
![](./images/show.png)
