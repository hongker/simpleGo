# Simple Web FrameWork of Golang

## 说明
  我就想整一个golang的web框架，借此了解golang的项目开发流程

## 用法

### 路由
  用法
  ```
  router := vendor.NewRouter()

  // 添加根目录
	router.RegisterController("/", &controllers.IndexController{})

  // 添加子模块的方法
	router.RegisterController("/user/hello", &controllers.UserController{})

  // 传递参数
	router.RegisterController("/user/info/:id", &controllers.UserController{})
  ```
