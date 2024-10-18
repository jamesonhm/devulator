package scpi

type objNative struct {
	function nativeFn
}

type nativeFn func(argCount int, args []value) value

func newNative(function nativeFn) *objNative {
	return &objNative{
		function: function,
	}
}
