package benchmarksortedmaps

import (
	"cmp"
	"fmt"
	"testing"

	omap "github.com/akalinux/orderedmap"
	sm "github.com/egregors/sortedmap"
	btree "github.com/google/btree"
)

type KV[V any] struct {
	key   string
	Value V
}

func (a KV[any]) Less(than btree.Item) bool {
	val, _ := than.(KV[any])
	return a.key < val.key
}

func BenchmarkAll(b *testing.B) {
	format := "%04d"
	for _, size := range []int{1000, 2000, 5000} {
		b.Logf("Working with set size: %d", size)
		var m map[string]any = make(map[string]any, size)
		// Write tests
		b.Log("** Write tests")
		b.Run(fmt.Sprintf("Native Map Put: %d", size), func(b *testing.B) {
			for k := range size {
				key := fmt.Sprintf(format, k)
				m[key] = nil
			}
		})

		var ct *omap.CenterTree[string, any]
		b.Run(fmt.Sprintf("CenterTree Put: %d", size), func(b *testing.B) {
			ct = omap.NewCenterTree[string, any](size>>1, cmp.Compare)
			for k := range size {
				key := fmt.Sprintf(format, k)
				ct.Put(key, nil)
			}
		})
		var so *sm.SortedMap[map[string]any, string, any]
		b.Run("sortedmap Put", func(b *testing.B) {
			so = sm.New[map[string]any](func(i, j sm.KV[string, any]) bool { return cmp.Less(i.Key, j.Key) })
			for k := range size {
				key := fmt.Sprintf(format, k)
				so.Insert(key, nil)
			}
		})

		var bt *btree.BTree
		b.Run("btree Put", func(b *testing.B) {
			bt = btree.New(2)
			for k := range size {
				key := fmt.Sprintf(format, k)
				bt.ReplaceOrInsert(KV[any]{key: key, Value: nil})
			}
		})

		// Read tests
		b.Log("** Read tests")
		b.Run(fmt.Sprintf("Native Map Get: %d", size), func(b *testing.B) {
			for k := range size {
				key := fmt.Sprintf(format, k)
				_, _ = m[key]
			}
		})
		b.Run(fmt.Sprintf("CenterTree Get: %d", size), func(b *testing.B) {
			for k := range size {
				key := fmt.Sprintf(format, k)
				ct.Get(key)
			}
		})
		b.Run(fmt.Sprintf("sortedmap Get: %d", size), func(b *testing.B) {
			for k := range size {
				key := fmt.Sprintf(format, k)
				so.Get(key)
			}
		})
		b.Run(fmt.Sprintf("btree Get: %d", size), func(b *testing.B) {
			for k := range size {
				key := fmt.Sprintf(format, k)
				bt.Get(KV[any]{key: key, Value: nil})
			}
		})

		b.Log("Count all values from mid to end")

		mid := size / 2
		start := fmt.Sprintf(format, mid)
		end := fmt.Sprintf(format, size-1)
		b.Run(fmt.Sprintf("Native Count: %d", size), func(b *testing.B) {
			count := 0
			for k := range m {
				if k >= start {
					count++
				}
			}
		})
		b.Run(fmt.Sprintf("CenterTree Count: %d", size), func(b *testing.B) {
			ct.Between(start, end, omap.LAST_KEY)
		})
		b.Run(fmt.Sprintf("sortedmap Count: %d", size), func(b *testing.B) {
			count := 0
			for k := range so.All() {
				if k >= start {
					count++
				}
			}

		})
		b.Run(fmt.Sprintf("btree Count: %d", size), func(b *testing.B) {
			count := 0
			start := fmt.Sprintf(format, mid-1)
			bt.DescendGreaterThan(KV[any]{key: start}, func(item btree.Item) bool {
				count++
				return true
			})
		})

	}
}
