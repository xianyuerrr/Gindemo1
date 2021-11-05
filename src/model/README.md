# model层数据格式

## Rule

从http请求中解析的格式，和文档中一一对应

其中，特殊的是版本号，格式为x.x.x.x，并且没有前置0

## NewRules

为了适配数据库中的表单结构，新增加了和whitelist表多对多关系的gorm配置

## clientInfo

