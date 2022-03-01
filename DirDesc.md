# 目录结构

```text
.
├── DirDesc.md
├── LICENSE
├── Makefile
├── README.md
├── Tutorial.md
├── api // 下面维护了微服务使用的proto文件以及根据它们所生成的go文件
│    └── user
│        └── service
│            └── v1
│                ├── user.pb.go
│                ├── user.proto
│                └── user_grpc.pb.go
├── app // 应用
│    └── user
│        └── service
│            ├── Dockerfile
│            ├── LICENSE
│            ├── Makefile
│            ├── README.md
│            ├── cmd // 应用启动目录
│            │    └── server
│            │        ├── main.go // 启动文件
│            │        ├── wire.go // 使用wire来维护依赖注入
│            │        └── wire_gen.go
│            ├── configs // 配置文件
│            │    ├── config.yaml
│            │    └── registry.yaml
│            ├── generate.go
│            ├── internal // 该服务所有不对外暴露的代码，通常的业务逻辑都在这下面，使用internal避免错误引用
│            │    ├── biz // 业务逻辑的组装层，类似 DDD 的 domain 层，data 类似 DDD 的 repo，repo 接口在这里定义，使用依赖倒置的原则。
│            │    │    ├── README.md
│            │    │    ├── address.go
│            │    │    ├── biz.go
│            │    │    ├── card.go
│            │    │    └── user.go
│            │    ├── conf // 内部使用的config的结构定义，使用proto格式生成
│            │    │    ├── conf.pb.go
│            │    │    └── conf.proto
│            │    ├── data // 业务数据访问，包含 cache、db 等封装，实现了 biz 的 repo 接口。我们可能会把 data 与 dao 混淆在一起，data 偏重业务的含义，它所要做的是将领域对象重新拿出来，我们去掉了 DDD 的 infra层。
│            │    │    ├── README.md
│            │    │    ├── address.go
│            │    │    ├── card.go
│            │    │    ├── data.go
│            │    │    ├── ent // 使用 ent 生成数据库表 Entity
│            │    │    │    ├── address
│            │    │    │    │    ├── address.go
│            │    │    │    │    └── where.go
│            │    │    │    ├── address.go
│            │    │    │    ├── address_create.go
│            │    │    │    ├── address_delete.go
│            │    │    │    ├── address_query.go
│            │    │    │    ├── address_update.go
│            │    │    │    ├── card
│            │    │    │    │    ├── card.go
│            │    │    │    │    └── where.go
│            │    │    │    ├── card.go
│            │    │    │    ├── card_create.go
│            │    │    │    ├── card_delete.go
│            │    │    │    ├── card_query.go
│            │    │    │    ├── card_update.go
│            │    │    │    ├── client.go
│            │    │    │    ├── config.go
│            │    │    │    ├── context.go
│            │    │    │    ├── ent.go
│            │    │    │    ├── enttest
│            │    │    │    │    └── enttest.go
│            │    │    │    ├── hook
│            │    │    │    │    └── hook.go
│            │    │    │    ├── migrate
│            │    │    │    │    ├── migrate.go
│            │    │    │    │    └── schema.go
│            │    │    │    ├── mutation.go
│            │    │    │    ├── predicate
│            │    │    │    │    └── predicate.go
│            │    │    │    ├── runtime
│            │    │    │    │    └── runtime.go
│            │    │    │    ├── runtime.go
│            │    │    │    ├── schema // mysql 表结构定义
│            │    │    │    │    ├── address.go
│            │    │    │    │    ├── card.go
│            │    │    │    │    └── user.go
│            │    │    │    ├── tx.go
│            │    │    │    ├── user
│            │    │    │    │    ├── user.go
│            │    │    │    │    └── where.go
│            │    │    │    ├── user.go
│            │    │    │    ├── user_create.go
│            │    │    │    ├── user_delete.go
│            │    │    │    ├── user_query.go
│            │    │    │    └── user_update.go
│            │    │    └── user.go
│            │    ├── pkg
│            │    │    └── util
│            │    │        └── password.go
│            │    ├── server // http和grpc实例的创建和配置
│            │    │    ├── grpc.go
│            │    │    └── server.go
│            │    ├── service // 实现了 api 定义的服务层，类似 DDD 的 application 层，处理 DTO 到 biz 领域实体的转换(DTO -> DO)，同时协同各类 biz 交互，但是不应处理复杂逻辑
│            │    │    ├── README.md
│            │    │    ├── service.go
│            │    │    └── user.go
│            │    └── test // 测试目录
│            │        └── user.go
│            └── openapi.yaml
├── go.mod
├── go.sum
├── openapi.yaml
└── third_party // api 依赖的第三方proto
    ├── README.md
    ├── errors
    │    └── errors.proto
    ├── google
    │    ├── api
    │    │    ├── annotations.proto
    │    │    ├── client.proto
    │    │    ├── field_behavior.proto
    │    │    ├── http.proto
    │    │    └── httpbody.proto
    │    └── protobuf
    │        └── descriptor.proto
    └── validate
        ├── README.md
        └── validate.proto

35 directories, 89 files


```

