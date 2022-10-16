package events

// constants for hooks event, there are default allowed event names
const (
	// OnAppInitBefore On app init before
	OnAppInitBefore = "app.init.before"
	// OnAppInitAfter On app init after
	OnAppInitAfter = "app.init.after"
	// OnAppStop   = "app.stopped"

	// OnAppBindOptsBefore bind app options
	OnAppBindOptsBefore = "app.bind.opts.before"
	OnAppBindOptsAfter  = "app.bind.opts.after"

	// OnAppCmdAdd on app cmd add
	OnAppCmdAdd = "app.cmd.add.before"

	// OnAppCmdAdded on app cmd added
	OnAppCmdAdded = "app.cmd.added"

	// OnAppOptsParsed event
	//
	// Data:
	// 	{args: app-args}
	OnAppOptsParsed = "app.opts.parsed"

	// OnAppPrepared prepare for run
	OnAppPrepared = "app.run.prepared"

	OnAppRunBefore = "app.run.before"
	OnAppRunAfter  = "app.run.after"
	OnAppRunError  = "app.run.error"

	OnCmdInitBefore = "cmd.init.before"
	OnCmdInitAfter  = "cmd.init.after"

	// OnCmdNotFound on app-command or subcommand not found.
	//
	// Data:
	// 	{name: command-name}
	OnCmdNotFound = "cmd.not.found"
	// OnAppCmdNotFound on app command not found
	OnAppCmdNotFound = "app.cmd.not.found"
	// OnCmdSubNotFound on subcommand not found
	OnCmdSubNotFound = "cmd.sub.not.found"

	// OnCmdOptParsed event
	//
	// Data:
	// 	{args: command-args}
	OnCmdOptParsed = "cmd.opts.parsed"

	// OnCmdRunBefore cmd run
	OnCmdRunBefore = "cmd.run.before"
	OnCmdRunAfter  = "cmd.run.after"
	OnCmdRunError  = "cmd.run.error"

	// OnCmdExecBefore cmd exec
	OnCmdExecBefore = "cmd.exec.before"
	OnCmdExecAfter  = "cmd.exec.after"
	OnCmdExecError  = "cmd.exec.error"

	// OnGlobalOptsParsed event
	//
	// Data:
	// 	{args: remain-args}
	OnGlobalOptsParsed = "gcli.gopts.parsed"
)
