package scpi

//import "fmt"

type InterpretResult int

type VM struct {
}

const (
	INTERPRET_OK InterpretResult = iota
	INTERPRET_COMPILE_ERROR
	INTERPRET_RUNTIME_ERROR
)

func (vm *VM) Interpret(src string) InterpretResult {
	compile(src)
	return INTERPRET_OK
}

type void struct{}

var m void

func wordInSet(s string, set map[string]void) bool {
	_, ok := set[s]
	if ok {
		return true
	}
	return false
}

var SPECIALS = map[string]void{
	"MIN": m,
	"MAX": m,
	"DEF": m,
}

var BOOLS = map[string]void{
	"ON":  m,
	"OFF": m,
}

var COMMONS = map[string]void{
	"*CLS": m,
	"*DMC": m,
	"*EMC": m,
	"*ESE": m,
	"*ESR": m,
	"*IDN": m,
	"*LRN": m,
	"*OPC": m,
	"*PCB": m,
	"*RST": m,
	"*SRE": m,
	"*STB": m,
	"*TRG": m,
	"*TST": m,
	"*WAI": m,
}

var kws = map[string]void{
	"ABOR":      m,
	"ABORT":     m,
	"AFR":       m,
	"CALC":      m,
	"CALCULATE": m,
	"CALP":      m,
	"CALPOD":    m,
	"CONT":      m,
	"CONTROL":   m,
	"CSET":      m,
	"DISP":      m,
	"DISTLAY":   m,
	"FORM":      m,
	"FORMAT":    m,
	"HCOP":      m,
	"HCOPY":     m,
	"INIT":      m,
	"INITIATE":  m,
	"LXI":       m,
	"MMEM":      m,
	"MMEMORY":   m,
	"OUTP":      m,
	"OUTPUT":    m,
	"SENS":      m,
	"SENSE":     m,
	"SOUR":      m,
	"SOURCE":    m,
	"STAT":      m,
	"STATUS":    m,
	"SYST":      m,
	"SYSTEM":    m,
}
