package styles

import "github.com/alecthomas/chroma/v2"

var (
	macchiatoFlamingo = "#F0C6C6"
	macchiatoPink     = "#F5BDE6"
	macchiatoMauve    = "#C6A0F6"
	macchiatoRed      = "#ED8796"
	macchiatoMaroon   = "#EE99A0"
	macchiatoPeach    = "#F5A97F"
	macchiatoYellow   = "#EED49F"
	macchiatoGreen    = "#A6DA95"
	macchiatoTeal     = "#8BD5CA"
	macchiatoSky      = "#91D7E3"
	macchiatoBlue     = "#8AADF4"
	macchiatoLavender = "#B7BDF8"
	macchiatoText     = "#CAD3F5"
	macchiatoOverlay0 = "#6E738D"
	macchiatoSurface2 = "#5B6078"
	macchiatoSurface0 = "#363A4F"
	macchiatoBase     = "#24273A"
)

// Catppuccin a soothing pastel theme for the high-spirited (macchiato variant)
var CatppuccinMacchiato = Register(chroma.MustNewStyle("catppuccin-macchiato", chroma.StyleEntries{
	chroma.TextWhitespace:        macchiatoSurface0,
	chroma.Comment:               "italic " + macchiatoSurface2,
	chroma.CommentPreproc:        macchiatoBlue,
	chroma.Keyword:               macchiatoMauve,
	chroma.KeywordPseudo:         "bold " + macchiatoMauve,
	chroma.KeywordType:           macchiatoFlamingo,
	chroma.KeywordConstant:       "italic " + macchiatoMauve,
	chroma.Operator:              macchiatoSky,
	chroma.OperatorWord:          "bold " + macchiatoSky,
	chroma.Name:                  macchiatoLavender,
	chroma.NameBuiltin:           "italic " + macchiatoText,
	chroma.NameFunction:          macchiatoSky,
	chroma.NameClass:             macchiatoYellow,
	chroma.NameNamespace:         macchiatoYellow,
	chroma.NameException:         macchiatoMaroon,
	chroma.NameVariable:          macchiatoPeach,
	chroma.NameConstant:          macchiatoYellow,
	chroma.NameLabel:             macchiatoYellow,
	chroma.NameEntity:            macchiatoPink,
	chroma.NameAttribute:         macchiatoYellow,
	chroma.NameTag:               macchiatoMauve,
	chroma.NameDecorator:         macchiatoPink,
	chroma.NameOther:             macchiatoPeach,
	chroma.Punctuation:           macchiatoText,
	chroma.LiteralString:         macchiatoGreen,
	chroma.LiteralStringDoc:      macchiatoGreen,
	chroma.LiteralStringInterpol: macchiatoGreen,
	chroma.LiteralStringEscape:   macchiatoBlue,
	chroma.LiteralStringRegex:    macchiatoBlue,
	chroma.LiteralStringSymbol:   macchiatoGreen,
	chroma.LiteralStringOther:    macchiatoGreen,
	chroma.LiteralNumber:         macchiatoTeal,
	chroma.GenericHeading:        "bold " + macchiatoSky,
	chroma.GenericSubheading:     "bold " + macchiatoSky,
	chroma.GenericDeleted:        macchiatoMaroon,
	chroma.GenericInserted:       macchiatoGreen,
	chroma.GenericError:          macchiatoMaroon,
	chroma.GenericEmph:           "italic",
	chroma.GenericStrong:         "bold",
	chroma.GenericPrompt:         "bold " + macchiatoOverlay0,
	chroma.GenericOutput:         macchiatoPeach,
	chroma.GenericTraceback:      macchiatoMaroon,
	chroma.Error:                 macchiatoRed,
	chroma.Background:            macchiatoPeach + " bg:" + macchiatoBase,
}))
