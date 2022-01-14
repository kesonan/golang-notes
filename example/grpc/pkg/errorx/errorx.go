package errorx

import "log"

func MustFatalln(err error) {
	if err == nil {
		return
	}
	log.Fatalln(err)
}
