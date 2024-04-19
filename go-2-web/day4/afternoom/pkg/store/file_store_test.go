package store

import (
	"os"
	"testing"
)

const filename = "testing.json"

func TestRead(test *testing.T) {
	var db FileStore = FileStore{filename}
	data := []struct {
		name    string
		age     int
		testing float32
	}{{
		name:    "luiz",
		age:     19,
		testing: 4.5,
	}, {
		name:    "Milena",
		age:     20,
		testing: 4.6,
	},
	}
	os.Create(filename)
	test.Run("testando se a resposta do context funciona", func(t *testing.T) {
		err := db.Read(&data)
		if err != nil {
			t.Errorf("the read was not possible %s", err.Error())
		}
	})
	os.Remove(filename)

	test.Run("testando se o erro funciona", func(t *testing.T) {
		err := db.Read(&data)
		if err != nil {
			t.Errorf("the read was not possible %s", err.Error())
		}
	})
}
