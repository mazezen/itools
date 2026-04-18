package ilo

import (
	"fmt"
	"testing"
)

func TestPaginate(t *testing.T) {
	items := []string{"a", "b", "c", "d", "e"}

  	res := Paginate(items, 3, 2)
	fmt.Printf("%v", res)
}