package ver

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	rep = "https://github.com/phoebetron/backup"
	sha = "n/a"
	ver = "n/a"
)

type run struct{}

func (r *run) run(cmd *cobra.Command, args []string) {
	fmt.Fprintf(os.Stdout, "Git Sha       %s\n", sha)
	fmt.Fprintf(os.Stdout, "Repository    %s\n", rep)
	fmt.Fprintf(os.Stdout, "Version       %s\n", ver)
	fmt.Fprintf(os.Stdout, "Go Version    %s\n", runtime.Version())
	fmt.Fprintf(os.Stdout, "Go Arch       %s\n", runtime.GOARCH)
	fmt.Fprintf(os.Stdout, "Go OS         %s\n", runtime.GOOS)
}
