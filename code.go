package ogs

// CodeOK default ok code is 0, you can change it
var CodeOK interface{} = 0

// ChangeOKCode change the value of CodeOK
func ChangeOKCode(v interface{}) {
	CodeOK = v
}
