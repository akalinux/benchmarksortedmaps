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
	for _, size := range []int{1000} {
		b.Logf("Working with set size: %d", size)
		var m map[string]any
		// Write tests
		b.Log("** Write tests")
		BuildMap := func(size int) {
			m = make(map[string]any, size)
			for k := range size {
				key := fmt.Sprintf(format, k)
				m[key] = nil
			}
		}
		b.Run("Native Map Put", func(b *testing.B) {
			for range b.N {
				BuildMap(size)
			}
		})

		var ct *omap.CenterTree[string, any]
		BuildCt := func(size int) {
			ct = omap.NewCenterTree[string, any](size, cmp.Compare)
			for k := range size {
				key := fmt.Sprintf(format, k)
				ct.Put(key, nil)
			}
		}
		b.Run("CenterTree Put", func(b *testing.B) {
			for range b.N {
				BuildCt(size)
			}
		})
		var so *sm.SortedMap[map[string]any, string, any]
		BuildSm := func(size int) {

			so = sm.New[map[string]any](func(i, j sm.KV[string, any]) bool { return cmp.Less(i.Key, j.Key) })
			for k := range size {
				key := fmt.Sprintf(format, k)
				so.Insert(key, nil)
			}
		}
		b.Run("sortedmap Put", func(b *testing.B) {
			for range b.N {
				BuildSm(size)
			}
		})

		var bt *btree.BTree
		BuildBt := func(size int) {
			bt = btree.New(2)
			for k := range size {
				key := fmt.Sprintf(format, k)
				bt.ReplaceOrInsert(KV[any]{key: key, Value: nil})
			}
		}
		b.Run("btree Put", func(b *testing.B) {
			for range b.N {
				BuildBt(size)
			}
		})

		// Read tests
		b.Log("** Single Read tests")
		b.Run("Native Map Ge", func(b *testing.B) {
			for range b.N {

				for k := range size {
					key := fmt.Sprintf(format, k)
					_, _ = m[key]
				}
			}
		})
		b.Run("CenterTree Get", func(b *testing.B) {
			for range b.N {
				for k := range size {
					key := fmt.Sprintf(format, k)
					ct.Get(key)
				}
			}
		})
		b.Run("sortedmap Get", func(b *testing.B) {
			for range b.N {
				for k := range size {
					key := fmt.Sprintf(format, k)
					so.Get(key)
				}
			}
		})
		b.Run("btree Get", func(b *testing.B) {
			for range b.N {
				for k := range size {
					key := fmt.Sprintf(format, k)
					bt.Get(KV[any]{key: key, Value: nil})
				}
			}
		})

		b.Log("Count all values from begin to mid")

		mid := size / 2
		start := fmt.Sprintf(format, mid)
		end := fmt.Sprintf(format, mid/2)
		b.Run("Native Count", func(b *testing.B) {
			for range b.N {
				count := 0
				for k := range m {
					if k >= start && k <= end {
						count++
					}
				}
			}
		})
		b.Run("CenterTree Count", func(b *testing.B) {
			for range b.N {
				ct.Between(start, end, omap.LAST_KEY)
			}
		})
		b.Run("sortedmap Count", func(b *testing.B) {
			for range b.N {
				count := 0
				for k := range so.All() {
					if k >= start && k <= end {
						count++
					} else {
						break
					}
				}
			}

		})
		b.Run("btree Count", func(b *testing.B) {
			for range b.N {
				count := 0
				bt.Ascend(func(item btree.Item) bool {
					kv, _ := item.(KV[any])
					count++
					return kv.key > end
				})
			}
		})

		b.Log("Find and delete first half of all elements")
		for _, cb := range []func(int){BuildBt, BuildCt, BuildMap, BuildSm} {
			cb(size)
		}

		b.Run("Native Mass Remove", func(b *testing.B) {
			for range b.N {

				for k := range m {
					if k <= end {
						delete(m, k)
					} else {
						break
					}
				}

				b.StopTimer()
				BuildMap(size)
				b.StartTimer()
			}
		})
		b.Run("CenterTree Mass Remove", func(b *testing.B) {
			for range b.N {
				ct.RemoveBetween("", end, omap.FIRST_KEY)
				b.StopTimer()
				BuildCt(size)
				b.StartTimer()
			}
		})
		b.Run("sortedmap Mass Remove", func(b *testing.B) {
			for range b.N {
				for k := range so.Keys() {
					key := fmt.Sprintf(format, k)
					so.Delete(key)
					if k > end {
						break
					}
				}
				b.StopTimer()
				BuildSm(size)
				b.StartTimer()
			}
		})
		b.Run("btree Mass Remove", func(b *testing.B) {

			for range b.N {
				bt.Ascend(
					func(item btree.Item) bool {
						kv, _ := item.(KV[any])
						bt.Delete(item)
						return kv.key > end
					},
				)
				b.StopTimer()
				BuildBt(size)
				b.StartTimer()
			}
		})
	}
}
