package styles

import "github.com/alecthomas/chroma/v2"

var (
	frappeFlamingo = "#EEBEBE"
	frappePink     = "#F4B8E4"
	frappeMauve    = "#CA9EE6"
	frappeRed      = "#E78284"
	frappeMaroon   = "#EA999C"
	frappePeach    = "#EF9F76"
	frappeYellow   = "#E5C890"
	frappeGreen    = "#A6D189"
	frappeTeal     = "#81C8BE"
	frappeSky      = "#99D1DB"
	frappeBlue     = "#8CAAEE"
	frappeLavender = "#BABBF1"
	frappeText     = "#c6d0f5"
	frappeOverlay0 = "#737994"
	frappeSurface2 = "#626880"
	frappeSurface0 = "#414559"
	frappeBase     = "#303446"
)

// Catppuccin a soothing pastel theme for the high-spirited (frappe variant)
var CatppuccinFrappe = Register(chroma.MustNewStyle("catppuccin-frappe", chroma.StyleEntries{
	chroma.TextWhitespace:        frappeSurface0,
	chroma.Comment:               "italic " + frappeSurface2,
	chroma.CommentPreproc:        frappeBlue,
	chroma.Keyword:               frappeMauve,
	chroma.KeywordPseudo:         "bold " + frappeMauve,
	chroma.KeywordType:           frappeFlamingo,
	chroma.KeywordConstant:       "italic " + frappeMauve,
	chroma.Operator:              frappeSky,
	chroma.OperatorWord:          "bold " + frappeSky,
	chroma.Name:                  frappeLavender,
	chroma.NameBuiltin:           "italic " + frappeText,
	chroma.NameFunction:          frappeSky,
	chroma.NameClass:             frappeYellow,
	chroma.NameNamespace:         frappeYellow,
	chroma.NameException:         frappeMaroon,
	chroma.NameVariable:          frappePeach,
	chroma.NameConstant:          frappeYellow,
	chroma.NameLabel:             frappeYellow,
	chroma.NameEntity:            frappePink,
	chroma.NameAttribute:         frappeYellow,
	chroma.NameTag:               frappeMauve,
	chroma.NameDecorator:         frappePink,
	chroma.NameOther:             frappePeach,
	chroma.Punctuation:           frappeText,
	chroma.LiteralString:         frappeGreen,
	chroma.LiteralStringDoc:      frappeGreen,
	chroma.LiteralStringInterpol: frappeGreen,
	chroma.LiteralStringEscape:   frappeBlue,
	chroma.LiteralStringRegex:    frappeBlue,
	chroma.LiteralStringSymbol:   frappeGreen,
	chroma.LiteralStringOther:    frappeGreen,
	chroma.LiteralNumber:         frappeTeal,
	chroma.GenericHeading:        "bold " + frappeSky,
	chroma.GenericSubheading:     "bold " + frappeSky,
	chroma.GenericDeleted:        frappeMaroon,
	chroma.GenericInserted:       frappeGreen,
	chroma.GenericError:          frappeMaroon,
	chroma.GenericEmph:           "italic",
	chroma.GenericStrong:         "bold",
	chroma.GenericPrompt:         "bold " + frappeOverlay0,
	chroma.GenericOutput:         frappePeach,
	chroma.GenericTraceback:      frappeMaroon,
	chroma.Error:                 frappeRed,
	chroma.Background:            frappePeach + " bg:" + frappeBase,
}))
