package translations

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Translator struct {
	printer *message.Printer
}

func (l *Translator) Translate(key message.Reference, args ...any) string {
	return l.printer.Sprintf(key, args...)
}

var EnTranslator *Translator
var PtTranslator *Translator

func init() {
	EnTranslator = &Translator{message.NewPrinter(language.MustParse("en-US"))}
	PtTranslator = &Translator{message.NewPrinter(language.MustParse("pt-BR"))}
}
