# gluster-storage-gateway

## 架构概览
![event-module](./document/gluster-oss-arc.jpg)
## 功能介绍
- 以glusterfs为存储底座，gluster提供无副本、单副本、EC等实现方式，基于这个底座实现对象存储的功能，
- 添加元数据中心，用户所有的元数据查询从元数据中心获取
- 提供小文件合并的方案
