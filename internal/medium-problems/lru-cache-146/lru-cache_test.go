package lru_cache_146

import (
	"fmt"
	"testing"
)

func TestLruCache_Get(t *testing.T) {
	tests := []struct {
		name            string
		expectedError   error
		initialiseCache func() Cache
		lookupKey       string
		expectedValue   string
	}{
		{
			name:          "lookup on nil entry, assert error",
			expectedError: noElementError{},
			initialiseCache: func() Cache {
				return New(10)
			},
			lookupKey:     "",
			expectedValue: "",
		},
		{
			name:          "get element, no error",
			expectedError: nil,
			initialiseCache: func() Cache {
				c := New(10)
				c.Put("hello", "world")
				return c
			},
			lookupKey:     "hello",
			expectedValue: "world",
		},
		{
			name:          "call to Get() updates mru node",
			expectedError: noElementError{},
			initialiseCache: func() Cache {
				c := New(2)
				c.Put("blue", "jean")
				c.Put("let's", "dance")
				_, _ = c.Get("let's")
				c.Put("china", "girl")
				return c
			},
			lookupKey:     "blue",
			expectedValue: "jean",
		},
		//{
		//	name: "Get on an uninitialised cache",
		//},
	}

	for i, test := range tests {
		fmt.Printf("test %d: %s\n", i, test.name)
		c := test.initialiseCache()
		str, err := c.Get(test.lookupKey)
		if err != nil && test.expectedError != nil {
			fmt.Printf("\tPASS: %s\n", err.Error())
			continue
		}

		if str != nil && test.expectedValue != *str {
			t.Errorf("FAIL: expected %s, got %s\n", test.expectedValue, *str)
			continue
		} else {
			fmt.Printf("\tPASS\n")
		}
	}
}

func TestLruCache_Put(t *testing.T) {
	v := func(s string) *string {
		return &s
	}
	tests := []struct {
		name            string
		expectedError   error
		initialiseCache func() Cache
		lookupKey       string
		expectedValue   *string
	}{
		//{
		//	name:          "initialise cache with capacity of 0, cause panic",
		//	expectedError: nil,
		//	initialiseCache: func() Cache {
		//		c := New(0)
		//		return c
		//	},
		//},
		{
			name:          "call to Put()",
			expectedError: nil,
			initialiseCache: func() Cache {
				c := New(1)
				c.Put("let's", "dance")
				return c
			},
			lookupKey:     "let's",
			expectedValue: v("dance"),
		},
		{
			name:          "call to Put(), testing cache eviction policy",
			expectedError: noElementError{},
			initialiseCache: func() Cache {
				c := New(1)
				c.Put("let's", "dance")
				c.Put("china", "girl")
				return c
			},
			lookupKey:     "let's",
			expectedValue: nil,
		},
		{
			name:          "call to Put(), add 2 elements to cache",
			expectedError: nil,
			initialiseCache: func() Cache {
				c := New(2)
				c.Put("let's", "dance")
				c.Put("china", "girl")
				return c
			},
			lookupKey:     "let's",
			expectedValue: v("dance"),
		},
		{
			name:          "call to Put(), testing cache eviction policy",
			expectedError: noElementError{},
			initialiseCache: func() Cache {
				c := New(3)
				c.Put("let's", "dance")
				c.Put("blue", "jean")
				c.Put("china", "girl")
				_, _ = c.Get("let's")
				_, _ = c.Get("blue")
				c.Put("diamond", "dogs")
				return c
			},
			lookupKey:     "china",
			expectedValue: nil,
		},
	}

	//defer func() error {
	//	if recover() != nil {
	//		return errors.New("capacity must be at least 1")
	//	}
	//	return nil
	//}()

	for i, test := range tests {
		fmt.Printf("test %d: %s\n", i, test.name)

		c := test.initialiseCache()

		resultVal, err := c.Get(test.lookupKey)

		if err != nil && test.expectedError == nil {
			t.Errorf("FAIL: unsuccessful Get() on key %s", test.lookupKey)
		}
		if resultVal == nil && test.expectedValue == nil {
			fmt.Printf("\tPASS\n")
			continue
		}
		if *resultVal != *test.expectedValue {
			t.Errorf("FAIL: expected: %s, got: %s", *test.expectedValue, *resultVal)
		} else {
			fmt.Printf("\tPASS\n")
		}
	}
}

