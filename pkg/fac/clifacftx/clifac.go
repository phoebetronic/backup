package clifacftx

import (
	"github.com/go-numb/go-ftx/auth"
	"github.com/go-numb/go-ftx/rest"
)

type FTX struct {
	key string
	sec string
}

func New(con Config) (*FTX, error) {
	{
		con.Verify()
	}

	f := &FTX{
		key: con.Key,
		sec: con.Sec,
	}

	return f, nil
}

func (f *FTX) New(acc ...string) *rest.Client {
	if len(acc) > 1 {
		panic("must not create more than one account at a time")
	}

	var uid int
	{
		uid = 7
	}

	var sub []auth.SubAccount
	{
		if len(acc) == 1 {
			sub = append(sub, auth.SubAccount{
				UUID:     uid,
				Nickname: acc[0],
			})
		}
	}

	var aut *auth.Config
	{
		aut = auth.New(
			f.key,
			f.sec,
			sub...,
		)
	}

	var cli *rest.Client
	{
		cli = rest.New(aut)
	}

	{
		if len(acc) == 1 {
			cli.Auth.UseSubAccountID(uid)
		}
	}

	return cli
}
