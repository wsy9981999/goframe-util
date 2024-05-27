package log

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
)

type Option struct {
	Name    string
	Context context.Context
}

func getLogInstance(opt *Option) (*glog.Logger, context.Context) {
	var o *Option
	if opt == nil {
		o = &Option{
			Name:    "",
			Context: context.Background(),
		}
	} else {
		o = opt
		if o.Context == nil {
			o.Context = gctx.GetInitCtx()
		}
	}

	if o.Name == "" {
		return g.Log(), o.Context
	}
	return g.Log(o.Name), o.Context
}

func Error(opt *Option, v ...any) {
	instance, ctx := getLogInstance(opt)
	instance.Error(ctx, v...)
}
func Errorf(opt *Option, format string, v ...any) {
	instance, ctx := getLogInstance(opt)
	instance.Errorf(ctx, format, v...)
}
func Fatal(opt *Option, v ...any) {
	instance, ctx := getLogInstance(opt)
	instance.Fatal(ctx, v...)
}
func Fatalf(opt *Option, format string, v ...any) {
	instance, ctx := getLogInstance(opt)
	instance.Fatalf(ctx, format, v...)
}
func Debug(opt *Option, v ...any) {
	instance, ctx := getLogInstance(opt)
	instance.Debug(ctx, v...)
}
func Debugf(opt *Option, format string, v ...any) {
	instance, ctx := getLogInstance(opt)
	instance.Debugf(ctx, format, v...)
}
func Info(opt *Option, v ...any) {
	instance, ctx := getLogInstance(opt)
	instance.Info(ctx, v...)
}
func Infof(opt *Option, format string, v ...any) {
	instance, ctx := getLogInstance(opt)
	instance.Infof(ctx, format, v...)
}
func Warning(opt *Option, v ...any) {
	instance, ctx := getLogInstance(opt)
	instance.Warning(ctx, v...)
}
func Warningf(opt *Option, format string, v ...any) {
	instance, ctx := getLogInstance(opt)
	instance.Warningf(ctx, format, v...)
}
func Panic(opt *Option, v ...any) {
	instance, ctx := getLogInstance(opt)
	instance.Panic(ctx, v...)
}
func Panicf(opt *Option, format string, v ...any) {
	instance, ctx := getLogInstance(opt)
	instance.Panicf(ctx, format, v...)
}
func Notice(opt *Option, v ...any) {
	instance, ctx := getLogInstance(opt)
	instance.Notice(ctx, v...)
}
func Noticef(opt *Option, format string, v ...any) {
	instance, ctx := getLogInstance(opt)
	instance.Noticef(ctx, format, v...)
}
func Critical(opt *Option, v ...any) {
	instance, ctx := getLogInstance(opt)
	instance.Critical(ctx, v...)
}
func Criticalf(opt *Option, format string, v ...any) {
	instance, ctx := getLogInstance(opt)
	instance.Criticalf(ctx, format, v...)
}

func Print(opt *Option, v ...any) {
	instance, ctx := getLogInstance(opt)
	instance.Print(ctx, v...)
}
func Printf(opt *Option, format string, v ...any) {
	instance, ctx := getLogInstance(opt)
	instance.Printf(ctx, format, v...)
}
