package apicliaws

func Default() *AWS {
	var err error

	var aws *AWS
	{
		c := Config{}

		aws, err = New(c)
		if err != nil {
			panic(err)
		}
	}

	return aws
}
