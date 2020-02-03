package main

import (
	"io"
	"log"

	"github.com/devit-tel/goxid"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"

	"github.com/devit-tel/gogo-blueprint/app"
	"github.com/devit-tel/gogo-blueprint/config"
	companyRepo "github.com/devit-tel/gogo-blueprint/repository/company/store"
	staffRepo "github.com/devit-tel/gogo-blueprint/repository/staff/store"
	companyService "github.com/devit-tel/gogo-blueprint/service/company"
	companyServiceTracer "github.com/devit-tel/gogo-blueprint/service/company/withtracer"
	staffService "github.com/devit-tel/gogo-blueprint/service/staff"
	staffServiceTracer "github.com/devit-tel/gogo-blueprint/service/staff/withtracer"
)

func setupJaeger(appConfig *config.Config) io.Closer {
	cfg, err := jaegercfg.FromEnv()
	if err != nil {
		log.Panic(err)
	}
	cfg.ServiceName = appConfig.AppName
	cfg.Sampler.Type = "const"
	cfg.Sampler.Param = 1
	cfg.Reporter = &jaegercfg.ReporterConfig{LogSpans: true}

	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		log.Panic(err)
	}
	opentracing.SetGlobalTracer(tracer)

	return closer
}

func newApp(appConfig *config.Config) *app.App {
	xid := goxid.New()

	companyStore := companyRepo.New(appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBCompanyTableName)
	company := companyServiceTracer.Wrap(companyService.New(xid, companyStore))

	staffStore := staffRepo.New(appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBStaffTableName)
	staff := staffServiceTracer.Wrap(staffService.New(xid, staffStore, companyStore))

	return app.New(staff, company)
}

func setupLog() *logrus.Logger {
	lr := logrus.New()
	lr.SetFormatter(&logrus.JSONFormatter{})

	return lr
}
