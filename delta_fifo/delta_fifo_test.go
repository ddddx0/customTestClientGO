package delta_fifo

import (
	"k8s.io/client-go/tools/cache"
	"testing"
)

func TestDeltaFIFO_addUpdateDelete(t *testing.T) {
	f := cache.NewDeltaFIFO(testFifoObjectKeyFunc, nil)
	f.Add(mkFifoObj("foo", 10))
	f.Update(mkFifoObj("foo", 12))
	f.Delete(mkFifoObj("foo", 15))
}

func TestDeltaFIFO_replace(t *testing.T) {
	f := cache.NewDeltaFIFO(
		testFifoObjectKeyFunc,
		customObjects{mkFifoObj("foo", 5), mkFifoObj("bar", 6), mkFifoObj("baz", 7)},
	)
	f.Delete(mkFifoObj("baz", 10))
	f.Replace([]interface{}{mkFifoObj("foo", 5)}, "0")
}

func TestDeltaFIFO_resync(t *testing.T) {
	f := cache.NewDeltaFIFO(
		testFifoObjectKeyFunc,
		customObjects{mkFifoObj("foo", 5), mkFifoObj("bar", 6), mkFifoObj("baz", 7)},
	)
	f.Delete(mkFifoObj("foo", 10))
	f.Resync()
}

func TestDeltaFIFO_pop(t *testing.T) {
	f := cache.NewDeltaFIFO(testFifoObjectKeyFunc, nil)

	f.Add(mkFifoObj("foo", 10))
	_, err := f.Pop(customPopFunc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
