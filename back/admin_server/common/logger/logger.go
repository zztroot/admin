package logger

import "github.com/zztroot/zztlog"

func InitLogger(){
	if err := zztlog.InitConfigFile("./common/logger/zztlog.json"); err != nil {
		zztlog.Error(err)
		return
	}
	// 测试logger是否初始化成功
	zztlog.Debug("logger init success")
}
