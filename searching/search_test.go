package searching

import (
	"fmt"
	"testing"
)

var context SearchContext

func init() {
	context = SearchContext{
		Systems: []System{
			System{Capacity: 50, Id: "9996788", Company: "Tr. Wu"},
			System{Capacity: 35, Id: "9998855", Company: "Whoa Whoa"},
			System{Capacity: 65, Id: "8836722", Company: "That sounds like a death trap"},
		},
	}
}

func TestEnclosing(t *testing.T) {
	results := context.Search("8, capacity > 50")
	t.Log(len(results))
	results = context.Search("T, capacity < 75")
	t.Log(len(results))
}

func TestConverstion(t *testing.T) {
	result := SearchQueryToCQL("Longneck, 99967")
	fmt.Println(result)
	t.Error(result)
}
