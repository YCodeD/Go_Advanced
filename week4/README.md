### Week04 作业题目
1. 按照自己的构想，写一个项目满足基本的目录结构和工程;
2. 代码包含数据层、业务层、API注册;
3. main函数对于服务的注册和启动，信号处理;
4. wire 构建依赖


### 要点
1. 项目大体结构
    internal 使用，biz、data、service 等目录，可以携带myapp应用名或者不带（比如单项目），internal/pkg 为项目里共用
    cmd/myapp，cmd下需要带上app名字
    api/serivice/v1，按照gRPC 服务名，以及版本号
    configs，放配置文件
2. 对象初始化，biz、data、service，依赖的对象必须作为参数传入，在main里使用wire构建（IoC使用google wire 实现）和消费资源
3. biz中定义了 repository 的接口，实现在 data 目录中，biz中包含了 DomainObject的定义
4. main.go 中使用wire 进行对象后，有lifecycle进行服务的注册和启动