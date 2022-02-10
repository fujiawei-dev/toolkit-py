{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"fmt"
	"testing"
)

func TestSplit(t *testing.T) {
	fmt.Println(Split("c:/images/cover.jpg"))
}

func TestListDir(t *testing.T) {
	fmt.Println(ListDir("."))
	fmt.Println(ListDir2("."))
}

func TestMustRel(t *testing.T) {
	fmt.Println(MustRel("/mnt/extra/images", "meizi"))
}
