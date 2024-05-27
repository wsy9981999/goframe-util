package cfg

import (
	"github.com/gogf/gf/v2/util/gconv"
)

func GetStrs(opt *Option, pattern string, defaultValue ...[]string) ([]string, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return nil, err
	}
	return get.Strings(), nil
}
func GetInts(opt *Option, pattern string, defaultValue ...[]int) ([]int, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return nil, err
	}
	return get.Ints(), nil
}
func GetMaps(opt *Option, pattern string, defaultValue ...[]map[string]interface{}) ([]map[string]interface{}, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return nil, err
	}
	if opt != nil {
		return get.Maps(opt.MapOption), nil
	}
	return get.Maps(), nil
}
func GetSlice(opt *Option, pattern string, defaultValue ...[]interface{}) ([]interface{}, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return nil, err
	}
	return get.Slice(), nil
}
func GetStructs(opt *Option, pattern string, pointer any, m map[string]string, defaultValue ...map[any]any) error {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return err
	}
	return get.Structs(pointer, m)
}
func GetBytes(opt *Option, pattern string, defaultValue ...[]byte) ([]byte, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return nil, err
	}
	return get.Bytes(), err
}
func GetFloat64s(opt *Option, pattern string, defaultValue ...[]float64) ([]float64, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return nil, err
	}
	return get.Float64s(), err
}
func GetFloat32s(opt *Option, pattern string, defaultValue ...[]float32) ([]float32, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return nil, err
	}
	return get.Float32s(), err
}
func GetUints(opt *Option, pattern string, defaultValue ...[]uint) ([]uint, error) {
	str := gconv.Interfaces(defaultValue)
	get, err := Get(opt, pattern, str...)
	if err != nil {
		return nil, err
	}
	return get.Uints(), err
}
