//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/nextzeus/kratos-shop/app/cart/service/internal/biz"
	"github.com/nextzeus/kratos-shop/app/cart/service/internal/conf"
	"github.com/nextzeus/kratos-shop/app/cart/service/internal/data"
	"github.com/nextzeus/kratos-shop/app/cart/service/internal/server"
	"github.com/nextzeus/kratos-shop/app/cart/service/internal/service"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
