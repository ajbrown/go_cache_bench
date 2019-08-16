package go_cache_bench

import (
	"fmt"
	"github.com/coocood/freecache"
	"github.com/goburrow/cache"
	"math/rand"
	"testing"
)

var itemSizes = []int{
	100,
	1000,
	10000,
}

func randBytes(size int) []byte {
	b := make([]byte, size)
	for i := 0; i < size; i++ {
		b[i] = byte(rand.Int())
	}

	return b
}

func BenchmarkSingleKey_Get(b *testing.B) {

	for _, s := range itemSizes {
		b.Run(fmt.Sprintf("FreeCache/keySize=%d", s), func(b *testing.B) {
			fc := freecache.NewCache(s * 1024)
			v := randBytes(s)
			k := []byte("key")
			_ = fc.Set(k, v, -1)

			b.ResetTimer()

			for r := 0; r < b.N; r++ {
				_, _ = fc.Get(k)
			}
		})

		b.Run(fmt.Sprintf("MangoStandard/keySize=%d", s), func(b *testing.B) {
			gbc := cache.New(cache.WithMaximumSize(s + 10))
			v := randBytes(s)
			k := "key"
			gbc.Put(k, v)

			b.ResetTimer()

			for r := 0; r < b.N; r++ {
				_, _ = gbc.GetIfPresent(k)
			}
		})

		b.Run(fmt.Sprintf("MangoLoading/keySize=%d", s), func(b *testing.B) {
			b.SkipNow()
			gbc := cache.NewLoadingCache(func(k cache.Key) (cache.Value, error) { return randBytes(s), nil }, cache.WithMaximumSize(s+10))

			v := randBytes(s)
			k := "key"
			gbc.Put(k, v)

			b.ResetTimer()

			for r := 0; r < b.N; r++ {
				_, _ = gbc.Get(k)
			}
		})
	}
}

func BenchmarkSingleKey_SetAndGet(b *testing.B) {
	for _, s := range itemSizes {
		b.Run(fmt.Sprintf("FreeCache/keySize=%d", s), func(b *testing.B) {
			var v []byte
			fc := freecache.NewCache(s * 1024)
			k := []byte("key")

			b.ResetTimer()

			for r := 0; r < b.N; r++ {
				b.StopTimer()
				v = randBytes(s)
				b.StartTimer()

				_ = fc.Set(k, v, -1)
				_, _ = fc.Get(k)
			}
		})

		b.Run(fmt.Sprintf("MangoStandard/keySize=%d", s), func(b *testing.B) {
			var v []byte
			gbc := cache.New(cache.WithMaximumSize(s + 10))
			k := "key"

			b.ResetTimer()

			for r := 0; r < b.N; r++ {
				b.StopTimer()
				v = randBytes(s)
				b.StartTimer()

				gbc.Put(k, v)
				_, _ = gbc.GetIfPresent(k)
			}
		})

		b.Run(fmt.Sprintf("MangoLoading/keySize=%d", s), func(b *testing.B) {
			b.SkipNow()

			var v []byte
			gbc := cache.NewLoadingCache(func(k cache.Key) (cache.Value, error) { return randBytes(s), nil }, cache.WithMaximumSize(s+10))
			k := "key"

			b.ResetTimer()

			for r := 0; r < b.N; r++ {
				b.StopTimer()
				v = randBytes(s)
				b.StartTimer()

				gbc.Put(k, v)
				_, _ = gbc.Get(k)
			}
		})
	}
}

// This benchmark represents a more typical behavior in a service or website.
// - lookup the value
// - write a new value
// - look up the value multiple times
func BenchmarkSingleKey_GetSetMultiGet(b *testing.B) {
	var numGets = []int{
		2,
		4,
		8,
	}

	for _, s := range itemSizes {

		for _, ng := range numGets {
			b.Run(fmt.Sprintf("FreeCache/keySize=%d/reads=%d", s, ng), func(b *testing.B) {
				var v []byte
				fc := freecache.NewCache(s * 1024)
				k := []byte("key")

				b.ResetTimer()

				for r := 0; r < b.N; r++ {
					b.StopTimer()
					v = randBytes(s)
					b.StartTimer()

					_, _ = fc.Get(k)
					_ = fc.Set(k, v, -1)
					for i := 0; i < ng; i++ {
						_, _ = fc.Get(k)
					}
				}
			})

			b.Run(fmt.Sprintf("MangoStandard/keySize=%d/reads=%d", s, ng), func(b *testing.B) {
				var v []byte
				gbc := cache.New(cache.WithMaximumSize(s + 10))
				k := "key"

				b.ResetTimer()

				for r := 0; r < b.N; r++ {
					b.StopTimer()
					v = randBytes(s)
					b.StartTimer()

					gbc.GetIfPresent(k)
					gbc.Put(k, v)
					for i := 0; i < ng; i++ {
						_, _ = gbc.GetIfPresent(k)
					}
				}
			})

			b.Run(fmt.Sprintf("MangoLoading/keySize=%d/reads=%d", s, ng), func(b *testing.B) {
				b.SkipNow()
				var v []byte
				gbc := cache.NewLoadingCache(func(k cache.Key) (cache.Value, error) { return randBytes(s), nil }, cache.WithMaximumSize(s+10))
				k := "key"

				b.ResetTimer()

				for r := 0; r < b.N; r++ {
					b.StopTimer()
					v = randBytes(s)
					b.StartTimer()

					_, _ = gbc.Get(k)
					gbc.Put(k, v)
					for i := 0; i < ng; i++ {
						_, _ = gbc.Get(k)
					}
				}
			})

		}
	}
}
