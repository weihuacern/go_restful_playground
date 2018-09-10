package pam

import "sync"

var CallBack struct {
	sync.Mutex
	m map[int]interface{}
	c int
}

func init() {
	CallBack.m = make(map[int]interface{})
}

func CallBackAdd(v interface{}) int {
	CallBack.Lock()
	defer CallBack.Unlock()
	CallBack.c++
	CallBack.m[CallBack.c] = v
	return CallBack.c
}

func CallBackGet(c int) interface{} {
	CallBack.Lock()
	defer CallBack.Unlock()
	if v, ok := CallBack.m[c]; ok {
		return v
	}
	panic("Callback pointer not found")
}

func CallBackDelete(c int) {
	CallBack.Lock()
	defer CallBack.Unlock()
	if _, ok := CallBack.m[c]; !ok {
		panic("Callback pointer not found")
	}
	delete(CallBack.m, c)
}
