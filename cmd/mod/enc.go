package mod

import "github.com/phoebetron/series/buck"

func (r *run) enc(lis []string) []string {
	return append([]string{buck.Inds(lis[0])}, lis[1:]...)
}
