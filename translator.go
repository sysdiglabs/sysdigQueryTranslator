package main

import "C"
import . "github.com/sysdiglabs/sysdigQueryTranslator/pkg/translator"

func main() {
}

//export RemoveSysdigLabelsInExpression
func RemoveSysdigLabelsInExpression(cExpr, cExcludeScope *C.char) *C.char {
	expr := C.GoString(cExpr)
	excludeScope := C.GoString(cExcludeScope)
	expr = RemoveSysdigLabels(expr, excludeScope)
	cStr := C.CString(expr)
	return cStr
}
