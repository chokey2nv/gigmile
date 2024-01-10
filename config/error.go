package config

import (
	"fmt"
	"log"
	"runtime"
)

func Errorf(any interface{}, a ...interface{}) error {
	if any != nil {
		err := (error)(nil)

		switch any.(type) {
		case string:
			err = fmt.Errorf(any.(string), a...)
		case error:
			err = fmt.Errorf(any.(error).Error(), a...)
		default:
			err = fmt.Errorf("%+v", any)
		}
		_, fn, line, _ := runtime.Caller(1)
		log.Printf("ERROR: [%s:%d] %v \n", fn, line, err)
		return err
	}
	return nil
}
