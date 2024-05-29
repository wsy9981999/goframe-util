package goframe_utils

import (
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/wsy9981999/goframe-util/cfg"
	"time"
)

func GetIntWithDefault(pattern string, defaultValue ...int) (int, error) {
	return cfg.GetInt(nil, pattern, defaultValue...)
}
func GetBoolWithDefault(pattern string, defaultValue ...bool) (bool, error) {
	return cfg.GetBool(nil, pattern, defaultValue...)

}
func GetMapWithDefault(pattern string, defaultValue ...map[string]interface{}) (map[string]interface{}, error) {
	return cfg.GetMap(nil, pattern, defaultValue...)
}
func GetMapStrStrWithDefault(pattern string, defaultValue ...map[string]string) (map[string]string, error) {
	return cfg.GetMapStrStr(nil, pattern, defaultValue...)
}
func GetInt8WithDefault(pattern string, defaultValue ...int8) (int8, error) {
	return cfg.GetInt8(nil, pattern, defaultValue...)
}
func GetUint8WithDefault(pattern string, defaultValue ...uint8) (uint8, error) {
	return cfg.GetUint8(nil, pattern, defaultValue...)

}
func GetInt16WithDefault(pattern string, defaultValue ...int16) (int16, error) {
	return cfg.GetInt16(nil, pattern, defaultValue...)

}
func GetUint16WithDefault(pattern string, defaultValue ...uint16) (uint16, error) {
	return cfg.GetUint16(nil, pattern, defaultValue...)
}
func GetInt32WithDefault(pattern string, defaultValue ...int32) (int32, error) {
	return cfg.GetInt32(nil, pattern, defaultValue...)

}
func GetUint32WithDefault(pattern string, defaultValue ...uint32) (uint32, error) {
	return cfg.GetUint32(nil, pattern, defaultValue...)
}
func GetInt64WithDefault(pattern string, defaultValue ...int64) (int64, error) {
	return cfg.GetInt64(nil, pattern, defaultValue...)

}
func GetUint64WithDefault(pattern string, defaultValue ...uint64) (uint64, error) {
	return cfg.GetUint64(nil, pattern, defaultValue...)
}
func GetFloat32WithDefault(pattern string, defaultValue ...float32) (float32, error) {
	return cfg.GetFloat32(nil, pattern, defaultValue...)
}
func GetValWithDefault(pattern string, defaultValue ...interface{}) (interface{}, error) {
	return cfg.GetVal(nil, pattern, defaultValue...)

}
func GetFloat64WithDefault(pattern string, defaultValue ...float64) (float64, error) {
	return cfg.GetFloat64(nil, pattern, defaultValue...)

}
func GetInterfaceWithDefault(pattern string, defaultValue ...interface{}) (interface{}, error) {
	return cfg.GetInterface(nil, pattern, defaultValue...)
}
func GetDurationWithDefault(pattern string, defaultValue ...time.Duration) (time.Duration, error) {
	return cfg.GetDuration(nil, pattern, defaultValue...)
}
func GetMapStrVarWithDefault(pattern string, defaultValue ...map[string]*gvar.Var) (map[string]*gvar.Var, error) {
	return cfg.GetMapStrVar(nil, pattern, defaultValue...)
}
func GetMapStrAnyWithDefault(pattern string, defaultValue ...map[string]interface{}) (map[string]interface{}, error) {
	return cfg.GetMapStrAny(nil, pattern, defaultValue...)
}
func ScanWithDefault(pattern string, pointer interface{}, m map[string]string, defaultValue ...map[any]any) error {
	return cfg.Scan(nil, pattern, pointer, m, defaultValue...)
}
func StructWithDefault(pattern string, pointer interface{}, m map[string]string, defaultValue ...map[any]any) error {
	return cfg.Struct(nil, pattern, pointer, m, defaultValue...)
}
func GetTimeWithDefault(pattern string, format string, defaultValue ...map[any]any) *gtime.Time {
	return cfg.GetTime(nil, pattern, format, defaultValue...)
}
