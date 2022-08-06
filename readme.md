# Ginblog

原创作者仓库：https://gitee.com/wejectchan/ginblog

<div align="center">
<img  src="https://s3.bmp.ovh/imgs/2022/07/31/cc4d0ebbc896a9eb.jpg" width="600" height="350"/>
</div>

## 实现功能

1.  简单的用户管理权限设置
2.  用户密码加密存储
3.  文章分类自定义
4.  列表分页
5.  图片上传七牛云
6.  JWT 认证
7.  自定义日志功能
8.  跨域 cors 设置
9.  文章评论功能

## 项目预览

- 前端展示页面
  ![](https://s3.bmp.ovh/imgs/2022/07/31/05ec23559fec2a13.png)

- 前端展示页面
  ![](https://s3.bmp.ovh/imgs/2022/07/31/233658d4e01e5151.png)

- 后台登录页面

  ![](https://s3.bmp.ovh/imgs/2022/07/31/1c70342bc57c803c.jpg)

- 后台管理页面

  ![](https://s3.bmp.ovh/imgs/2022/07/31/8e42cd300655056b.jpg)

## 技术栈

- golang
  - Gin web framework
  - gorm(v1 && v2)
  - jwt-go
  - scrypt
  - logrus
  - gin-contrib/cors
  - go-playground/validator/v10
  - go-ini
- JavaScript
  - vue
  - vue cli
  - vue router
  - ant design vue
  - vuetify
  - axios
  - tinymce
  - moment
- MySQL 

## 前端环境



后台管理界面采用 Vue2 + Ant-design设计

博客界面采用 Vue2 + Vuetify 设计 

文本编辑区域采用 TinyMCE 框架





<div align="center">
<img  src="https://s3.bmp.ovh/imgs/2022/07/31/cc4d0ebbc896a9eb.jpg" width="600" height="350"/>
</div>

## 目录结构

```shell
├─  .gitignore
│  go.mod // 项目依赖
│  go.sum
│  LICENSE
│  main.go //主程序
│  README.md
│  tree.txt
│          
├─api         
├─config // 项目配置入口   
├─database  // 数据库备份文件（初始化）
├─log  // 项目日志
├─middleware  // 中间件
├─model // 数据模型层
├─routes
│      router.go // 路由入口    
├─static // 打包静态文件
│  ├─admin  // 后台管理页面 (已废弃，打包静态文件在web/admin/dist下)         
│  └─front  // 前端展示页面 (已废弃，打包静态文件在web/front/dist下)            
├─upload   
├─utils // 项目公用工具库
│  │  setting.go 
│  ├─errmsg   
│  └─validator         
└─web // 前端开发源码（VUECLI项目源文件)
    ├─admin             
    └─front
```



## 后端环境

1. 克隆项目

2. 转到下面文件夹下

```shell
cd yourPath/GinBlog
```

3. 安装依赖

```
go mod tidy
```

4. 初始化项目配置config.ini

```ini
./config/config.ini


[server]
AppMode=debug
HttpPort=:8081
JwtKey=8afiehoui
[database]
Db=mysql
DbHost=localhost
DbPort=3306
DbUser=
DbPass=
DbName=ginblog
[qiniu]
AccessKey  =
SecretKey  =
Bucket     =
QiniuServer=
```

5. 在database中将 sql 文件导入数据库

	推荐 navicat 或者其他 sql 管理工具导入

6. 启动项目

```shell
 go run main.go
```

此时，项目启动，你可以访问页面

```shell
首页
http://localhost:8081/front
后台管理页面
http://localhost:8081/admin

默认管理员:wakk1  密码:111222333
```







