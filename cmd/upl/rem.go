package upl

func (r *run) rem() {
	for _, h := range r.misfra {
		err := r.stotra.Delete(h.Sta)
		if err != nil {
			panic(err)
		}
	}
}
