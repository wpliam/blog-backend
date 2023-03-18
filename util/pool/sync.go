package pool

import (
	"fmt"
	"sync"
)

func GoAndWait(handlers ...func() error) error {
	var (
		once sync.Once
		wg   sync.WaitGroup
		err  error
	)
	for _, handler := range handlers {
		wg.Add(1)
		go func(f func() error) {
			defer func() {
				wg.Done()
				if e := recover(); e != nil {
					once.Do(func() {
						err = fmt.Errorf("[PANIC] e:%v", e)
					})
				}
			}()
			if e := f(); e != nil {
				once.Do(func() {
					err = e
				})
			}
		}(handler)
	}
	wg.Wait()
	return err
}
