package PublicProfileService

func NewClient(ahost, aport string) Client {

	return &pubprofileclient{
		host: ahost,
		port: aport,
	}
}
