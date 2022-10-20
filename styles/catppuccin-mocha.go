package styles

import "github.com/alecthomas/chroma/v2"

var (
	mochaFlamingo = "#F2CDCD"
	mochaPink     = "#F5C2E7"
	mochaMauve    = "#CBA6F7"
	mochaRed      = "#F38BA8"
	mochaMaroon   = "#EBA0AC"
	mochaPeach    = "#FAB387"
	mochaYellow   = "#F9E2AF"
	mochaGreen    = "#A6E3A1"
	mochaTeal     = "#94E2D5"
	mochaSky      = "#89DCEB"
	mochaBlue     = "#87B0F9"
	mochaLavender = "#B4BEFE"
	mochaText     = "#C6D0F5"
	mochaOverlay0 = "#696D86"
	mochaSurface2 = "#565970"
	mochaSurface0 = "#313244"
	mochaBase     = "#1E1E2E"
)

// Catppuccin a soothing pastel theme for the high-spirited (mocha variant)
var CatppuccinMocha = Register(chroma.MustNewStyle("catppuccin-mocha", chroma.StyleEntries{
	chroma.TextWhitespace:        mochaSurface0,
	chroma.Comment:               "italic " + mochaSurface2,
	chroma.CommentPreproc:        mochaBlue,
	chroma.Keyword:               mochaMauve,
	chroma.KeywordPseudo:         "bold " + mochaMauve,
	chroma.KeywordType:           mochaFlamingo,
	chroma.KeywordConstant:       "italic " + mochaMauve,
	chroma.Operator:              mochaSky,
	chroma.OperatorWord:          "bold " + mochaSky,
	chroma.Name:                  mochaLavender,
	chroma.NameBuiltin:           "italic " + mochaText,
	chroma.NameFunction:          mochaSky,
	chroma.NameClass:             mochaYellow,
	chroma.NameNamespace:         mochaYellow,
	chroma.NameException:         mochaMaroon,
	chroma.NameVariable:          mochaPeach,
	chroma.NameConstant:          mochaYellow,
	chroma.NameLabel:             mochaYellow,
	chroma.NameEntity:            mochaPink,
	chroma.NameAttribute:         mochaYellow,
	chroma.NameTag:               mochaMauve,
	chroma.NameDecorator:         mochaPink,
	chroma.NameOther:             mochaPeach,
	chroma.Punctuation:           mochaText,
	chroma.LiteralString:         mochaGreen,
	chroma.LiteralStringDoc:      mochaGreen,
	chroma.LiteralStringInterpol: mochaGreen,
	chroma.LiteralStringEscape:   mochaBlue,
	chroma.LiteralStringRegex:    mochaBlue,
	chroma.LiteralStringSymbol:   mochaGreen,
	chroma.LiteralStringOther:    mochaGreen,
	chroma.LiteralNumber:         mochaTeal,
	chroma.GenericHeading:        "bold " + mochaSky,
	chroma.GenericSubheading:     "bold " + mochaSky,
	chroma.GenericDeleted:        mochaMaroon,
	chroma.GenericInserted:       mochaGreen,
	chroma.GenericError:          mochaMaroon,
	chroma.GenericEmph:           "italic",
	chroma.GenericStrong:         "bold",
	chroma.GenericPrompt:         "bold " + mochaOverlay0,
	chroma.GenericOutput:         mochaPeach,
	chroma.GenericTraceback:      mochaMaroon,
	chroma.Error:                 mochaRed,
	chroma.Background:            mochaPeach + " bg:" + mochaBase,
}))
