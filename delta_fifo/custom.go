package delta_fifo

import (
	"fmt"
	"k8s.io/client-go/tools/cache"
)

const (
	Added   cache.DeltaType = "Added"
	Updated cache.DeltaType = "Updated"
	Deleted cache.DeltaType = "Deleted"
	// The other types are obvious. You'll get Sync deltas when:
	//  * A watch expires/errors out and a new list/watch cycle is started.
	//  * You've turned on periodic syncs.
	// (Anything that trigger's DeltaFIFO's Replace() method.)
	Sync cache.DeltaType = "Sync"
)

func customPopFunc(obj interface{}) error {
	deltas := obj.(cache.Deltas)
	if deltas[len(deltas)-1].Type != Added {
		fmt.Println("Added obj : " + deltas[len(deltas)-1].Object.(testFifoObject).name)
		return nil
	}
	return cache.ErrRequeue{Err: nil}
}

func testFifoObjectKeyFunc(obj interface{}) (string, error) {
	return obj.(testFifoObject).name, nil
}

func mkFifoObj(name string, val interface{}) testFifoObject {
	return testFifoObject{name: name, val: val}
}

type testFifoObject struct {
	name string
	val  interface{}
}

type customObjects []testFifoObject

func (kl customObjects) ListKeys() []string {
	result := []string{}
	for _, fifoObj := range kl {
		result = append(result, fifoObj.name)
	}
	return result
}

func (kl customObjects) GetByKey(key string) (interface{}, bool, error) {
	for _, v := range kl {
		if v.name == key {
			return v, true, nil
		}
	}
	return nil, false, nil
}
