package scpi

//import "fmt"

type InterpretResult int

type VM struct {
	chunk *Chunk
	ip    uint8
}

const (
	INTERPRET_OK InterpretResult = iota
	INTERPRET_COMPILE_ERROR
	INTERPRET_RUNTIME_ERROR
)

func (vm *VM) Interpret(src string) InterpretResult {
	chunk := InitChunk()

	if !compile(src, &chunk) {
		return INTERPRET_COMPILE_ERROR
	}

	vm.chunk = &chunk
	vm.ip = vm.chunk.code[0]

	//	var result InterpretResult = run()

	return INTERPRET_OK
}

type void struct{}

var v void

func wordInSet(s string, set map[string]void) bool {
	_, ok := set[s]
	if ok {
		return true
	}
	return false
}

var SPECIALS = map[string]void{
	"MIN": v,
	"MAX": v,
	"DEF": v,
}

var BOOLS = map[string]void{
	"ON":  v,
	"OFF": v,
}

var COMMONS = map[string]void{
	"*CLS": v,
	"*DMC": v,
	"*EMC": v,
	"*ESE": v,
	"*ESR": v,
	"*IDN": v,
	"*LRN": v,
	"*OPC": v,
	"*PCB": v,
	"*RST": v,
	"*SRE": v,
	"*STB": v,
	"*TRG": v,
	"*TST": v,
	"*WAI": v,
}

var kws = map[string]void{
	"ABOR":      v,
	"ABORT":     v,
	"AFR":       v,
	"CALC":      v,
	"CALCULATE": v,
	"CALP":      v,
	"CALPOD":    v,
	"CONT":      v,
	"CONTROL":   v,
	"CSET":      v,
	"DISP":      v,
	"DISTLAY":   v,
	"FORM":      v,
	"FORMAT":    v,
	"HCOP":      v,
	"HCOPY":     v,
	"INIT":      v,
	"INITIATE":  v,
	"LXI":       v,
	"MMEM":      v,
	"MMEMORY":   v,
	"OUTP":      v,
	"OUTPUT":    v,
	"SENS":      v,
	"SENSE":     v,
	"SOUR":      v,
	"SOURCE":    v,
	"STAT":      v,
	"STATUS":    v,
	"SYST":      v,
	"SYSTEM":    v,
}

type deviceFn func(args ...value) value

var deviceFns = map[string]deviceFn{
	"MEASUREVOLTAGERISETIME": measureVoltageRiseTime,
}

func measureVoltageRiseTime(arg value) value {
	return value{}
}
