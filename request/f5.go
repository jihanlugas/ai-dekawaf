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

type F5CreateSecurityPolicy struct {
	Name                 string `json:"name"`
	ApplicationLanguage  string `json:"applicationLanguage"`
	EnforcementMode      string `json:"enforcementMode"`
	CaseInsensitive      bool   `json:"caseInsensitive"`
	ProtocolIndependent  bool   `json:"protocolIndependent"`
	Active               bool   `json:"active"`
	PolicyBuilderEnabled bool   `json:"policyBuilderEnabled"`
}

type F5AssignSecurityPolicyToVS struct {
	VirtualServers []string `json:"virtualServers"`
}

type F5Link struct {
	Link string `json:"link"`
}

type F5AssignServerTecnologyToSecurityPolicy struct {
	ServerTechnologyReference F5Link `json:"serverTechnologyReference"`
}

type F5SetLearningModePolicy struct {
	LearningMode string `json:"learningMode"`
}

type F5AttachBotDefenseandLogProfileToVS struct {
	SecurityLogProfiles []string `json:"securityLogProfiles"`
	BotDefenseProfile   F5Link   `json:"botDefenseProfile"`
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

type F5CreateVS struct {
	Name        string        `json:"name"`
	Partition   string        `json:"partition"`
	Destination string        `json:"destination"`
	Mask        string        `json:"mask"`
	IPProtocol  string        `json:"ipProtocol"`
	Pool        string        `json:"pool"`
	Profiles    []F5VSProfile `json:"profiles"`
}

type F5VSProfile struct {
	Name    string `json:"name"`
	Context string `json:"context"`
}
