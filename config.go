package learning

type PublicPerson struct {
	hiddenName string
	PublicName string
}

type privatePerson struct {
	name string
}

func Config() string {
	return "learning go"
}

func privateConfig() string {
	return "not visible to anyone"
}
