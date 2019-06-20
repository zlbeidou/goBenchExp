package goBenchExp

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func BenchmarkInlineFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		inlineFunc(1, 2)
	}
}

func BenchmarkInlineFuncParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			inlineFunc(1, 2)
		}
	})
}

func BenchmarkEmptyFuncVs0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if i == math.MaxInt32 {
			fmt.Println(0)
		}
	}
}

func BenchmarkEmptyFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		emptyFunc(0)
	}
}

func BenchmarkSimpleDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simpleDefer(0)
	}
}

func BenchmarkGoEmptyFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go emptyFunc(0)
	}
}

var testStr = `Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.`

func BenchmarkStrToByteSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strToByteSlice(testStr)
	}
}

func BenchmarkByteSliceToStr(b *testing.B) {
	testB := []byte(testStr)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		byteSliceToStr(testB)
	}
}

func BenchmarkByteSliceToStrDirect(b *testing.B) {
	testB := []byte(testStr)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		byteSliceToStrDirect(testB)
	}
}

func BenchmarkTypeAssert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		typeAssert()
	}
}

func BenchmarkSimpleReflectTypeOf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simpleReflectTypeOf()
	}
}

func BenchmarkSimpleReflectValueOf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simpleReflectValueOf()
	}
}

func BenchmarkGetFromCh(b *testing.B) {
	ch := make(chan string)
	go func() {
		for {
			ch <- "hello world"
		}
	}()
	time.Sleep(time.Millisecond)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getFromCh(ch)
	}
}

func BenchmarkGetFromChGoroutine2(b *testing.B) {
	ch := make(chan string)
	go func() {
		for {
			ch <- "hello world"
		}
	}()
	go func() {
		for {
			ch <- "hello world"
		}
	}()
	time.Sleep(time.Millisecond)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getFromCh(ch)
	}
}

func BenchmarkGetFromChGo2Parallel(b *testing.B) {
	ch := make(chan string)
	go func() {
		for {
			ch <- "hello world"
		}
	}()
	go func() {
		for {
			ch <- "hello world"
		}
	}()
	time.Sleep(time.Millisecond)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			getFromCh(ch)
		}
	})
}

func BenchmarkGetFromBufferCh(b *testing.B) {
	ch := make(chan string, 1000)
	go func() {
		for {
			ch <- "hello world"
		}
	}()
	time.Sleep(time.Millisecond)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getFromCh(ch)
	}
}

func BenchmarkGetFromBufferChCap2(b *testing.B) {
	ch := make(chan string, 10000)
	go func() {
		for {
			ch <- "hello world"
		}
	}()
	time.Sleep(time.Millisecond)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getFromCh(ch)
	}
}

func BenchmarkGetFromBufChGo2Parallel(b *testing.B) {
	ch := make(chan string, 1000)
	go func() {
		for {
			ch <- "hello world"
		}
	}()
	go func() {
		for {
			ch <- "hello world"
		}
	}()
	time.Sleep(time.Millisecond)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			getFromCh(ch)
		}
	})
}

var jsonStr = `{"apple":"appleItem","boy":1,"cat":"catItem","dog":{"aItem":[1,2,3],"bItem":[4,5,6]}}`

type jsonSt struct {
	Apple string           `json:"apple"`
	Boy   int              `json:"boy"`
	Cat   string           `json:"cat"`
	Dog   map[string][]int `json:"dog"`
}

func BenchmarkJsonStd(b *testing.B) {
	jsonB := []byte(jsonStr)
	jsonV := &jsonSt{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		jsonUnmarshalStd(jsonB, jsonV)
	}
}

func BenchmarkJsonStdParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		jsonB := []byte(jsonStr)
		jsonV := &jsonSt{}

		for pb.Next() {
			jsonUnmarshalStd(jsonB, jsonV)
		}
	})
}

func BenchmarkJson3rd(b *testing.B) {
	jsonB := []byte(jsonStr)
	jsonV := &jsonSt{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		jsonUnmarshal3rd(jsonB, jsonV)
	}
}

func BenchmarkJson3rdParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		jsonB := []byte(jsonStr)
		jsonV := &jsonSt{}

		for pb.Next() {
			jsonUnmarshal3rd(jsonB, jsonV)
		}
	})
}

var cnt int64

func BenchmarkAtomicAdd(b *testing.B) {
	var cnt int64
	for i := 0; i < b.N; i++ {
		atomic.AddInt64(&cnt, 1)
	}
}

func BenchmarkAtomicAddParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.AddInt64(&cnt, 1)
		}
	})
}

func BenchmarkAtomicCas(b *testing.B) {
	var cnt int64
	for i := 0; i < b.N; i++ {
		oldValue := atomic.LoadInt64(&cnt)
		for !atomic.CompareAndSwapInt64(&cnt, oldValue, int64(oldValue+1)) {
			oldValue = atomic.LoadInt64(&cnt)
		}
	}
}

func BenchmarkAtomicCasParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			oldValue := atomic.LoadInt64(&cnt)
			for !atomic.CompareAndSwapInt64(&cnt, oldValue, int64(oldValue+1)) {
				oldValue = atomic.LoadInt64(&cnt)
			}
		}
	})
}

func BenchmarkMutex(b *testing.B) {
	var mut sync.Mutex
	for i := 0; i < b.N; i++ {
		mut.Lock()
		cnt = cnt + 1
		mut.Unlock()
	}
}

func BenchmarkMutexParallel(b *testing.B) {
	var mut sync.Mutex
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mut.Lock()
			cnt = cnt + 1
			mut.Unlock()
		}
	})
}
