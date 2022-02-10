{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import "testing"

func TestPathWritable(t *testing.T) {
	t.Log(PathWritable("c:/developer"))
}
