### 创建标准工程目录：
- api 放置grpc protoc和接口
- cmd 放置启动文件和启动配置
- configs 放置全局配置
- internal
- - service 放置自定义服务
- - dao 放置数据访问逻辑
- - model 放置PO对象 

### 服务说明：
- 使用grpc实现一个rpc服务User
- api实现登录接口Login
- dao使用mysql查询DB
- service接收数据，调用dao查询并处理返回

### 其他：
未使用wire，感觉wire生成代码比较怪异，放置到任何目录都不太合适。
