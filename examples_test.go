package gomap_test

import (
	"fmt"
	"sync"

	"github.com/ybs-github/gomap"
)

func Example() {
	m := gomap.NewGoMap()

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		gi := i
		go func() {
			defer wg.Done()
			m.Set(fmt.Sprintf("key%d", gi), fmt.Sprintf("value%d", gi))
		}()
	}
	wg.Wait()

	for i := 0; i < 10; i++ {
		fmt.Println(m.Get(fmt.Sprintf("key%d", i)))
	}

	// Output: value0
	//value1
	//value2
	//value3
	//value4
	//value5
	//value6
	//value7
	//value8
	//value9
}
