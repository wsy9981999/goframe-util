package goframe_utils

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gstructs"
)

const (
	metaName = "Meta"
)

func GetMetaTag(r *ghttp.Request) map[string]string {
	name, b := r.GetServeHandler().Handler.Info.Type.In(1).Elem().FieldByName(metaName)
	if !b {
		return nil
	}
	return gstructs.ParseTag(string(name.Tag))
}
func GetMetaFromCtx(ctx context.Context) map[string]string {
	return GetMetaTag(ghttp.RequestFromCtx(ctx))
}
