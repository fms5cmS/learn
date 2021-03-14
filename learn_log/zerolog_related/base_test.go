package zerolog_test

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"testing"
)

func TestZerolog_Base(t *testing.T) {
	// UNIX Time is faster and smaller than most timestamps
	// If you set zerolog.TimeFieldFormat to an empty string,
	// logs will write with UNIX time
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// 默认情况下，日志写入标准错误（os_rela.Stderr），而 log.Print 默认的错误级别是 debug，print 建议只是临时调试用。
	log.Print("proto world")
}

// 日志应该由字段和消息组成，而字段就是记录日志的上下文
// 注意，所有的链式 API，最后都需要调用 Msg 或 Msgf 或 Send()，否则，日志不会输出。
// Send() 等价于 Msg("")
func TestZerolog_Context(t *testing.T) {
	zerolog.TimeFieldFormat = zerolog.TimeFieldFormat
	log.Debug().
		Str("Scale","833 cents").
		Float64("Interval",833.09).
		Msg("Fibonacci is everywhere")
	log.Debug().
		Str("Name", "Tom").
		Send()
}

// zerolog 由高到低支持如下等级：
// panic (zerolog.PanicLevel, 5)
// fatal (zerolog.FatalLevel, 4)
// error (zerolog.ErrorLevel, 3)
// warn (zerolog.WarnLevel, 2)
// info (zerolog.InfoLevel, 1)
// debug (zerolog.DebugLevel, 0)
// trace (zerolog.TraceLevel, -1)
func TestZerolog_level(t *testing.T) {
	// 设置全局的日志等级，无论选择哪个级别，都会写入所有大于或等于该级别的日志
	// 如果要完全关闭日志记录，传递 zerolog.Disabled 常量即可
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Debug().Str("debug-k","debug-v").Msg("debug level information")
	log.Warn().Str("warn-k","warn-v").Msg("warn level information")
}

func TestZerolog_logfile(t *testing.T) {
	logfile, _ := os.Create("log.txt")
	// github.com/rs/zerolog/log 中的全局 logger 默认是输出到 os_rela.Stderr，修改输出位置即可将日志写入文件中
	// 方法一：
	log.Logger = zerolog.New(logfile).With().Timestamp().Logger()
	// 方法二：
	log.Logger = log.Output(logfile)
	log.Print("this message will be writed in a file")
}

// 定制自己的上下文信息
// 希望日志都带上固定的一些信息，比如 traceId，可以通过子日志的形式实现
func TestZerolog_customizeContext(t *testing.T) {
	subLogger := log.With().Str("traceId", "132").Logger()
	subLogger.Info().Msg("go")
}

func TestZerolog_TTYwithColor(t *testing.T) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Str("foo","bar").Msg("colorful")
}