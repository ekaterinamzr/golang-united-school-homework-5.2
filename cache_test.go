package cache

import (
	"testing"
	"time"
)

type testStruct struct {
	key      string
	value    string
	deadline time.Time
}

const expiration = 2e9

func TestGet(t *testing.T) {
	tests := []testStruct{{key: "Name", value: "Kate"},
		{key: "Uni", value: "BMSTU", deadline: time.Now().Add(expiration * 2)},
		{key: "Color", value: "Blue", deadline: time.Now().Add(expiration)},
	}

	cache := NewCache()

	for _, test := range tests {
		if test.deadline.IsZero() {
			cache.Put(test.key, test.value)
		} else {
			cache.PutTill(test.key, test.value, test.deadline)
		}
	}

	time.Sleep(expiration)

	for _, test := range tests {
		value, ok := cache.Get(test.key)
		if !ok && (test.deadline.IsZero() || time.Now().Before(test.deadline)) {
			t.Error("Key ", test.key, "not found")
		} else if ok && !test.deadline.IsZero() && !time.Now().Before(test.deadline) {
			t.Error("Key ", test.key, "expired but found")
		} else if ok && value != test.value {
			t.Error("For key ", test.key, "expected ", test.value, "got ", value)
		}
	}
}
