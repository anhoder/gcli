package cliapp

import (
	"os"
)

// constants for error level 0 - 4
const (
	VerbQuiet uint = iota // don't report anything
	VerbError             // reporting on error
	VerbWarn
	VerbInfo
	VerbDebug
	// VerbCrazy
)

// HelpVar allow var replace in help info. like "{$binName}" "{$cmd}"
const HelpVar = "{$%s}"

// Logo app logo, ASCII logo
type Logo struct {
	Text  string // ASCII logo string
	Style string // eg "info"
}

// GlobalOpts global flags
type GlobalOpts struct {
	noColor  bool
	verbose  uint // message report level
	showVer  bool
	showHelp bool
}

// Application the cli app definition
type Application struct {
	// Name app name
	Name string
	// Version app version. like "1.0.1"
	Version string
	// Description app description
	Description string
	// Logo ASCII logo setting
	Logo Logo
	// Hooks can setting some hooks func on running.
	// allow hooks: "init", "before", "after", "error"
	Hooks map[string]func(app *Application)
	// Strict use strict mode. short opt must be begin '-', long opt must be begin '--'
	Strict bool

	// vars you can add some vars map for render help info
	vars map[string]string
	// command names. key is name, value is name string length
	// eg. {"test": 4, "example": 7}
	names map[string]int
	// command aliases map. {alias: name}
	aliases map[string]string
	// current command name
	commandName string
	// default command name
	defaultCommand string
}

// global options
var gOpts = GlobalOpts{verbose: VerbError}

// bin script name eg "./cliapp"
var binName = os.Args[0]

// the app work dir path
var workDir, _ = os.Getwd()

// Exit program
func Exit(code int) {
	os.Exit(code)
}

// WorkDir get work dir
func WorkDir() string {
	return workDir
}

// BinName get bin script name
func BinName() string {
	return binName
}

// Verbose returns verbose level
func Verbose() uint {
	return gOpts.verbose
}

// NewApp create new app
// settings (name, version, description)
// cli.NewApp("my cli app", "1.0.1", "The is is my cil application")
func NewApp(settings ...string) *Application {
	app = &Application{Name: "My CLI Application", Version: "1.0.0"}
	app.Logo.Style = "info"

	for k, v := range settings {
		switch k {
		case 0:
			app.Name = v
		case 1:
			app.Version = v
		case 2:
			app.Description = v
		}
	}

	// init
	app.Init()
	return app
}

// LogoText
func (app *Application) LogoText(logo string) {
	app.Logo.Text = logo
}

// DebugMode level
func (app *Application) DebugMode() {
	gOpts.verbose = VerbDebug
}

// QuietMode level
func (app *Application) QuietMode() {
	gOpts.verbose = VerbQuiet
}

// SetVerbose level
func (app *Application) SetVerbose(verbose uint) {
	gOpts.verbose = verbose
}

// DefaultCommand set default command name
func (app *Application) DefaultCommand(name string) {
	app.defaultCommand = name
}

// Init
func (app *Application) Init() {
	app.names = make(map[string]int)

	// init some tpl vars
	app.vars = map[string]string{
		"workDir": workDir,
		"binName": binName,
	}
}

// Add add a command
func (app *Application) Add(c *Command, more ...*Command) {
	app.addCommand(c)

	if len(more) > 0 {
		for _, cmd := range more {
			app.addCommand(cmd)
		}
	}
}

func (app *Application) addCommand(c *Command) {
	if c.IsDisabled() {
		Logf(VerbDebug, "command %s has been disabled, skip add", c.Name)
		return
	}

	commands[c.Name] = c
	app.names[c.Name] = len(c.Name)
	Logf(VerbDebug, "add command: %s", c.Name)

	// will call it on input './cliapp command -h'
	c.Flags.Usage = func() {
		c.ShowHelp(true)
	}

	// init
	c.Init()
	// add app vars to cmd
	c.AddVars(app.vars)
	// add aliases for the command
	app.AddAliases(c.Name, c.Aliases)
}

// AddVar get command name
func (app *Application) AddVar(name string, value string) {
	app.vars[name] = value
}

// AddVars add multi tpl vars
func (app *Application) AddVars(vars map[string]string) {
	for n, v := range vars {
		app.AddVar(n, v)
	}
}

// GetVar get a help var by name
func (app *Application) GetVar(name string) string {
	if v, ok := app.vars[name]; ok {
		return v
	}

	return ""
}

// GetVars get all tpl vars
func (app *Application) GetVars(name string, value string) map[string]string {
	return app.vars
}
