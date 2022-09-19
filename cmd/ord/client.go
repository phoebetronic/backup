package ord

import (
	"github.com/phoebetron/backup/pkg/cli/apiclidydx"
)

func (r *run) newcli() Client {
	var cli Client
	switch r.flags.Exchange {
	case "dydx":
		cli = apiclidydx.Default(r.flags.Asset)
	}

	return cli
}
