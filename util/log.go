package util

import (
	"github.com/cihub/seelog"
	"log"
	"sync"
)

type mylog struct {
	File string //普通日志
	Log  seelog.LoggerInterface
}

var (
	my_log     *mylog
	mylog_once sync.Once
)

// Mylog  创建mylog单实例
func Mylog() *mylog {
	mylog_once.Do(func() {
		my_log = &mylog{}
		if err := my_log.LoadConfigure(); err != nil {
			log.Fatalln(err)
		}
	})
	return my_log
}

// LoadConfigure 从file里读取seelog 配置
func (l *mylog) LoadConfigure() error {
	conf := GetConfigure()
	file := conf.Log.File
	logger, err := seelog.LoggerFromConfigAsFile(file)
	if err != nil {
		return err
	}
	l.Log = logger

	return nil
}

// Infof 输出info信息
func (l *mylog) Infof(format string, v ...interface{}) {
	l.Log.Infof(format, v...)
}

// Debugf 输出debug信息
func (l *mylog) Debugf(format string, v ...interface{}) {
	l.Log.Debugf(format, v...)
}

// Warnf 输出warn信息
func (l *mylog) Warnf(format string, v ...interface{}) {
	l.Log.Warnf(format, v...)
}

// Errorf 输出error信息
func (l *mylog) Errorf(format string, v ...interface{}) {
	l.Log.Errorf(format, v...)
}

func (l *mylog) Info(v ...interface{}) {
	l.Log.Info(v...)
}

func (l *mylog) Debug(v ...interface{}) {
	l.Log.Debug(v...)
}

func (l *mylog) Warn(v ...interface{}) {
	l.Log.Warn(v...)
}

func (l *mylog) Error(v ...interface{}) {
	l.Log.Error(v...)
}

// Infof 输出info信息
func (l *mylog) Stop() {
	l.Log.Flush()
}

func GetMylog() *mylog {
	return my_log
}
