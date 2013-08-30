package user_error

import (
	"log"
)

type callback interface{}

const (
	NORMAL  = 0
	FATAL   = 1
	ERROR   = 2
	WARNING = 3
	NOTICE  = 4
	UNKNOWN = 5
)

func Check_error(err error, level int) {
	if err != nil {
		switch level {
		case FATAL:
			log.Println("Event Class:FATAL!", err)
		case ERROR:
			log.Println("Event Class:ERROR!", err)
		case WARNING:
			log.Println("Event Class:WAUNING!", err)
		case NOTICE:
			log.Println("Event Class:NOTICE!", err)
		default:
			log.Println("Event Class:UNKNOWN!", err)
		}
	}
}

func WipeAss(err error, level int, callback func(...interface{}), Parameters4CallBack ...interface{}) {
	Check_error(err, level)
	callback(Parameters4CallBack)
}
