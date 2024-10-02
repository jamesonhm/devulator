package scpi

type Scanner struct {
	start   *rune
	current *rune
	line    uint
}

func InitScanner(source string) {

}

func Scan(src string) []string {
	return []string{"from", "the", "scanner"}
}
