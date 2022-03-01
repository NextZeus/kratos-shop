
# command

```text

xiaodong@bogon ~/g/s/g/n/kratos-shop (main)> mkdir -p api/cart/service/v1
xiaodong@bogon ~/g/s/g/n/kratos-shop (main)> kratos proto add api/cart/service/v1/cart.proto
xiaodong@bogon ~/g/s/g/n/kratos-shop (main)> kratos proto add api/cart/service/v1/cart_error.proto
xiaodong@bogon ~/g/s/g/n/kratos-shop (main)> mkdir -p app/cart/
xiaodong@bogon ~/g/s/g/n/kratos-shop (main)> cd app/cart/
xiaodong@bogon ~/g/s/g/n/k/a/cart (main)> kratos new service
🚀 Creating service service, layout repo is https://github.com/go-kratos/kratos-layout.git, please wait a moment.

Already up to date.

CREATED service/.gitignore (528 bytes)
CREATED service/Dockerfile (459 bytes)
CREATED service/LICENSE (1066 bytes)
CREATED service/Makefile (2089 bytes)
CREATED service/README.md (1062 bytes)
CREATED service/api/helloworld/v1/error_reason.pb.go (5439 bytes)
CREATED service/api/helloworld/v1/error_reason.proto (407 bytes)
CREATED service/api/helloworld/v1/error_reason_errors.pb.go (952 bytes)
CREATED service/api/helloworld/v1/greeter.pb.go (8229 bytes)
CREATED service/api/helloworld/v1/greeter.proto (692 bytes)
CREATED service/api/helloworld/v1/greeter_grpc.pb.go (3449 bytes)
CREATED service/api/helloworld/v1/greeter_http.pb.go (2139 bytes)
CREATED service/cmd/service/main.go (1709 bytes)
CREATED service/cmd/service/wire.go (569 bytes)
CREATED service/cmd/service/wire_gen.go (1037 bytes)
CREATED service/configs/config.yaml (266 bytes)
CREATED service/generate.go (56 bytes)
CREATED service/go.mod (238 bytes)
CREATED service/go.sum (18354 bytes)
CREATED service/internal/biz/README.md (6 bytes)
CREATED service/internal/biz/biz.go (128 bytes)
CREATED service/internal/biz/greeter.go (683 bytes)
CREATED service/internal/conf/conf.pb.go (21314 bytes)
CREATED service/internal/conf/conf.proto (758 bytes)
CREATED service/internal/data/README.md (7 bytes)
CREATED service/internal/data/data.go (469 bytes)
CREATED service/internal/data/greeter.go (511 bytes)
CREATED service/internal/server/grpc.go (816 bytes)
CREATED service/internal/server/http.go (820 bytes)
CREATED service/internal/server/server.go (150 bytes)
CREATED service/internal/service/README.md (10 bytes)
CREATED service/internal/service/greeter.go (860 bytes)
CREATED service/internal/service/service.go (136 bytes)
CREATED service/openapi.yaml (941 bytes)
CREATED service/third_party/README.md (14 bytes)
CREATED service/third_party/errors/errors.proto (411 bytes)
CREATED service/third_party/google/api/annotations.proto (1051 bytes)
CREATED service/third_party/google/api/client.proto (3395 bytes)
CREATED service/third_party/google/api/field_behavior.proto (3011 bytes)
CREATED service/third_party/google/api/http.proto (15140 bytes)
CREATED service/third_party/google/api/httpbody.proto (2671 bytes)
CREATED service/third_party/google/protobuf/descriptor.proto (38027 bytes)
CREATED service/third_party/validate/README.md (81 bytes)
CREATED service/third_party/validate/validate.proto (31270 bytes)

🍺 Project creation succeeded service
💻 Use the following command to start the project 👇:

$ cd service
$ go generate ./...
$ go build -o ./bin/ ./...
$ ./bin/service -conf ./configs

🤝 Thanks for using Kratos
📚 Tutorial: https://go-kratos.dev/docs/getting-started/start

xiaodong@bogon ~/g/s/g/n/k/a/cart (main)> cd ..
xiaodong@bogon ~/g/s/g/n/k/app (main)> cd ..
xiaodong@bogon ~/g/s/g/n/kratos-shop (main)> kratos proto client api/cart/service/v1/cart.proto
proto: api/cart/service/v1/cart.proto
xiaodong@bogon ~/g/s/g/n/kratos-shop (main)> kratos proto server api/cart/service/v1/cart.proto -t app/cart/service/internal/service
app/cart/service/internal/service/cart.go
xiaodong@bogon ~/g/s/g/n/kratos-shop (main)> rm -rf app/cart/service/api
xiaodong@bogon ~/g/s/g/n/kratos-shop (main)> rm -rf app/cart/service/third_party/
xiaodong@bogon ~/g/s/g/n/kratos-shop (main)> rm -rf app/cart/service/go.mod
xiaodong@bogon ~/g/s/g/n/kratos-shop (main)> rm -rf app/cart/service/go.sum
xiaodong@bogon ~/g/s/g/n/kratos-shop (main)> mv app/cart/service/cmd/service app/cart/service/cmd/server

```