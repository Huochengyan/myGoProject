package log

//log emergency level
func Emergency(msg string) {
	Logger.Emergency(msg)
}

//log emergency format
func Emergencyf(format string, a ...interface{}) {
	Logger.Emergencyf(format, a)
}

//log alert level
func Alert(msg string) {
	Logger.Alert(msg)
}

//log alert format
func Alertf(format string, a ...interface{}) {
	Logger.Alertf(format, a)
}

//log critical level
func Critical(msg string) {
	Logger.Critical(msg)
}

//log critical format
func Criticalf(format string, a ...interface{}) {
	Logger.Criticalf(format, a)
}

//log error level
func Error(msg string) {
	Logger.Error(msg)
}

//log error format
func Errorf(format string, a ...interface{}) {
	Logger.Errorf(format, a)
}

//log warning level
func Warning(msg string) {
	Logger.Warning(msg)
}

//log warning format
func Warningf(format string, a ...interface{}) {
	Logger.Warningf(format, a)
}

//log notice level
func Notice(msg string) {
	Logger.Notice(msg)
}

//log notice format
func Noticef(format string, a ...interface{}) {
	Logger.Noticef(format, a)
}

//log info level
func Info(msg string) {
	Logger.Info(msg)
}

//log info format
func Infof(format string, a ...interface{}) {
	Logger.Infof(format, a)
}

//log debug level
func Debug(msg string) {
	Logger.Debug(msg)
}

//log debug format
func Debugf(format string, a ...interface{}) {
	Logger.Debugf(format, a)
}
