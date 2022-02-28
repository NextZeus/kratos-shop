//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"

	"github.com/nextzeus/kratos-shop/app/user/service/internal/conf"
	"github.com/nextzeus/kratos-shop/app/user/service/internal/data"
	"github.com/nextzeus/kratos-shop/app/user/service/internal/biz"
	"github.com/nextzeus/kratos-shop/app/user/service/internal/server"
	"github.com/nextzeus/kratos-shop/app/user/service/internal/service"
)

// initApp init kratos application.
// wire injector
// 通过 wire 自动生成一个完整的 initApp 函数, 在 wire_gen.go
func initApp(*conf.Server, *conf.Registry, *conf.Data, *conf.Auth, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
