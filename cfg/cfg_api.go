package cfg

import (
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

func GetStr(opt *Option, pattern string, defaultValue ...string) (string, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return "", err
	}
	return get.String(), nil
}
func GetInt(opt *Option, pattern string, defaultValue ...int) (int, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return 0, err
	}
	return get.Int(), nil
}
func GetBool(opt *Option, pattern string, defaultValue ...bool) (bool, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return false, err
	}
	return get.Bool(), nil
}
func GetMap(opt *Option, pattern string, defaultValue ...map[string]interface{}) (map[string]interface{}, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return nil, err
	}
	if opt != nil {
		return get.Map(opt.MapOption), nil
	}
	return get.Map(), nil
}
func GetMapStrStr(opt *Option, pattern string, defaultValue ...map[string]string) (map[string]string, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return nil, err
	}
	if opt != nil {
		return get.MapStrStr(opt.MapOption), nil
	}
	return get.MapStrStr(), err
}
func GetInt8(opt *Option, pattern string, defaultValue ...int8) (int8, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return 0, err
	}
	return get.Int8(), nil
}
func GetUint8(opt *Option, pattern string, defaultValue ...uint8) (uint8, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return 0, err
	}
	return get.Uint8(), nil
}
func GetInt16(opt *Option, pattern string, defaultValue ...int16) (int16, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return 0, err
	}
	return get.Int16(), nil
}
func GetUint16(opt *Option, pattern string, defaultValue ...uint16) (uint16, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return 0, err
	}
	return get.Uint16(), nil
}
func GetInt32(opt *Option, pattern string, defaultValue ...int32) (int32, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str)
	if err != nil {
		return 0, err
	}
	return get.Int32(), nil
}
func GetUint32(opt *Option, pattern string, defaultValue ...uint32) (uint32, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str)
	if err != nil {
		return 0, err
	}
	return get.Uint32(), nil
}
func GetInt64(opt *Option, pattern string, defaultValue ...int64) (int64, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str)
	if err != nil {
		return 0, err
	}
	return get.Int64(), nil
}
func GetUint64(opt *Option, pattern string, defaultValue ...uint64) (uint64, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str)
	if err != nil {
		return 0, err
	}
	return get.Uint64(), nil
}
func GetFloat32(opt *Option, pattern string, defaultValue ...float32) (float32, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return 0, err
	}
	return get.Float32(), nil
}
func GetVal(opt *Option, pattern string, defaultValue ...interface{}) (interface{}, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return 0, err
	}
	return get.Val(), nil
}
func GetFloat64(opt *Option, pattern string, defaultValue ...float64) (float64, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return 0, err
	}
	return get.Float64(), nil
}
func GetInterface(opt *Option, pattern string, defaultValue ...interface{}) (interface{}, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return 0, err
	}
	return get.Interface(), nil
}
func GetDuration(opt *Option, pattern string, defaultValue ...time.Duration) (time.Duration, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return 0, err
	}
	return get.Duration(), nil
}
func GetMapStrVar(opt *Option, pattern string, defaultValue ...map[string]*gvar.Var) (map[string]*gvar.Var, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return nil, err
	}
	return get.MapStrVar(), nil
}
func GetMapStrAny(opt *Option, pattern string, defaultValue ...map[string]interface{}) (map[string]interface{}, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return nil, err
	}
	return get.MapStrAny(), nil
}
func Scan(opt *Option, pattern string, pointer interface{}, m map[string]string, defaultValue ...map[any]any) error {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return nil
	}
	return get.Scan(pointer, m)
}
func Struct(opt *Option, pattern string, pointer interface{}, m map[string]string, defaultValue ...map[any]any) error {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return nil
	}
	return get.Struct(pointer, m)
}
func GetTime(opt *Option, pattern string, format string, defaultValue ...map[any]any) *gtime.Time {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return nil
	}
	return get.GTime(format)
}
