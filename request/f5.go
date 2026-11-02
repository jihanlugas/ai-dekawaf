package request

type F5Login struct {
	Username          string `json:"username"`
	Password          string `json:"password"`
	LoginProviderName string `json:"loginProviderName"`
}

type F5CreateNode struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type F5CreatePool struct {
	Name    string `json:"name"`
	Members []struct {
		Name string `json:"name"`
	} `json:"members"`
}
