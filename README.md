# CrzayORM
## 目前实现
- 引擎和session的封装
- 日志库
- 适配sqlite3数据库，映射数据类型和特定的SQL语句，创建 Dialect 层屏蔽数据库差异。
- 设计Schema，利用(Reflect)完成结构体和数据库表结构的映射
- 构造创建、删除、存在性的SQL语句完成数据库表的基本操作
- 实现记录的新增和查询