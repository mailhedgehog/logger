package logger

import "fmt"

// LogLevel represents an log level name
type LogLevel = string

// MinLogLevel minimal level to display
var MinLogLevel = INFO

var (
    LogHandler          func(message string, context ...interface{})
    LogEmergencyHandler func(message string, context ...interface{})
    LogAlertHandler     func(message string, context ...interface{})
    LogCriticalHandler  func(message string, context ...interface{})
    LogErrorHandler     func(message string, context ...interface{})
    LogWarningHandler   func(message string, context ...interface{})
    LogNoticeHandler    func(message string, context ...interface{})
    LogInfoHandler      func(message string, context ...interface{})
    LogDebugHandler     func(message string, context ...interface{})
)

const (
    EMERGENCY LogLevel = "emergency"
    ALERT     LogLevel = "alert"
    CRITICAL  LogLevel = "critical"
    ERROR     LogLevel = "error"
    WARNING   LogLevel = "warning"
    NOTICE    LogLevel = "notice"
    INFO      LogLevel = "info"
    DEBUG     LogLevel = "debug"
)

var (
    EmergencyColor = Red
    AlertColor     = Red
    CriticalColor  = Red
    ErrorColor     = Red
    WarningColor   = Yellow
    InfoColor      = Teal
    NoticeColor    = Purple
    DebugColor     = Green
)

var (
    Red    = Color("\033[1;31m%s\033[0m")
    Green  = Color("\033[1;32m%s\033[0m")
    Yellow = Color("\033[1;33m%s\033[0m")
    Purple = Color("\033[1;34m%s\033[0m")
    Teal   = Color("\033[1;36m%s\033[0m")
)

func PanicIfError(e error) {
    if e != nil {
        Critical(e.Error())

        panic(e)
    }
}

func Color(colorString string) func(...interface{}) string {
    sprint := func(args ...interface{}) string {
        return fmt.Sprintf(colorString,
            fmt.Sprint(args...))
    }
    return sprint
}

func logLevelNumber(logLevel LogLevel) int {
    switch logLevel {
    case EMERGENCY:
        return 7
    case ALERT:
        return 6
    case CRITICAL:
        return 5
    case ERROR:
        return 4
    case WARNING:
        return 3
    case NOTICE:
        return 2
    case INFO:
        return 1
    case DEBUG:
        return 0
    default:
        panic(fmt.Sprintf("Wrong log level: %s", logLevel))
    }
}

func Log(level LogLevel, message string, context ...interface{}) {
    switch level {
    case EMERGENCY:
        Emergency(message, context...)
    case ALERT:
        Alert(message, context...)
    case CRITICAL:
        Critical(message, context...)
    case ERROR:
        Error(message, context...)
    case WARNING:
        Warning(message, context...)
    case NOTICE:
        Notice(message, context...)
    case INFO:
        Info(message, context...)
    case DEBUG:
        Debug(message, context...)
    default:
        panic(fmt.Sprintf("Wrong log level: %s", level))
    }
}

func Emergency(message string, context ...interface{}) {
    if logLevelNumber(EMERGENCY) < logLevelNumber(MinLogLevel) {
        return
    }

    if LogEmergencyHandler != nil {
        LogEmergencyHandler(message, context...)
        return
    }

    if LogHandler != nil {
        LogHandler(message, context...)
        return
    }

    if len(context) > 0 {
        fmt.Println(EmergencyColor(message), context)
        return
    }
    fmt.Println(EmergencyColor(message))
}

func Alert(message string, context ...interface{}) {
    if logLevelNumber(ALERT) < logLevelNumber(MinLogLevel) {
        return
    }

    if LogAlertHandler != nil {
        LogAlertHandler(message, context...)
        return
    }

    if LogHandler != nil {
        LogHandler(message, context...)
        return
    }

    if len(context) > 0 {
        fmt.Println(AlertColor(message), context)
        return
    }
    fmt.Println(AlertColor(message))
}

func Critical(message string, context ...interface{}) {
    if logLevelNumber(CRITICAL) < logLevelNumber(MinLogLevel) {
        return
    }

    if LogCriticalHandler != nil {
        LogCriticalHandler(message, context...)
        return
    }

    if LogHandler != nil {
        LogHandler(message, context...)
        return
    }

    if len(context) > 0 {
        fmt.Println(CriticalColor(message), context)
        return
    }
    fmt.Println(CriticalColor(message))
}

func Error(message string, context ...interface{}) {
    if logLevelNumber(ERROR) < logLevelNumber(MinLogLevel) {
        return
    }

    if LogErrorHandler != nil {
        LogErrorHandler(message, context...)
        return
    }

    if LogHandler != nil {
        LogHandler(message, context...)
        return
    }

    if len(context) > 0 {
        fmt.Println(ErrorColor(message), context)
        return
    }
    fmt.Println(ErrorColor(message))
}

func Warning(message string, context ...interface{}) {
    if logLevelNumber(WARNING) < logLevelNumber(MinLogLevel) {
        return
    }

    if LogWarningHandler != nil {
        LogWarningHandler(message, context)
        return
    }

    if LogHandler != nil {
        LogHandler(message, context...)
        return
    }

    if len(context) > 0 {
        fmt.Println(WarningColor(message), context)
        return
    }
    fmt.Println(WarningColor(message))
}

func Notice(message string, context ...interface{}) {
    if logLevelNumber(NOTICE) < logLevelNumber(MinLogLevel) {
        return
    }

    if LogNoticeHandler != nil {
        LogNoticeHandler(message, context...)
        return
    }

    if LogHandler != nil {
        LogHandler(message, context...)
        return
    }

    if len(context) > 0 {
        fmt.Println(NoticeColor(message), context)
        return
    }
    fmt.Println(NoticeColor(message))
}

func Info(message string, context ...interface{}) {
    if logLevelNumber(INFO) < logLevelNumber(MinLogLevel) {
        return
    }

    if LogInfoHandler != nil {
        LogInfoHandler(message, context...)
        return
    }

    if LogHandler != nil {
        LogHandler(message, context...)
        return
    }

    if len(context) > 0 {
        fmt.Println(InfoColor(message), context)
        return
    }
    fmt.Println(InfoColor(message))
}

func Debug(message string, context ...interface{}) {
    if logLevelNumber(DEBUG) < logLevelNumber(MinLogLevel) {
        return
    }

    if LogDebugHandler != nil {
        LogDebugHandler(message, context...)
        return
    }

    if LogHandler != nil {
        LogHandler(message, context...)
        return
    }

    if len(context) > 0 {
        fmt.Println(DebugColor(message), context)
        return
    }
    fmt.Println(DebugColor(message))
}
