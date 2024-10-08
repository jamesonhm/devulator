package scpi

type void struct{}

var m void

var kws = map[string]void{
	"ABOR":      m,
	"ABORT":     m,
	"AFR":       m,
	"CALC":      m,
	"CALCULATE": m,
}

type state struct {
}
