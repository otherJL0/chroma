package styles

import "github.com/alecthomas/chroma/v2"

var (
	latteFlamingo = "#DD7878"
	lattePink     = "#ea76cb"
	latteMauve    = "#8839EF"
	latteRed      = "#D20F39"
	latteMaroon   = "#E64553"
	lattePeach    = "#FE640B"
	latteYellow   = "#df8e1d"
	latteGreen    = "#40A02B"
	latteTeal     = "#179299"
	latteSky      = "#04A5E5"
	latteBlue     = "#1e66f5"
	latteLavender = "#7287FD"
	latteText     = "#4C4F69"
	latteOverlay0 = "#9CA0B0"
	latteSurface2 = "#ACB0BE"
	latteSurface0 = "#CCD0DA"
	latteBase     = "#EFF1F5"
)

// Catppuccin a soothing pastel theme for the high-spirited (latte variant)
var CatppuccinLatte = Register(chroma.MustNewStyle("catppuccin-latte", chroma.StyleEntries{
	chroma.TextWhitespace:        latteSurface0,
	chroma.Comment:               "italic " + latteSurface2,
	chroma.CommentPreproc:        latteBlue,
	chroma.Keyword:               latteMauve,
	chroma.KeywordPseudo:         "bold " + latteMauve,
	chroma.KeywordType:           latteFlamingo,
	chroma.KeywordConstant:       "italic " + latteMauve,
	chroma.Operator:              latteSky,
	chroma.OperatorWord:          "bold " + latteSky,
	chroma.Name:                  latteLavender,
	chroma.NameBuiltin:           "italic " + latteText,
	chroma.NameFunction:          latteSky,
	chroma.NameClass:             latteYellow,
	chroma.NameNamespace:         latteYellow,
	chroma.NameException:         latteMaroon,
	chroma.NameVariable:          lattePeach,
	chroma.NameConstant:          latteYellow,
	chroma.NameLabel:             latteYellow,
	chroma.NameEntity:            lattePink,
	chroma.NameAttribute:         latteYellow,
	chroma.NameTag:               latteMauve,
	chroma.NameDecorator:         lattePink,
	chroma.NameOther:             lattePeach,
	chroma.Punctuation:           latteText,
	chroma.LiteralString:         latteGreen,
	chroma.LiteralStringDoc:      latteGreen,
	chroma.LiteralStringInterpol: latteGreen,
	chroma.LiteralStringEscape:   latteBlue,
	chroma.LiteralStringRegex:    latteBlue,
	chroma.LiteralStringSymbol:   latteGreen,
	chroma.LiteralStringOther:    latteGreen,
	chroma.LiteralNumber:         latteTeal,
	chroma.GenericHeading:        "bold " + latteSky,
	chroma.GenericSubheading:     "bold " + latteSky,
	chroma.GenericDeleted:        latteMaroon,
	chroma.GenericInserted:       latteGreen,
	chroma.GenericError:          latteMaroon,
	chroma.GenericEmph:           "italic",
	chroma.GenericStrong:         "bold",
	chroma.GenericPrompt:         "bold " + latteOverlay0,
	chroma.GenericOutput:         lattePeach,
	chroma.GenericTraceback:      latteMaroon,
	chroma.Error:                 latteRed,
	chroma.Background:            lattePeach + " bg:" + latteBase,
}))
