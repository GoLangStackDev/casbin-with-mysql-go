# Go语言 casbin 系列教程

稍微大一点的项目就一定会涉及到权限管理，这个系列我们给大家介绍一个非常好用的鉴权库 casbin。

![casbin](https://casbin.org/img/casbin.svg)

它是一个跨平台的库，支持很多种语言的鉴权。

官方文档： [https://casbin.org/docs/zh-CN/overview](https://casbin.org/docs/zh-CN/overview)

光讲解这个库的 API 很显然是枯燥的，所以我会结合 Gin，模拟实际开发中的鉴权去给大家讲解他的使用。

所以我们的工程里面会用到到的库如下：

```bash
go get github.com/casbin/casbin/v2
go get -u github.com/gin-gonic/gin
```

## 这个系列你能学到什么？

我在出这个系列文章的时候，最先考虑的就是这个问题，所以我给下列的答案：

- 我们的文章是围绕 RBAC 权限管理开展的，所以如果你要学习 ABAC, ACL 这类的，将不能满足你的需求
- 你能很清楚的理解 casbin 的配置文件
- 你能比较熟悉 casbin 库的常用API
- 你将拥有能把 casbin 整合到 gin 工程里面
- 你能掌握基于 mysql 来控制你的权限
- ...

## 我们的官方公众号

![GoLang全栈](https://static.golangstack.com/%20upload/qrcode_for_gh_e41ae96a4b33_258.jpg)

## 系列文章链接

- [不会处理鉴权？那用Casbin吧，快速了解入门 ](http://mp.weixin.qq.com/s?__biz=MzAxMDM4OTE4Ng==&mid=2247484735&idx=1&sn=285422f74c6675a3f97bb9b5d5a9719e&chksm=9b505692ac27df84b57b41a4459c5bac23af639928b1c5a5719519c146ac5618a293fe0f9f19#rd)
- [面试官：说说Casbin配置文件里的设计哲学（配置详解）](http://mp.weixin.qq.com/s?__biz=MzAxMDM4OTE4Ng==&mid=2247484736&idx=1&sn=c4cb73640b090da4cf22c147a10fda27&chksm=9b5056edac27dffb5203a9f63b0fd83464556309e37d30f699d0d9062a3e5e4da166a75a8779#rd)
- [如何把Casbin整合到Gin里面，让他支持RBAC鉴权？](http://mp.weixin.qq.com/s?__biz=MzAxMDM4OTE4Ng==&mid=2247484768&idx=1&sn=f83bac92fdb71106f9173841de0defa5&chksm=9b5056cdac27dfdb66fef9cd40e264d79e44c4c7549d380fe40e8b53ca8f2a672cc807be95da#rd)
- [如何使用数据库来配置Casbin，手把手教你整合GORM适配器](http://mp.weixin.qq.com/s?__biz=MzAxMDM4OTE4Ng==&mid=2247484800&idx=1&sn=2a779682f2edccc59c319ea367380ca0&chksm=9b50562dac27df3b9cf155fa0fbc41a9152df5ca5e23a2da2b261738d1b01608957d487cb808#rd)


