package searching

import (
	"testing"
)

var context SearchContext

func init() {
	context = SearchContext{
		Systems: []System{
			System{capacity: 50, id: "9996788"},
			System{capacity: 35, id: "9998855"},
			System{capacity: 65, id: "8836722"},
		},
	}
}

func TestEnclosing(t *testing.T) {
	results := context.Search("capacity < 50")
	t.Log(len(results))
}
