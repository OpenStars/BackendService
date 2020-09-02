package PubProfileClient

func NewClient(ahost, aport string) Client {

	return &pubprofileclient{
		host: ahost,
		port: aport,
	}
}
