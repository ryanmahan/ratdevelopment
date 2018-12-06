package searching

import (
	"testing"
)

var context SearchContext

func init() {
	context = SearchContext{
		Systems: []System{
			System{capacity: 50, id: "9996788", tenants: []string{"hpe", "jake"}},
			System{capacity: 35, id: "9998855", tenants: []string{"hpe"}},
			System{capacity: 65, id: "8836722", tenants: []string{"hpe", "john"}},
		},
	}
}

func TestEnclosing(t *testing.T) {
	results := context.Search("8, capacity > 50")
	t.Log(len(results))
	results = context.Search("j, capacity < 55")
	t.Log(len(results))
}
