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
	for _, size := range []int{50} {
		b.Logf("Working with set size: %d", size)
		var m map[string]any
		// Write tests
		b.Log("** Write tests")
		BuildMap := func() {
			m = make(map[string]any, size)
			for k := range size {
				key := fmt.Sprintf(format, k)
				m[key] = nil
			}
		}
		b.Run(fmt.Sprintf("Native Map Put: %d", size), func(b *testing.B) {
			for range b.N {
				BuildMap()
			}
		})

		var ct *omap.CenterTree[string, any]
		BuildCt := func() {
			ct = omap.NewCenterTree[string, any](size, cmp.Compare)
			for k := range size {
				key := fmt.Sprintf(format, k)
				ct.Put(key, nil)
			}
		}
		b.Run(fmt.Sprintf("CenterTree Put: %d", size), func(b *testing.B) {
			for range b.N {
				BuildCt()
			}
		})
		var so *sm.SortedMap[map[string]any, string, any]
		BuildSm := func() {

			so = sm.New[map[string]any](func(i, j sm.KV[string, any]) bool { return cmp.Less(i.Key, j.Key) })
			for k := range size {
				key := fmt.Sprintf(format, k)
				so.Insert(key, nil)
			}
		}
		b.Run("sortedmap Put", func(b *testing.B) {
			for range b.N {

				BuildSm()
			}
		})

		var bt *btree.BTree
		BuildBt := func() {
			bt = btree.New(2)
			for k := range size {
				key := fmt.Sprintf(format, k)
				bt.ReplaceOrInsert(KV[any]{key: key, Value: nil})
			}
		}
		b.Run("btree Put", func(b *testing.B) {
			for range b.N {
				BuildBt()
			}
		})

		// Read tests
		b.Log("** Read tests")
		b.Run(fmt.Sprintf("Native Map Get: %d", size), func(b *testing.B) {
			for range b.N {

				for k := range size {
					key := fmt.Sprintf(format, k)
					_, _ = m[key]
				}
			}
		})
		b.Run(fmt.Sprintf("CenterTree Get: %d", size), func(b *testing.B) {
			for range b.N {
				for k := range size {
					key := fmt.Sprintf(format, k)
					ct.Get(key)
				}
			}
		})
		b.Run(fmt.Sprintf("sortedmap Get: %d", size), func(b *testing.B) {
			for range b.N {
				for k := range size {
					key := fmt.Sprintf(format, k)
					so.Get(key)
				}
			}
		})
		b.Run(fmt.Sprintf("btree Get: %d", size), func(b *testing.B) {
			for range b.N {
				for k := range size {
					key := fmt.Sprintf(format, k)
					bt.Get(KV[any]{key: key, Value: nil})
				}
			}
		})

		b.Log("Count all values from mid to end")

		mid := size / 2
		start := fmt.Sprintf(format, mid)
		end := fmt.Sprintf(format, size-1)
		b.Run(fmt.Sprintf("Native Count: %d", size), func(b *testing.B) {
			for range b.N {
				count := 0
				for k := range m {
					if k >= start {
						count++
					}
				}
			}
		})
		b.Run(fmt.Sprintf("CenterTree Count: %d", size), func(b *testing.B) {
			for range b.N {
				ct.Between(start, end, omap.LAST_KEY)
			}
		})
		b.Run(fmt.Sprintf("sortedmap Count: %d", size), func(b *testing.B) {
			for range b.N {
				count := 0
				for k := range so.All() {
					if k >= start {
						count++
					}
				}
			}

		})
		b.Run(fmt.Sprintf("btree Count: %d", size), func(b *testing.B) {
			for range b.N {
				count := 0
				start := fmt.Sprintf(format, mid-1)
				bt.DescendGreaterThan(KV[any]{key: start}, func(item btree.Item) bool {
					count++
					return true
				})
			}
		})

		b.Log("Delete first 999 elements")
		end = fmt.Sprintf(format, 998)

		b.Run(fmt.Sprintf("Native Mass Remove: %d", mid), func(b *testing.B) {
			for range b.N {
				for k := range mid {
					key := fmt.Sprintf(format, k)
					delete(m, key)
				}
				b.StopTimer()
				BuildMap()
				b.StartTimer()
			}
		})
		b.Run(fmt.Sprintf("CenterTree Mass Remove: %d", mid), func(b *testing.B) {
			for range b.N {
				ct.RemoveBetween("", end, omap.FIRST_KEY)
				b.StopTimer()
				BuildCt()
				b.StartTimer()
			}
		})
		b.Run(fmt.Sprintf("sortedmap Mass Remove: %d", mid), func(b *testing.B) {
			for range b.N {
				for k := range mid {
					key := fmt.Sprintf(format, k)
					so.Delete(key)
					b.StopTimer()
					BuildSm()
					b.StartTimer()
				}
			}
		})
		b.Run(fmt.Sprintf("btree Mass Remove: %d", mid), func(b *testing.B) {

			for range b.N {
				for k := range mid {
					key := fmt.Sprintf(format, k)
					bt.Delete(KV[any]{key: key, Value: nil})
					b.StopTimer()
					BuildBt()
					b.StartTimer()
				}
			}
		})
	}
}
