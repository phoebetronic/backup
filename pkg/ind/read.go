package ind

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Read(p string) Index {
	var err error

	var cur []byte
	{
		cur, err = ioutil.ReadFile(filepath.Join(p, "ind", "ind.json"))
		if os.IsNotExist(err) {
			cur = []byte(string("{}"))
		} else if err != nil {
			panic(err)
		}
	}

	var ind Index
	{
		err = json.Unmarshal(cur, &ind)
		if err != nil {
			panic(err)
		}
	}

	return ind
}
