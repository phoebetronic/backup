package main

import (
	"math/rand"
	"time"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"

	"github.com/phoebetron/backup/cmd"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	{
		err := mainE()
		if err != nil {
			tracer.Panic(err)
		}
	}
}

func mainE() error {
	var err error

	var c *cobra.Command
	{
		c, err = cmd.New()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = c.Execute()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
