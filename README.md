# 2018 冰岩程序组夏令营

### 前言：

+ 不会的话先 Google ，搜索完之后还是没有解决再来问我们

+ 第一天先 fork 仓库：[bingyan-summer-camp2018](https://github.com/tofar/bingyan-summer-camp2018)，本次夏令营要求代码，日报周报等全部托管在 你们 fork 之后的 github 仓库上

  日报、周报不需要太多，只需要介绍每天学习了什么即可

+ 之后的 code review，采用 PR --> MR 的形式

  当你需要代码写了一部分的时候，可以向你们fork的仓库提交 pull request，我们会审核代码，如果有点问题，会附上建议，你们修改好之后我们再将你们的 PR merge 进我们仓库

+ 在夏令营开始的时候会给你们每人分配一个合适的导师

+ 我们每周一次组内分享，欢迎新人做分享

+ 没做完，可以夏令营结束之后继续写，欢迎来问我们

+ 坚持就是胜利，我希望你们在最后还能保持着刚刚进来的热情，很多事情没有你想象的那么难，当然也没有你想象的那么简单，但是很难的事情也是一步一步做完的，希望冰岩夏令营能成为你成为你大学的一个契机、一个跳板。加油！

### 操作说明：

- 首先 fork [此仓库](https://github.com/tofar/bingyan-summer-camp2018)

- 在你的仓库操作的时候请不要对他人目录进行任何操作

- 你的操作权限仅限于你的目录，目录名字为你的 githubID，若仓库中没有你的目录请自行创建

- 提交 PR 的时候自行查看是否存在代码冲突，如果存在自行解决之后再提交 PR

- 提交 PR 是提交到 dev 分支，不是 master 分支

- 提交之后最好跟导师说一声，让导师及时检查

- 目录结构推荐如下：

  + README.md
    必须注明相关开发环境，以及一些特殊说明
  + .gitignore
    忽略一些不需要的文件
  + client
      前端代码，如果不用webpack之类的打包的话，直接一个 dist 文件夹即可
      + dist
        编译打包后的文件目录
        + index.html
        + css
        + js
      + src
        源码
    + server
      服务端代码
      + src
        源码

## 一、准备工作

**注：准备工作主要是作为参考，不做要求，可以跳过，这里涉及到一些你写项目的时候可能会需要的知识**

### 1. 环境准备

+ 推荐使用 Linux or Mac 作为开发系统

+ git

+ 搭建本地开发环境

  高性能反向代理代理服务器软件: Nginx/Openresty

  数据库:MySQL/MongoDB/PostGreSQL

  缓存数据库：Redis/memcached

  语言环境：Python/Go/Java/PHP/node/Kotlin等

  容器工具：Docker，docker-compose

  云服务器：可以选择 阿里云或者腾讯云的学生套餐, 仅10元一月，前期学习只需要在本地即可，后期可能会用到

  注：斜杠划分的选择其一即可

+ 选择好自己适合的开发工具，如：编辑器 vscode，用不惯的话 IDE 亦可

+ 安装一款后台接口测试工具，如：Postman

### 2. 语言 

- Python
  - [廖雪峰的Python3教程](https://www.liaoxuefeng.com/wiki/0014316089557264a6b348958f449949df42a6d3a2e542c000)
        [廖雪峰的Python2教程](https://www.liaoxuefeng.com/wiki/001374738125095c955c1e6d8bb493182103fac9270762a000)
  - [Python3入门指南-官方中文](http://www.pythondoc.com/pythontutorial3/)
        [Python2入门指南-官方中文](http://www.pythondoc.com/pythontutorial27/index.html)
  - [Python 代码规范(Google 开源项目风格指南)](http://zh-google-styleguide.readthedocs.io/en/latest/google-python-styleguide/python_style_rules/) (必需)

- Go
  - [官方链接](https://golang.org/)
  - [官方中文教程](https://tour.go-zh.org/welcome/1)
  - [语言规范](https://go-zh.org/ref/spec)

- PHP

  + [PHP 代码规范](https://www.php-fig.org/psr/) 

- Java
  - [菜鸟教程](http://www.runoob.com/java/java-tutorial.html) 
  - Java 代码规范

- Node

  关键字：单线程，异步，回调地域，Promise，async/await

  - [菜鸟教程-Node](http://www.runoob.com/nodejs/nodejs-install-setup.html)
  - [Node10.5 中文文档](http://nodejs.cn/api/) 
  - [Node 官方文档](https://nodejs.org/api/) 
  - [airbnb node 代码规范](https://github.com/airbnb/javascript) , [node 代码规范](https://github.com/dead-horse/node-style-guide)

- Kotlin
  - [官方教程](http://kotlinlang.org/docs/tutorials/)

### 3. 框架

- Python

  可能会用到 uWSGI 启动

  - Flask
  - Django
  - sanic

- Go

  其实原生的以及封装的很好了

  - Gin
  - Echo
  - Beego

- PHP
   - Laravel

   - CodeIgniter

- Java

  Servlet 写小一点的东西也行

  - SSM + Spring Boot

  - Play

    https://github.com/playframework/playframework

  - Spark

    https://github.com/perwendel/spark/

- Node

  - express
  - Koa

- Kotlin

### 4. 涉及知识

#### 认证：

熟悉以下三种前后端认证方式，一般在登录时使用

- cookie
- session
- JWT

#### 加密算法：

不同需求，对应不同加密方式，先了解

- 对称加密
- 非对称加密
- 哈希算法

### 4. 基本前端知识和前后端交互

注：初学前端请使用 Chrome

- 基本 HTML、CSS 知识

  **关键字：**

  + HTML：布局，表格表单，区块，元素，列表
  + CSS：盒模型，样式表，选择器，浮动，定位

- 基本 JavaScript

  **关键字：事件，DOM操作，作用域，变量，函数**

  注：前端可使用相关UI框架和 JS 框架（react，vue），以及一些比较好用的 js 包（如请求包 axios），可自行选择，具体可以咨询前端组的同学。

- HTTP基本知识

  如：HTTP 方法：GET、POST、PUT、UPDATE等，HTTP状态码：404，500， 200，301等，HTTP URL，HTTP 基本传输格式：json，form等

  **关键字：HTTP 方法，HTTP状态码，HTTP传输格式，HTTP头部**

- 前后端如何交互

  如：前端如何获取后端返回的数据，如何发送请求，后端如何根据前端发过来的请求，回应请求，如何辨别不同的请求

  **关键字： js 请求库(axios, fetch, superagent选其一，原生亦可)，URL， 域名，ip**

### 5. 数据库

学习基本操作（增删改查）即可，以后可自行研究

- MySQL/MongoDB/PostGreSQL
- Redis

可在菜鸟教程上速成

## 二、热身

注：前端能看就行，不是硬性要求，但是要求采用前后端分离的方式，拒绝后端渲染

#### 简易成员管理系统：

具体要求：

+ 管理员登录/注册

  管理员与普通成员信息类似

  注：登录之后注意用户认证问题，如：从浏览器退出此页面之后，再次进入页面如何辨认此用户，登录过期等问题

+ 普通成员注册

  注册之后，需要管理员审核通过才能成为组员

+ 查看未审核的成员，审核成员注册是否通过

+ 添加成员

  成员必须信息如下：用户ID（字符串），密码（要求在数据库中加密存储），邮箱，手机号，昵称，组别

  其他信息自行思考

+ 删除成员

+ 修改成员信息

+ 获取所有成员信息

+ 可以根据组别显示成员

**关键字： 认证，数据库，成员管理**

## 三、项目

注：前端能看就行，不是硬性要求，但是要求采用前后端分离的方式，拒绝后端渲染

### 商城系统：

形式不定：网页，小程序，桌面程序均可

基本功能：

+ 登录注册

  用户密码使用不可逆加密

+ 商品按照类别查询

  如：商品类别：电子设备、书籍资料、宿舍百货、美妆护肤、女装、男装、鞋帽配饰、门票卡券、其他

+ 商品按照地域查询

  如：韵苑、沁苑、紫菘、其他

+ 热门查询、最新查询

  热门查询可在后台记录用户的浏览数据等信息

+ 商品页面

  + 商品详细信息

    标题、简介、价格等

  + 图片

    图片可以存在本地，或者使用七牛云存储

+ 个人信息页

  + 个人基本信息
  + 浏览量

进阶功能：

+ 图片压缩

  浏览时显示压缩的小图片，详细页显示大一点的图片

+ 收藏夹

+ 商品浏览量、收藏量等

+ 后台系统

  + 商品上架、下架
  + 商品信息变动
  + 系统通知

+ 消息提醒

  如：降价提醒、系统推送

+ 接入微信或者QQ

## 四、项目部署

### 1. 配置nginx

学习配置 nginx 做中间代理层，具体可从以下链接中选取部分学习，作为示例，夏令营之后可以好好研究，当然夏令营期间有时间也可以自行研究，遇到坑可以问我们。

[nginx 配置简介](https://juejin.im/post/5ad96864f265da0b8f62188f) 

[openresty 实践](https://juejin.im/post/5aae659c6fb9a028d375308b)

### 2. 配置 docker

[Docker 从入门到实践](https://yeasy.gitbooks.io/docker_practice/content/install/ubuntu.html) 

[Docker 实践](https://juejin.im/post/5b34f0ac51882574ec30afce) 

### 3. 配置域名https (不要求)

前提：有已经备案的域名，有服务器

[Let's Encrypt 给网站加 HTTPS 完全指南](https://ksmx.me/letsencrypt-ssl-https/?utm_source=v2ex&utm_medium=forum&utm_campaign=20160529) 

## 五、附录

### 1. 夏令营聚合

+ 冰岩程序组夏令营：https://github.com/tofar/bingyan-summer-camp2018
+ 冰岩前端组夏令营：https://github.com/BingyanStudioFE/summber-camp-2018
+ 冰岩移动组夏令营: https://github.com/Liujiaohan/bingyan-summer-camp2018
+ 冰岩产品组夏令营：
+ 冰岩运营组夏令营：
+ 冰岩设计组夏令营:
+ 冰岩游戏组夏令营：https://github.com/CurryPseudo/bingyan-summer-camp-2018

### 2. 书籍推荐

#### Python：

- [Python 基础教程](https://zhiguangxiong.gitbooks.io/python/content/di-2-zhang-lie-biao-he-yuan-zu/tong-yong-xu-lie-cao-zuo.html) 

  不推荐快速入门

- [简明Python教程](https://bop.mol.uno/)

- [Effective Python 中文版](https://guoruibiao.gitbooks.io/effective-python/content/) [Effective Python 原版](https://hacktec.gitbooks.io/effective-python/content/en/) 

  很高质量的一本书，让你领略 Python 的一些哲学，以及一些 Python 的使用经验，适合有一定基础的 Python 选手

- [Flask Web开发](https://item.jd.com/11594082.html) 

  讲了很多 web 的基础知识，适合 web 入门

#### Go:

- [《The Go Programming Language》中文版](https://www.gitbook.com/book/yar999/gopl-zh/details)

- [《Effective Go》中英双语版](https://www.gitbook.com/book/bingohuang/effective-go-zh-en/details)

- [Go语言实战](http://download.csdn.net/download/truthurt/9858317)

- [Go Web编程](https://wizardforcel.gitbooks.io/build-web-application-with-golang/content/index.html) 

  可以了解基本web开发，推荐入门

- [Go入门指南](https://github.com/Unknwon/the-way-to-go_ZH_CN)

- [雨痕的学习笔记](https://github.com/qyuhen/book)

  一共两本，第二本讲原理多（如：go 的gc, 内存管理等），第一本适合入门

#### PHP:

  - 《深入php面向对象、模式与实践》
  - 《细说PHP》

#### Node: 

+ 深入浅出 node

#### JS:

+ [你不知道的JavaScript](https://github.com/getify/You-Dont-Know-JS) 上、中、下 
+ JS 高级程序设计
+ [ECMAScript 6 入门](http://es6.ruanyifeng.com/) 

#### Java:

+ effective Java
+ Java 核心技术卷
+ 深入理解jvm虚拟机

### 3. 算法任务（夏令营不要求）

只要你是一个程序员，那么算法对你就是必备！

+ 常见排序算法

  如：冒泡、插入、选择、希尔、堆排序、归并、快排

+ 常见散列

  如：分离链接法、开放定址法（线性探测、平方探测、双散列）

+ 常见数据结构

  如：链表、堆栈、队列、树

+ 常见算法了解

  如：贪婪算法、分治算法、动态规划、回溯

  这些算法主要了解其思想

+ 图论算法

  如：最短路径算法、最小生成树、深搜、宽搜

+ 。。。

书籍推荐：

+ 《数据结构与算法分析 C语言描述》黑皮书，质量很高
+ 《算法 第四版》普林斯顿的书，好像还不错
+ 《算法导论》（理论比较多）

### 4. 后台相关学习

+ 面向对象和抽象的思想

+ 熟悉 Github，熟悉 git 版本管理工具的基本命令操作，如：clone, add, commit, push, pull, merge, branch, checkout, tag

+ http 协议，主要阅读《http权威指南》前三章

+ tcp/ip协议，《计算机网络，自顶向下方法》tcp/ip协议至少要知道他的协议栈，每层是干嘛的，tcp连接建立、断开的过程，tcp/udp的区别

+ WebSocket、Socket、TCP、http，http2

+ Linux 常见操作学习，熟悉基本操作, 如：man, ls, mkdir, cd, cp, mv, scp, ssh, rm, ps, cat, head, tail, vim, wget, curl, chmod, chgrp, chown, sudo, grep

+ 了解 linux基本概念：用户组，权限，文件系统，软/硬连接，挂载，启动等

+ 了解 vim 基本操作，毕竟服务器上一般都是用vim操作，没有图像界面

+ 学习 shell 脚本程序

+ 正则表达式

  这个经常用到，不管是那个方向

+ 数据库设计，可以看下《SQL反模式》

+ 常见设计模式，如：MVC模式，装饰器模式

+ 常见安全问题

  常用工具：浏览器 F12，抓包工具：Fiddler， wireshark

  - DDOS攻击
  - 中间人攻击
  - sql注入
  - ip欺骗
  - xss攻击
  - csrf
  - 远程脚本执行

+ 操作系统基础知识

  基本特性：并发、共享、虚拟、异步

  + CPU 进程调度，线程，进程

  + 共享，**信号量，PV操作，锁**

  + 内存管理

    虚拟存储，段页式系统，缺页

  + 文件系统，ELF文件

+ 分布式了解以及应用

  + 数据库分布式

  + 分布式常见问题

    一致性问题，共识算法，FLP 不可能性原理，CAP 原理，ACID 原则，拜占庭问题，可靠性指标

  + 分布式算法

    Paxos 与 Raft

### 5. 相关建议

+ 良好的英文阅读能力对于一个优秀的程序员来说真的很重要，如果可以的话要养成看英文文档的习惯，毕竟很多好的文章、书籍都是国外的，翻译毕竟会有所损失，而且速度更新也没有看英文的快，不要因为觉得自己英语不好就不看英文文档和英文文章，如果只是阅读的话，还是OK的
+ 对于我们来说，技术好不好其实不一定重要，最重要是你有独立解决问题的能力和对 code 的热爱与勇气。
