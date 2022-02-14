package singleton

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type singleton struct {
	val string
}

var singleTonObj *singleton


func (*singleton) DoSomething() {
	fmt.Println("do something")
}

var once sync.Once

func GetSingletonObj() *singleton {
	once.Do(func() {
		singleTonObj = &singleton{}
	})
	return singleTonObj
}

var mutex sync.Mutex
var done uint32

// same as GetSingletonObj
func GetSingletonObjAtomic() *singleton {
	if atomic.LoadUint32(&done) == 0 {
		mutex.Lock()
		defer mutex.Unlock()

		if done == 0 {
			singleTonObj = &singleton{}
			atomic.StoreUint32(&done, 1)
		}
	}
	return singleTonObj
}
