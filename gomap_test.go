package gomap_test

import (
	"fmt"
	"sync"

	"github.com/ybs-github/gomap"

	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGoMap(t *testing.T) {
	Convey("GoMap", t, func() {
		m := gomap.NewGoMap()
		So(m.All(), ShouldResemble, map[string]interface{}{})

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

		So(m.All(), ShouldResemble, map[string]interface{}{
			"key0": "value0",
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
			"key4": "value4",
			"key5": "value5",
			"key6": "value6",
			"key7": "value7",
			"key8": "value8",
			"key9": "value9"})

		for i := 0; i < 10; i++ {
			So(m.Get(fmt.Sprintf("key%d", i)), ShouldEqual, fmt.Sprintf("value%d", i))
		}

		So(m.Get("foo"), ShouldBeNil)

		wg.Add(5)
		for i := 0; i < 5; i++ {
			gi := i
			go func() {
				defer wg.Done()
				m.Remove(fmt.Sprintf("key%d", gi))
			}()
		}
		wg.Wait()

		So(m.All(), ShouldResemble, map[string]interface{}{
			"key5": "value5",
			"key6": "value6",
			"key7": "value7",
			"key8": "value8",
			"key9": "value9"})

		m.Clear()
		So(m.All(), ShouldResemble, map[string]interface{}{})
	})
}
