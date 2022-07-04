package clifacftx

type Config struct {
	Key string
	Sec string
}

func (c Config) Verify() {
	if c.Key == "" {
		panic("Key must not be empty")
	}
	if c.Sec == "" {
		panic("Sec must not be empty")
	}
}
