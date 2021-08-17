package PublicProfileService

func NewClient(ahost, aport string) Client {

	c := &pubprofileclient{
		host: ahost,
		port: aport,
	}

	return c

}
