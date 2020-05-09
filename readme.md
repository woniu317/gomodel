# golang model例子

代码示例，数据库链接未配置

## 简介

### 初始化db链接

service/database/dbclient/engine.go:init

### 每个域设置引擎管理

model/student/engine.go

model/teacher/engine.go

相当于代理层，便于后续迁移库等操作.

### model只需引用engine

例如:
technologyEngine().Get……

### model增加错误处理

例如:
model/teacher/technology.go:technologyErrorProcess