package main

import (
	"sync"
	"unsafe"
)

func main() {

}

type (
	index    uint8
	store    map[index]interface{}
	callback struct {
		callback interface{}
		arg      unsafe.Pointer
	}
)

var (
	storage = make(store)
	mutex   sync.Mutex
)

//export evenNumberCallbackProxy
func evenNumberCallbackProxy(i uint, num int) {
	c := getCallback(index(i))
	c.callback.(func(int))(num)
}

//export userCallbackProxy
func userCallbackProxy(i uint) {
	c := getCallback(index(i))
	c.callback.(func(unsafe.Pointer))(c.arg)
}

func registerCallback(c interface{}, arg unsafe.Pointer) uint {
	mutex.Lock()
	i := index(len(storage))
	storage[i] = &callback{
		callback: c,
		arg:      arg,
	}
	mutex.Unlock()

	return uint(i)
}

func getCallback(i index) *callback {
	mutex.Lock()
	defer mutex.Unlock()
	return storage[i].(*callback)
}

func unregisterCallback(i uint) {
	mutex.Lock()
	delete(storage, index(i))
	mutex.Unlock()
}
