package cmd

import "github.com/gookit/gcli/v2"

// ShowDemoCommand create
func ShowDemoCommand() *gcli.Command {
	c := &gcli.Command{
		Name: "show",
		Func: runShow,
		//
		UseFor: "the command will show some data format methods",
	}

	return c
}

func runShow(c *gcli.Command, _ []string) error {

	return nil
}
