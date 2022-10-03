package recoverer

import (
	"errors"
	"log"
	"time"
)

type Recoverer struct {
	PrintError func(string)
	SleepTime  uint64
}

func (r Recoverer) RecoverErroredFunction(f func() error) {
	var err error
	for err = f(); err != nil; err = f() {
		r.PrintError(err.Error() + "\n#error")
		time.Sleep(time.Second * time.Duration(r.SleepTime))
	}
}

func (r Recoverer) RecoverPanicFunction(f func()) {
	r.RecoverErroredFunction(
		func() (err error) {
			defer func() {
				if message := recover(); message != nil {
					log.Printf("type %T\n	", message)
					switch v := message.(type) {
					default:
						log.Printf("unexpected type %T", v)
					case error:
						log.Printf("type %T\n	", v)
						err = v
					case string:
						log.Printf("type %T\n	", v)
						err = errors.New(v)
					}
				}
			}()
			f()
			return nil
		})
}
