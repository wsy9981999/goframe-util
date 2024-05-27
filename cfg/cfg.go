package cfg

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
)

type Option struct {
	Name      string
	Context   context.Context
	Must      bool
	Cmd       bool
	Env       bool
	MapOption gvar.MapOption
}

func getCfgInstance(opt *Option) (*gcfg.Config, context.Context) {
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
		return g.Cfg(), o.Context
	}
	return g.Cfg(o.Name), o.Context
}
func Get(opt *Option, pattern string, defaultValue ...any) (*gvar.Var, error) {
	instance, ctx := getCfgInstance(opt)
	var m func(ctx context.Context, pattern string, defaultValue ...any) (*gvar.Var, error)
	if opt.Must {
		m = func(ctx context.Context, pattern string, defaultValue ...any) (*gvar.Var, error) {
			if opt.Cmd {
				return instance.MustGetWithCmd(ctx, pattern, defaultValue...), nil
			} else if opt.Env {
				return instance.MustGetWithEnv(ctx, pattern, defaultValue...), nil
			}
			return instance.MustGet(ctx, pattern, defaultValue...), nil
		}

	}
	m = func(ctx context.Context, pattern string, defaultValue ...any) (*gvar.Var, error) {
		if opt.Cmd {
			return instance.GetWithCmd(ctx, pattern, defaultValue...)
		} else if opt.Env {
			return instance.GetWithEnv(ctx, pattern, defaultValue...)
		}
		return instance.Get(ctx, pattern, defaultValue...)
	}
	return m(ctx, pattern, defaultValue...)

}
