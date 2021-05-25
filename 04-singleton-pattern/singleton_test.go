package singleton

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	s := GetDocumentHelper()
	s1 := GetDocumentHelper()
	fmt.Println(s == s1)
}
