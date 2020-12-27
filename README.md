# gluster-storage-gateway


## 架构概览
![event-module](./document/gluster-oss-arc.jpg)
## 功能介绍
- 以glusterfs为存储底座，gluster提供无副本、单副本、EC等实现方式，基于这个底座实现对象存储的功能，
- 添加元数据中心，用户所有的元数据查询从元数据中心获取
- 提供小文件合并的方案

## 网关服务目录结构

- 目录结构
```
root@ubuntu:/mnt/dht# tree ./
./
├── bucket0
│   ├── block
│   │   ├── bucket0.block.0
│   │   ├── bucket0.block.1
│   │   ├── bucket0.block.10
│   │   ├── bucket0.block.11
│   │   ├── bucket0.block.12
│   │   ├── bucket0.block.13
│   │   ├── bucket0.block.14
│   │   ├── bucket0.block.15
│   │   ├── bucket0.block.16
│   │   ├── bucket0.block.17
│   │   ├── bucket0.block.18
│   │   ├── bucket0.block.19
│   │   ├── bucket0.block.2
│   │   ├── bucket0.block.20
│   │   ├── bucket0.block.21
│   │   ├── bucket0.block.22
│   │   ├── bucket0.block.23
│   │   ├── bucket0.block.24
│   │   ├── bucket0.block.25
│   │   ├── bucket0.block.26
│   │   ├── bucket0.block.27
│   │   ├── bucket0.block.28
│   │   ├── bucket0.block.29
│   │   ├── bucket0.block.3
│   │   ├── bucket0.block.30
│   │   ├── bucket0.block.31
│   │   ├── bucket0.block.4
│   │   ├── bucket0.block.5
│   │   ├── bucket0.block.6
│   │   ├── bucket0.block.7
│   │   ├── bucket0.block.8
│   │   └── bucket0.block.9
│   ├── bucket0.block.index
│   └── bucket0.object.index
└── bucket.meta

```

- bucket.meta:整个服务网关的存储bucket meta的文件，每个bucket更改都会持久化到这个文件，都是以append的方式
- block: 每个bucket的所有数据都存储在这个目录中，一个bucket可以创建多个block文件，每个block文件可以写入数据
- bucket0.block.index:这个文件是每个block目录下{bucket-name}.block.{index}数据块的元数据
- bucket0.object.index：这个是每个object的元数据，包括对象名称、属性、大小、所在block的起始位置、终止位置等
