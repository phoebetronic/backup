package ens

import "github.com/phoebetron/backup/pkg/ind"

func (r *run) enc(lis []string) []string {
	return append([]string{ind.EncI()[lis[0]]}, lis[1:]...)
}
