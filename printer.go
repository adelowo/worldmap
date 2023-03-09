package alien

import "os"

type Printer interface {
	Print(string) error
}

type StdOutPrinter struct {
}

func (ss *StdOutPrinter) Print(s string) error {
	_, err := os.Stdout.WriteString(s)
	return err
}

type NoopPrinter struct{}

func (_ *NoopPrinter) Print(_ string) error { return nil }

var defaultPrinter = &StdOutPrinter{}

func print(s string) error {
	return defaultPrinter.Print(s)
}
