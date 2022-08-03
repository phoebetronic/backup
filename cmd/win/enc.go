package win

func (r *run) enc(enc map[string]string, lis []string) []string {
	return append([]string{enc[lis[0]]}, lis[1:]...)
}
