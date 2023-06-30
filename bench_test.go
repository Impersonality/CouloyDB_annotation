package CouloyDB

import (
	"github.com/Kirov7/CouloyDB/public"
	"github.com/Kirov7/CouloyDB/public/utils/bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Benchmark_Put(b *testing.B) {
	options := DefaultOptions()
	options.SetDataFileSizeByte(8 * 1024 * 1024)
	options.SetSyncWrites(false)
	testDB, err := NewCouloyDB(options)
	assert.Nil(b, err)
	assert.NotNil(b, testDB)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		err := testDB.Put(bytes.IntToBytes(i), bytes.RandomBytes(1024))
		assert.Nil(b, err)
	}

	b.StopTimer()
	destroyCouloyDB(testDB)
}

func Benchmark_Get(b *testing.B) {
	options := DefaultOptions()
	options.SetDataFileSizeByte(8 * 1024 * 1024)
	testDB, err := NewCouloyDB(options)
	assert.Nil(b, err)
	assert.NotNil(b, testDB)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err = testDB.Get(bytes.IntToBytes(i))
		if err != nil && err != public.ErrKeyNotFound {
			b.Fatal(err)
		}
	}

	b.StopTimer()
	testDB.Close()
}
