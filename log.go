package goframe_utils

import "goframe-util/log"

func ErrorWithDefault(v ...any) {
	log.Error(nil, v...)
}
func ErrorfWithDefault(format string, v ...any) {
	log.Errorf(nil, format, v...)
}
func FatalWithDefault(v ...any) {
	log.Fatal(nil, v...)
}
func FatalfWithDefault(format string, v ...any) {
	log.Fatalf(nil, format, v...)
}
func DebugWithDefault(v ...any) {
	log.Debug(nil, v...)
}
func DebugfWithDefault(format string, v ...any) {
	log.Debugf(nil, format, v...)
}
func InfoWithDefault(v ...any) {
	log.Info(nil, v...)
}
func InfofWithDefault(format string, v ...any) {
	log.Infof(nil, format, v...)
}
func WarningWithDefault(v ...any) {
	log.Warning(nil, v...)
}
func Warningf(format string, v ...any) {
	log.Warningf(nil, format, v...)
}
func PanicWithDefault(v ...any) {
	log.Panic(nil, v...)
}
func PanicfWithDefault(format string, v ...any) {
	log.Panicf(nil, format, v...)
}
func NoticeWithDefault(v ...any) {
	log.Notice(nil, v...)
}
func NoticefWithDefault(format string, v ...any) {
	log.Noticef(nil, format, v...)
}
func CriticalWithDefault(v ...any) {
	log.Critical(nil, v...)
}
func CriticalfWithDefault(format string, v ...any) {
	log.Criticalf(nil, format, v...)
}
func PrintWithDefault(v ...any) {
	log.Print(nil, v...)
}
func PrintfWithDefault(format string, v ...any) {
	log.Printf(nil, format, v...)
}
