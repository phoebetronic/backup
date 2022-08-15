package upl

func (r *run) rem() {
	for _, h := range r.frames {
		err := r.storage.Delete(h.Sta)
		if err != nil {
			panic(err)
		}
	}
}
