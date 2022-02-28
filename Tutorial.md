# Command

```shell
xiaodong@bogon ~/go > mkdir -p github.com/nextzeus/kratos-shop 
xiaodong@bogon ~/go > cd github.com/nextzeus/kratos-shop 
xiaodong@bogon ~/g/s/g/n/kratos-shop (main)> go mod init github.com/nextzeus/kratos-shop
xiaodong@bogon ~/g/s/g/n/kratos-shop (main)> mkdir -p api
xiaodong@bogon ~/g/s/g/n/kratos-shop (main)> mkdir -p app
xiaodong@bogon ~/g/s/g/n/kratos-shop (main)> cd app/
xiaodong@bogon ~/g/~/g/s/g/n/k/app (main)> kratos new server # (app/server)
xiaodong@bogon ~/g/s/g/n/kratos-shop (main)> rm -rf app/server/api #(删除 api 目录)
xiaodong@bogon ~/g/s/g/n/kratos-shop (main)> rm app/server/internal/service/greeter.go 
xiaodong@bogon ~/g/s/g/n/kratos-shop (main)> mkdir -p api/server/service/v1
# 生成 proto 模板
xiaodong@bogon ~/g/s/g/n/kratos-shop (main)> kratos proto add api/server/service/v1/server.proto
# 生成 proto 源码
xiaodong@bogon ~/g/s/g/n/kratos-shop (main)> kratos proto client api/server/service/v1/server.proto
# 生成 server 模板
xiaodong@bogon ~/g/s/g/n/kratos-shop (main)> kratos proto server api/server/service/v1/server.proto -t app/server/internal/service/
app/server/internal/service/server.go


```

## 修改 app/user/internal/conf/conf.proto
```proto

message Bootstrap {
  ...
  Trace trace = 3;
  Auth auth = 4;
}

message Trace {
  string endpoint = 1;
}

message Auth {
  string key = 1;
}

message Registry {
  message Consul {
    string address = 1;
    string scheme = 2;
  }
  Consul consul = 1;
}

```

## 修改 app/user/configs/config.yaml
```yaml
...
trace: # 链路追踪的参数
  endpoint: http://127.0.0.1:14268/api/traces
auth: # jwt token key
  key: some-secret-key
```
## 修改 app/user/Makefile
- ./third_party 改成 ../../third_party
- 在 app/user 目录下执行 make config 生成 config.pb.go

## 创建 app/user/data/ent 目录
```shell
xiaodong@bogon ~/g/s/g/n/kratos-shop (main)> mkdir -p app/server/internal/data/ent/schema

```

## 新增 ent/schema/user.go
```shell
xiaodong@bogon ~/g/s/g/n/kratos-shop (main)> touch app/server/internal/data/ent/schema/server.go
```
### 填充代码
```go
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("username").Unique(),
		field.String("password_hash"),
		field.Time("created_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("updated_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("addresses", Address.Type),
		edge.To("cards", Card.Type),
	}
}

```
## 同上 添加 address.go card.go
## 安装 ent
```shell
xiaodong@bogon ~/go> go install entgo.io/ent/cmd/ent@latest
```

## 生成 ent 代码
```shell
xiaodong@bogon ~/g/s/g/n/k/a/u/service (main)> make ent
cd internal/data/ && ent generate ./ent/schema
```

## 修改 app/user/service/internal/biz/user.go
- UserUseCase
## 添加 app/user/service/internal/biz/address.go
- AddressUseCase
## 添加 app/user/service/internal/biz/card.go
- CardUseCase
## 修改 app/user/service/internal/biz/biz.go
```go
package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewUserUseCase, NewAddressUseCase, NewCardUseCase)
```
