package goframe_utils

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gmode"
)

func SetDbDebug(env []string, name ...string) error {
	var cfg gdb.ConfigNode
	var _name = "default"
	if len(name) > 0 {
		_name = name[0]
	}
	err := ScanWithDefault("database."+_name, &cfg, nil)
	if err != nil {
		return err
	}
	if gstr.InArray(env, gmode.Mode()) {
		cfg.Debug = true
	}
	gdb.AddDefaultConfigNode(cfg)
	return nil
}
func SetDbDebugWithDefault(name ...string) error {
	return SetDbDebug([]string{gmode.TESTING, gmode.DEVELOP}, name...)
}
