package log

import (
	"github.com/phachon/go-logger"
	"os"
)

var Logger = go_logger.NewLogger()

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func init() {
	_dir := "./logs"
	exist, err := PathExists(_dir)
	if err != nil {
		println("get dir error![%v]\n", err)
		return
	}
	if exist {
		println("has dir![%v]\n", _dir)
	} else {
		println("no dir![%v]\n", _dir)
		err := os.Mkdir(_dir, os.ModePerm)
		if err != nil {
			println("mkdir failed![%v]\n", err)
		} else {
			println("mkdir success!\n")
		}
	}

	Logger.Detach("console")

	// 命令行输出配置
	consoleConfig := &go_logger.ConsoleConfig{
		Color:      true,                                                                                                                 // 命令行输出字符串是否显示颜色
		JsonFormat: true,                                                                                                                 // 命令行输出字符串是否格式化
		Format:     "[Time:%millisecond_format%] [Level:%level_string%] [Body:%body%] [File:%file%] [Line:%line%] [Function:%function%]", // 如果输出的不是 json 字符串，JsonFormat: false, 自定义输出的格式
	}
	// 添加 console 为 logger 的一个输出
	Logger.Attach("console", go_logger.LOGGER_LEVEL_DEBUG, consoleConfig)

	// 文件输出配置
	fileConfig := &go_logger.FileConfig{
		Filename: "./logs/my.log", // 日志输出文件名，不自动存在
		// 如果要将单独的日志分离为文件，请配置LealFrimeNem参数。
		LevelFileName: map[int]string{
			Logger.LoggerLevel("error"): "./logs/error.log", // Error 级别日志被写入 error .log 文件
			Logger.LoggerLevel("info"):  "./logs/info.log",  // Info 级别日志被写入到 info.log 文件中
			Logger.LoggerLevel("debug"): "./logs/debug.log", // Debug 级别日志被写入到 debug.log 文件中
		},
		MaxSize:    1024 * 1024,                                                                                                          // 文件最大值（KB），默认值0不限
		MaxLine:    100000,                                                                                                               // 文件最大行数，默认 0 不限制
		DateSlice:  "d",                                                                                                                  // 文件根据日期切分， 支持 "Y" (年), "m" (月), "d" (日), "H" (时), 默认 "no"， 不切分
		JsonFormat: false,                                                                                                                // 写入文件的数据是否 json 格式化
		Format:     "[Time:%millisecond_format%] [Level:%level_string%] [Body:%body%] [File:%file%] [Line:%line%] [Function:%function%]", // 如果写入文件的数据不 json 格式化，自定义日志格式
	}

	// 添加 file 为 logger 的一个输出
	Logger.Attach("file", go_logger.LOGGER_LEVEL_DEBUG, fileConfig)
}