func TestExtraTestCases(t *testing.T) {
	c := New(2)
	c.Put("blue", "jean")        // cache is [{"blue":"jean"}]
	c.Put("let's", "dance")      // cache is [{"blue":"jean"}, {"let's":"dance"}]
	result, err := c.Get("blue") // return "jean", cache is now [{"let's":"dance"}, {"blue":"jean"}]
	if *result != "jean" || err != nil {
		t.Errorf("FAIL: expected %s, got %s\n", "jean", *result)
	}
	c.Put("china", "girl")       // LRU key was "let's", evicts key "let's", cache is [{"blue":"jean"}, {"china":"girl"}]
	result, err = c.Get("let's") // returns -1 (not found)
	if result != nil || err.Error() != new(noElementError).Error() {
		t.Errorf("FAIL: expected %v, got %v\n", nil, result)
	}
	c.Put("diamond", "dogs")    // LRU key was "blue", evicts key "blue", cache is [{"china":"girl"}, {"diamond":"dogs"}]
	result, err = c.Get("blue") // return noElementError (not found)
	if result != nil || err.Error() != new(noElementError).Error() {
		t.Errorf("FAIL: expected %v, got %v\n", nil, result)
	}
	result, err = c.Get("china") // return "girl", cache is [{"diamond":"dogs"}, {"china":"girl"}]
	if *result != "girl" || err != nil {
		t.Errorf("FAIL: expected %s, got %s\n", "girl", *result)
	}
	result, err = c.Get("diamond") // return 4, cache is [{"china":"girl"}, {"diamond":"dogs"}]
	if *result != "dogs" || err != nil {
		t.Errorf("FAIL: expected %s, got %s\n", "dogs", *result)
	}
}

func TestExtraTestCases2(t *testing.T) {
	c := New(2)
	c.Put("blue", "jean")         // cache is [{"blue":"jean"}]
	c.Put("let's", "dance")       // cache is [{"blue":"jean"}, {"let's":"dance"}]
	c.Put("blue", "jeans")        // cache is [{"let's":"dance"}, {"blue":"jeans"}]
	c.Put("china", "girl")        // cache is [{"blue":"jeans"}, {"china":"girl"}]
	result, err := c.Get("let's") // returns error
	if result != nil || err.Error() != new(noElementError).Error() {
		t.Errorf("FAIL: expected %v, got %s\n", nil, *result)
	}
	result, err = c.Get("blue") // returns "jeans", cache becomes [{"china":"girl"}, {"blue":"jeans"}]
	if *result != "jeans" || err != nil {
		t.Errorf("FAIL: expected %s, got %s\n", "jeans", *result)
	}
}

func TestExtraTestCases3(t *testing.T) {
	c := New(2)
	result, err := c.Get("let's") // cache is []
	if result != nil || err.Error() != new(noElementError).Error() {
		t.Errorf("FAIL: expected %v, got %s\n", nil, *result)
	}
	c.Put("let's", "dance")     // cache is [{"let's": &listNode1{val: "dance", next: nil}}]
	result, err = c.Get("blue") // returns -1
	if result != nil || err.Error() != new(noElementError).Error() {
		t.Errorf("FAIL: expected %v, got %s\n", nil, *result)
	}
	c.Put("blue", "jean")       // cache: [{"let's": &listNode1{val: "dance", next: listNode2}}, {"blue": &listNode2{val: "jean", next: nil}}]
	c.Put("blue", "jeans")      // cache: [{"let's": &listNode1{val: "dance", next: listNode2}}, {"blue": &listNode2{val: "jeans", next: nil}}]
	result, err = c.Get("blue") // returns "jeans"
	if *result != "jeans" || err != nil {
		t.Errorf("FAIL: expected %s, got %s\n", "jeans", *result)
	}
	result, err = c.Get("let's") // returns 6, // cache: [{"blue": &listNode2{val: "jeans", next: listNode1}}, {"let's": &listNode1{val: "dance", next: nil}}]
	if *result != "dance" || err != nil {
		t.Errorf("FAIL: expected %s, got %s\n", "dance", *result)
	}
}
