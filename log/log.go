package log

import (
	"github.com/mageddo/go-logging"
	"os"
)

var LOGGER logging.Log

func init(){

	mode := "dev"
	backend := logging.NewLogBackend(os.Stdout, "", 0)

	// setando o log dependendo do ambiente
	switch mode {
	case "prod":
		format := logging.MustStringFormatter(
			`%{time:06-01-02 15:04:05} %{level:.3s} %{message}`,
		)
		leveledBackend := logging.AddModuleLevel(logging.NewBackendFormatter(backend, format));
		leveledBackend.SetLevel(logging.INFO, "")
		logging.SetBackend(leveledBackend)
		break;
	default:
		format := logging.MustStringFormatter(
			`%{color}%{time:06-01-02 15:04:05.000} %{level:.3s} %{color:reset}%{message}`,
		)
		backend2Formatter := logging.NewBackendFormatter(backend, format)
		logging.SetBackend(backend2Formatter)
		break;
	}
	LOGGER = logging.NewLog(logging.NewContext())
}
