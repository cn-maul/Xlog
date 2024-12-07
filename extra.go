package Xlog

import "strings"

// isLevelAllowed 检查日志级别是否允许
func (l *Logger) isLevelAllowed(level string) bool {
	for _, allowedLevel := range l.Config.logLevels {
		if strings.EqualFold(allowedLevel, level) {
			return true
		}
	}
	return false
}

// shouldOutputToStdOut 检查是否应该输出到标准输出
func (l *Logger) shouldOutputToStdOut() bool {
	for _, output := range l.Config.logOutput {
		if strings.EqualFold(output, "stdOut") {
			return true
		}
	}
	return false
}

// Info 记录信息级别的日志
func (l *Logger) Info(content map[string]string) error {
	return l.log("info", content)
}

// Warn 记录警告级别的日志
func (l *Logger) Warn(content map[string]string) error {
	return l.log("warn", content)
}

// Error 记录错误级别的日志
func (l *Logger) Error(content map[string]string) error {
	return l.log("error", content)
}

// Fatal 记录致命错误级别的日志
func (l *Logger) Fatal(content map[string]string) error {
	return l.log("fatal", content)
}
