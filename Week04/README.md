## Go工程化

### 工程项目结构

https://github.com/golang-standards/project-layout/blob/master/README_zh.md

- /cmd

  项目主干。负责程序的启动、关闭、配置初始化，每个应用程序的目录名应该与想要的可执行文件的名称相匹配，例如`/cmd/myapp`。

- /internal

  私有应用程序和库代码。这个布局模式是由 Go 编译器本身执行的，并不局限于顶级 `internal` 目录，在项目树的任何级别上都可以有多个内部目录

- pkg

  外部应用程序可以使用的库代码

- api

  API 协议定义目录。例如grpc的pb文件及生成的代码

- configs

  配置文件模板或默认配置

- test

  额外的外部测试应用程序和测试数据。可以根据项目的大小、测试的需求创建多个子目录

**项目中不应该包含 /src 目录**

### Wire

go依赖注入工具：https://blog.golang.org/wire

### API设计

1.建立公司一个公共的api仓库，方便跨部门协作。 

- 版本管理，基于 git 控制
- 规范化检查，API lint
- API design review，变更 diff
- 权限管理，目录 OWNERS
- 项目中定义 proto，以 api 为包名根目录
- 在统一仓库中管理 proto ，以仓库为包名根目录