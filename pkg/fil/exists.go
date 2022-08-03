package fil

import "os"

func Exists(file string) bool {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	} else if err != nil {
		panic(err)
	}

	return true
}
