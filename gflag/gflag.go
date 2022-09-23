package gflag

import (
	"github.com/gookit/goutil/strutil"
)

const (
	// AlignLeft Align right, padding left
	AlignLeft = strutil.PosRight
	// AlignRight Align left, padding right
	AlignRight = strutil.PosLeft

	// default desc
	defaultDesc = "No description"

	// TagRuleNamed struct tag use named k-v rule.
	//
	// eg: `flag:"name=int0;shorts=i;required=true;desc=int option message"`
	TagRuleNamed = 0

	// TagRuleSimple struct tag use simple rule.
	// format: "desc;required;default;shorts"
	//
	// eg: `flag:"int option message;required;;i"`
	TagRuleSimple = 1
)

// FlagTagName default tag name on struct
var FlagTagName = "flag"

// FlagsConfig for render help information
type FlagsConfig struct {
	// WithoutType don't display flag data type on print help
	WithoutType bool
	// DescNewline flag desc at new line on print help
	DescNewline bool
	// Alignment flag name align left or right. default is: left
	Alignment uint8
	// TagName on struct
	TagName string
	// TagRuleType for struct tag value. default is TagRuleNamed
	TagRuleType uint8
	// DisableArg disable binding arguments.
	DisableArg bool
}

// OptCategory struct
type OptCategory struct {
	Name, Title string
	OptNames    []string
}
