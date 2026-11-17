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

type F5Install struct {
	Command       string `json:"command"`
	Name          string `json:"name"`
	FromLocalFile string `json:"from-local-file"`
}

type F5CreateSslSecurityProfile struct {
	Name         string `json:"name"`
	DefaultsFrom string `json:"defaultsFrom"`
	CertKeyChain []struct {
		Name  string `json:"name"`
		Cert  string `json:"cert"`
		Key   string `json:"key"`
		Chain string `json:"chain"`
	} `json:"certKeyChain"`
}

type F5ActiveSNI struct {
	ServerName string `json:"serverName"`
	SniDefault string `json:"sniDefault"`
}
