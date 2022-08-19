package tra

import (
	"github.com/phoebetron/backup/pkg/cli/apiclidydx"
	"github.com/phoebetron/backup/pkg/cli/apicliftx"
)

func (r *run) newcli() Client {
	var cli Client
	switch r.flags.Exchange {
	case "dydx":
		cli = apiclidydx.Default()
	case "ftx":
		cli = apicliftx.Default()
	}

	return cli
}
