package Xlog

import (
	"errors"
	"fmt"
	"time"
)

// LogConfig 日志服务的配置结构体
type LogConfig struct {
	logDir      string   //日志文件存放在本地的路径，比如./log/
	logFileName string   //日志文件的命名格式，比如2024-12-07 19:06:54.log
	logMaxNum   int      //在程序中写入本地文件前存储最大的日志数量，比如100
	logFormat   string   //输出的日志格式，比如json、xml、text
	logLevels   []string //允许的日志登记，例如[]string{"info","error","fatal"}
	logOutput   []string //输出日志到什么地方，比如stdOut
}

// LogItem 单条日志的结构体
type LogItem struct {
	timeStamp time.Time
	level     string
	Content   map[string]string
}

// Logger 向用户返回的日志服务结构体
type Logger struct {
	Config LogConfig
	List   []LogItem
}

// New 创建一个新的Logger实例，默认值
func New() (*Logger, error) {
	config := LogConfig{
		logDir:      "./log/",
		logFileName: "default.log",
		logMaxNum:   100,
		logFormat:   "text",
		logLevels:   []string{"Info", "Warn", "Error", "Fatal"},
		logOutput:   []string{"stdOut"},
	}

	return &Logger{
		Config: config,
		List:   make([]LogItem, 0),
	}, nil
}

// NewWithConfig 创建一个新的Logger实例，传入自定义配置
func NewWithConfig(config LogConfig) (*Logger, error) {
	if config.logDir == "" {
		return nil, errors.New("日志目录不能为空")
	}
	if config.logFileName == "" {
		return nil, errors.New("日志文件名不能为空")
	}
	if config.logMaxNum <= 0 {
		return nil, errors.New("日志最大条数不能小于1")
	}
	if config.logFormat == "" {
		return nil, errors.New("日志输出格式不能为空")
	}
	if len(config.logLevels) == 0 {
		return nil, errors.New("日志级别不能为空")
	}
	if len(config.logOutput) == 0 {
		return nil, errors.New("日志输出不能为空")
	}

	return &Logger{
		Config: config,
		List:   make([]LogItem, 0),
	}, nil
}

// Log 打印日志到终端
func (l *Logger) log(level string, content map[string]string) error {
	// 检查日志级别是否允许
	if !l.isLevelAllowed(level) {
		return fmt.Errorf("日志级别%s不在可允许的范围内", level)
	}

	// 创建新的日志项
	logItem := LogItem{
		timeStamp: time.Now(),
		level:     "info",
		Content:   content,
	}

	// 添加到日志列表
	l.List = append(l.List, logItem)

	// 打印到终端
	if l.shouldOutputToStdOut() {
		fmt.Printf("%s [%s] ", logItem.timeStamp.Format("2006-01-02 15:04:05"), logItem.level)
		for key, value := range logItem.Content {
			fmt.Printf("%s:%s ", key, value)
		}
		fmt.Println()
	}

	return nil
}
