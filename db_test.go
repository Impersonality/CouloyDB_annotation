package CouloyDB

import (
	"bytes"
	"fmt"
	"log"
	"testing"
)

func TestCouloyDB(t *testing.T) {
	conf := DefaultOptions()
	db, err := NewCouloyDB(conf)
	if err != nil {
		log.Fatal(err)
	}

	key := []byte("first key")
	value := []byte("first value")

	// Be careful, you can't use single non-displayable character in ASCII code as your key (0x00 ~ 0x1F and 0x7F),
	// because those characters will be used in CouloyDB as necessary operations in the preset key tagging system.
	// This may be changed in the next major release
	err = db.Put(key, value)
	if err != nil {
		log.Fatal(err)
	}

	v, err := db.Get(key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)

	err = db.Del(v)
	if err != nil {
		log.Fatal(err)
	}

	keys := db.ListKeys()
	for _, k := range keys {
		fmt.Println(k)
	}

	target := []byte("target")
	err = db.Fold(func(key []byte, value []byte) bool {
		if bytes.Equal(value, target) {
			fmt.Println("Include target")
			return false
		}
		return true
	})
	if err != nil {
		log.Fatal(err)
	}
}
