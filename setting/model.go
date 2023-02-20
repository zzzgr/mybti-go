package setting

type Config struct {
	Notify Notify `json:"notify"`
	User   []User `json:"user"`
}

type Notify struct {
	PushplusToken string `json:"pushplusToken"`
}

type User struct {
	Name        string `json:"name"`
	AccessToken string `json:"AccessToken"`
	Line        string `json:"line"`
	Station     string `json:"station"`
	Time        string `json:"time"`
}
