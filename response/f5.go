package response

type F5Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	// Referer         string      `json:"referer"`
	// RestOperationID int         `json:"restOperationId"`
	// Kind            string      `json:"kind"`
	// ErrorStack      interface{} `json:"errorStack"`
	// ApiError        int         `json:"apiError"`
}

type F5Login struct {
	Username       string `json:"username"`
	LoginReference struct {
		Link string `json:"link"`
	} `json:"loginReference"`
	LoginProviderName string `json:"loginProviderName"`
	Token             struct {
		Token            string `json:"token"`
		Name             string `json:"name"`
		UserName         string `json:"userName"`
		AuthProviderName string `json:"authProviderName"`
		User             struct {
			Link string `json:"link"`
		} `json:"user"`
		Timeout          int    `json:"timeout"`
		StartTime        string `json:"startTime"`
		Address          string `json:"address"`
		Partition        string `json:"partition"`
		Generation       int    `json:"generation"`
		LastUpdateMicros int64  `json:"lastUpdateMicros"`
		ExpirationMicros int64  `json:"expirationMicros"`
		Kind             string `json:"kind"`
		SelfLink         string `json:"selfLink"`
	} `json:"token"`
	Generation       int `json:"generation"`
	LastUpdateMicros int `json:"lastUpdateMicros"`
}

type F5CreateNode struct {
	Kind            string `json:"kind"`
	Name            string `json:"name"`
	Partition       string `json:"partition"`
	FullPath        string `json:"fullPath"`
	Generation      int    `json:"generation"`
	SelfLink        string `json:"selfLink"`
	Address         string `json:"address"`
	ConnectionLimit int    `json:"connectionLimit"`
	DynamicRatio    int    `json:"dynamicRatio"`
	Ephemeral       string `json:"ephemeral"`
	Fqdn            struct {
		AddressFamily string `json:"addressFamily"`
		Autopopulate  string `json:"autopopulate"`
		DownInterval  int    `json:"downInterval"`
		Interval      string `json:"interval"`
	} `json:"fqdn"`
	Logging   string `json:"logging"`
	Monitor   string `json:"monitor"`
	RateLimit string `json:"rateLimit"`
	Ratio     int    `json:"ratio"`
	Session   string `json:"session"`
	State     string `json:"state"`
}

type F5CreatePool struct {
	Kind                   string `json:"kind"`
	Name                   string `json:"name"`
	Partition              string `json:"partition"`
	FullPath               string `json:"fullPath"`
	Generation             int    `json:"generation"`
	SelfLink               string `json:"selfLink"`
	AllowNat               string `json:"allowNat"`
	AllowSnat              string `json:"allowSnat"`
	IgnorePersistedWeight  string `json:"ignorePersistedWeight"`
	IPTosToClient          string `json:"ipTosToClient"`
	IPTosToServer          string `json:"ipTosToServer"`
	LinkQosToClient        string `json:"linkQosToClient"`
	LinkQosToServer        string `json:"linkQosToServer"`
	LoadBalancingMode      string `json:"loadBalancingMode"`
	MinActiveMembers       int    `json:"minActiveMembers"`
	MinUpMembers           int    `json:"minUpMembers"`
	MinUpMembersAction     string `json:"minUpMembersAction"`
	MinUpMembersChecking   string `json:"minUpMembersChecking"`
	QueueDepthLimit        int    `json:"queueDepthLimit"`
	QueueOnConnectionLimit string `json:"queueOnConnectionLimit"`
	QueueTimeLimit         int    `json:"queueTimeLimit"`
	ReselectTries          int    `json:"reselectTries"`
	ServiceDownAction      string `json:"serviceDownAction"`
	SlowRampTime           int    `json:"slowRampTime"`
	MembersReference       struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"membersReference"`
}

type F5Upload struct {
	RemainingByteCount int `json:"remainingByteCount"`
	UsedChunks         struct {
		Num0 int `json:"0"`
	} `json:"usedChunks"`
	TotalByteCount    int    `json:"totalByteCount"`
	LocalFilePath     string `json:"localFilePath"`
	TemporaryFilePath string `json:"temporaryFilePath"`
	Generation        int    `json:"generation"`
	LastUpdateMicros  int64  `json:"lastUpdateMicros"`
}

type F5Install struct {
	Kind          string `json:"kind"`
	Command       string `json:"command"`
	Name          string `json:"name"`
	FromLocalFile string `json:"from-local-file"`
}

type F5CertAndChain struct {
	Kind         string `json:"kind"`
	Name         string `json:"name"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	SelfLink     string `json:"selfLink"`
	APIRawValues struct {
		CertificateKeySize string `json:"certificateKeySize"`
		Expiration         string `json:"expiration"`
		Issuer             string `json:"issuer"`
		PublicKeyType      string `json:"publicKeyType"`
	} `json:"apiRawValues"`
	City                    string `json:"city,omitempty"`
	CommonName              string `json:"commonName"`
	Country                 string `json:"country,omitempty"`
	Fingerprint             string `json:"fingerprint"`
	Organization            string `json:"organization,omitempty"`
	Ou                      string `json:"ou,omitempty"`
	State                   string `json:"state,omitempty"`
	CertValidatorsReference struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"certValidatorsReference"`
	EmailAddress           string `json:"emailAddress,omitempty"`
	SubjectAlternativeName string `json:"subjectAlternativeName,omitempty"`
}

type F5Key struct {
	Kind         string `json:"kind"`
	Name         string `json:"name"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	SelfLink     string `json:"selfLink"`
	KeySize      string `json:"keySize"`
	KeyType      string `json:"keyType"`
	SecurityType string `json:"securityType"`
}

type F5CreateSslSecurityProfile struct {
	Kind                     string `json:"kind"`
	Name                     string `json:"name"`
	Partition                string `json:"partition"`
	FullPath                 string `json:"fullPath"`
	Generation               int    `json:"generation"`
	SelfLink                 string `json:"selfLink"`
	AlertTimeout             string `json:"alertTimeout"`
	AllowDynamicRecordSizing string `json:"allowDynamicRecordSizing"`
	AllowExpiredCrl          string `json:"allowExpiredCrl"`
	AllowNonSsl              string `json:"allowNonSsl"`
	AppService               string `json:"appService"`
	Authenticate             string `json:"authenticate"`
	AuthenticateDepth        int    `json:"authenticateDepth"`
	BypassOnClientCertFail   string `json:"bypassOnClientCertFail"`
	BypassOnHandshakeAlert   string `json:"bypassOnHandshakeAlert"`
	C3DClientFallbackCert    string `json:"c3dClientFallbackCert"`
	C3DDropUnknownOcspStatus string `json:"c3dDropUnknownOcspStatus"`
	C3DOcsp                  string `json:"c3dOcsp"`
	CaFile                   string `json:"caFile"`
	CacheSize                int    `json:"cacheSize"`
	CacheTimeout             int    `json:"cacheTimeout"`
	Cert                     string `json:"cert"`
	CertReference            struct {
		Link string `json:"link"`
	} `json:"certReference"`
	CertExtensionIncludes  []string `json:"certExtensionIncludes"`
	CertLifespan           int      `json:"certLifespan"`
	CertLookupByIpaddrPort string   `json:"certLookupByIpaddrPort"`
	Chain                  string   `json:"chain"`
	ChainReference         struct {
		Link string `json:"link"`
	} `json:"chainReference"`
	CipherGroup           string `json:"cipherGroup"`
	Ciphers               string `json:"ciphers"`
	ClientCertCa          string `json:"clientCertCa"`
	Crl                   string `json:"crl"`
	CrlFile               string `json:"crlFile"`
	Data0Rtt              string `json:"data_0rtt"`
	DefaultsFrom          string `json:"defaultsFrom"`
	DefaultsFromReference struct {
		Link string `json:"link"`
	} `json:"defaultsFromReference"`
	Description                     string        `json:"description"`
	DestinationIPBlacklist          string        `json:"destinationIpBlacklist"`
	DestinationIPWhitelist          string        `json:"destinationIpWhitelist"`
	ForwardProxyBypassDefaultAction string        `json:"forwardProxyBypassDefaultAction"`
	GenericAlert                    string        `json:"genericAlert"`
	HandshakeTimeout                string        `json:"handshakeTimeout"`
	HelloExtensionIncludes          []interface{} `json:"helloExtensionIncludes"`
	HostnameBlacklist               string        `json:"hostnameBlacklist"`
	HostnameWhitelist               string        `json:"hostnameWhitelist"`
	InheritCaCertkeychain           string        `json:"inheritCaCertkeychain"`
	InheritCertkeychain             string        `json:"inheritCertkeychain"`
	Key                             string        `json:"key"`
	KeyReference                    struct {
		Link string `json:"link"`
	} `json:"keyReference"`
	LogPublisher          string `json:"logPublisher"`
	LogPublisherReference struct {
		Link string `json:"link"`
	} `json:"logPublisherReference"`
	LogSslC3DEvents                    string `json:"logSslC3dEvents"`
	LogSslClientAuthenticationEvents   string `json:"logSslClientAuthenticationEvents"`
	LogSslForwardProxyEvents           string `json:"logSslForwardProxyEvents"`
	LogSslHandshakeEvents              string `json:"logSslHandshakeEvents"`
	MaxActiveHandshakes                string `json:"maxActiveHandshakes"`
	MaxAggregateRenegotiationPerMinute string `json:"maxAggregateRenegotiationPerMinute"`
	MaxRenegotiationsPerMinute         int    `json:"maxRenegotiationsPerMinute"`
	MaximumRecordSize                  int    `json:"maximumRecordSize"`
	ModSslMethods                      string `json:"modSslMethods"`
	Mode                               string `json:"mode"`
	NotifyCertStatusToVirtualServer    string `json:"notifyCertStatusToVirtualServer"`
	OcspStapling                       string `json:"ocspStapling"`
	TmOptions                          string `json:"tmOptions"`
	PeerCertMode                       string `json:"peerCertMode"`
	PeerNoRenegotiateTimeout           string `json:"peerNoRenegotiateTimeout"`
	ProxyCaCert                        string `json:"proxyCaCert"`
	ProxyCaKey                         string `json:"proxyCaKey"`
	ProxySsl                           string `json:"proxySsl"`
	ProxySslPassthrough                string `json:"proxySslPassthrough"`
	RenegotiateMaxRecordDelay          string `json:"renegotiateMaxRecordDelay"`
	RenegotiatePeriod                  string `json:"renegotiatePeriod"`
	RenegotiateSize                    string `json:"renegotiateSize"`
	Renegotiation                      string `json:"renegotiation"`
	RetainCertificate                  string `json:"retainCertificate"`
	SecureRenegotiation                string `json:"secureRenegotiation"`
	ServerName                         string `json:"serverName"`
	SessionMirroring                   string `json:"sessionMirroring"`
	SessionTicket                      string `json:"sessionTicket"`
	SessionTicketTimeout               int    `json:"sessionTicketTimeout"`
	SniDefault                         string `json:"sniDefault"`
	SniRequire                         string `json:"sniRequire"`
	SourceIPBlacklist                  string `json:"sourceIpBlacklist"`
	SourceIPWhitelist                  string `json:"sourceIpWhitelist"`
	SslC3D                             string `json:"sslC3d"`
	SslForwardProxy                    string `json:"sslForwardProxy"`
	SslForwardProxyBypass              string `json:"sslForwardProxyBypass"`
	SslForwardProxyVerifiedHandshake   string `json:"sslForwardProxyVerifiedHandshake"`
	SslSignHash                        string `json:"sslSignHash"`
	StrictResume                       string `json:"strictResume"`
	UncleanShutdown                    string `json:"uncleanShutdown"`
	CertKeyChain                       []struct {
		Name          string `json:"name"`
		AppService    string `json:"appService"`
		Cert          string `json:"cert"`
		CertReference struct {
			Link string `json:"link"`
		} `json:"certReference"`
		Chain          string `json:"chain"`
		ChainReference struct {
			Link string `json:"link"`
		} `json:"chainReference"`
		Key          string `json:"key"`
		KeyReference struct {
			Link string `json:"link"`
		} `json:"keyReference"`
		Usage string `json:"usage"`
	} `json:"certKeyChain"`
}

type F5ActiveSNI struct {
	Kind                     string `json:"kind"`
	Name                     string `json:"name"`
	Partition                string `json:"partition"`
	FullPath                 string `json:"fullPath"`
	Generation               int    `json:"generation"`
	SelfLink                 string `json:"selfLink"`
	AlertTimeout             string `json:"alertTimeout"`
	AllowDynamicRecordSizing string `json:"allowDynamicRecordSizing"`
	AllowExpiredCrl          string `json:"allowExpiredCrl"`
	AllowNonSsl              string `json:"allowNonSsl"`
	AppService               string `json:"appService"`
	Authenticate             string `json:"authenticate"`
	AuthenticateDepth        int    `json:"authenticateDepth"`
	BypassOnClientCertFail   string `json:"bypassOnClientCertFail"`
	BypassOnHandshakeAlert   string `json:"bypassOnHandshakeAlert"`
	C3DClientFallbackCert    string `json:"c3dClientFallbackCert"`
	C3DDropUnknownOcspStatus string `json:"c3dDropUnknownOcspStatus"`
	C3DOcsp                  string `json:"c3dOcsp"`
	CaFile                   string `json:"caFile"`
	CacheSize                int    `json:"cacheSize"`
	CacheTimeout             int    `json:"cacheTimeout"`
	Cert                     string `json:"cert"`
	CertReference            struct {
		Link string `json:"link"`
	} `json:"certReference"`
	CertExtensionIncludes  []string `json:"certExtensionIncludes"`
	CertLifespan           int      `json:"certLifespan"`
	CertLookupByIpaddrPort string   `json:"certLookupByIpaddrPort"`
	Chain                  string   `json:"chain"`
	ChainReference         struct {
		Link string `json:"link"`
	} `json:"chainReference"`
	CipherGroup           string `json:"cipherGroup"`
	Ciphers               string `json:"ciphers"`
	ClientCertCa          string `json:"clientCertCa"`
	Crl                   string `json:"crl"`
	CrlFile               string `json:"crlFile"`
	Data0Rtt              string `json:"data_0rtt"`
	DefaultsFrom          string `json:"defaultsFrom"`
	DefaultsFromReference struct {
		Link string `json:"link"`
	} `json:"defaultsFromReference"`
	Description                     string        `json:"description"`
	DestinationIPBlacklist          string        `json:"destinationIpBlacklist"`
	DestinationIPWhitelist          string        `json:"destinationIpWhitelist"`
	ForwardProxyBypassDefaultAction string        `json:"forwardProxyBypassDefaultAction"`
	GenericAlert                    string        `json:"genericAlert"`
	HandshakeTimeout                string        `json:"handshakeTimeout"`
	HelloExtensionIncludes          []interface{} `json:"helloExtensionIncludes"`
	HostnameBlacklist               string        `json:"hostnameBlacklist"`
	HostnameWhitelist               string        `json:"hostnameWhitelist"`
	InheritCaCertkeychain           string        `json:"inheritCaCertkeychain"`
	InheritCertkeychain             string        `json:"inheritCertkeychain"`
	Key                             string        `json:"key"`
	KeyReference                    struct {
		Link string `json:"link"`
	} `json:"keyReference"`
	LogPublisher          string `json:"logPublisher"`
	LogPublisherReference struct {
		Link string `json:"link"`
	} `json:"logPublisherReference"`
	LogSslC3DEvents                    string `json:"logSslC3dEvents"`
	LogSslClientAuthenticationEvents   string `json:"logSslClientAuthenticationEvents"`
	LogSslForwardProxyEvents           string `json:"logSslForwardProxyEvents"`
	LogSslHandshakeEvents              string `json:"logSslHandshakeEvents"`
	MaxActiveHandshakes                string `json:"maxActiveHandshakes"`
	MaxAggregateRenegotiationPerMinute string `json:"maxAggregateRenegotiationPerMinute"`
	MaxRenegotiationsPerMinute         int    `json:"maxRenegotiationsPerMinute"`
	MaximumRecordSize                  int    `json:"maximumRecordSize"`
	ModSslMethods                      string `json:"modSslMethods"`
	Mode                               string `json:"mode"`
	NotifyCertStatusToVirtualServer    string `json:"notifyCertStatusToVirtualServer"`
	OcspStapling                       string `json:"ocspStapling"`
	TmOptions                          string `json:"tmOptions"`
	PeerCertMode                       string `json:"peerCertMode"`
	PeerNoRenegotiateTimeout           string `json:"peerNoRenegotiateTimeout"`
	ProxyCaCert                        string `json:"proxyCaCert"`
	ProxyCaKey                         string `json:"proxyCaKey"`
	ProxySsl                           string `json:"proxySsl"`
	ProxySslPassthrough                string `json:"proxySslPassthrough"`
	RenegotiateMaxRecordDelay          string `json:"renegotiateMaxRecordDelay"`
	RenegotiatePeriod                  string `json:"renegotiatePeriod"`
	RenegotiateSize                    string `json:"renegotiateSize"`
	Renegotiation                      string `json:"renegotiation"`
	RetainCertificate                  string `json:"retainCertificate"`
	SecureRenegotiation                string `json:"secureRenegotiation"`
	ServerName                         string `json:"serverName"`
	SessionMirroring                   string `json:"sessionMirroring"`
	SessionTicket                      string `json:"sessionTicket"`
	SessionTicketTimeout               int    `json:"sessionTicketTimeout"`
	SniDefault                         string `json:"sniDefault"`
	SniRequire                         string `json:"sniRequire"`
	SourceIPBlacklist                  string `json:"sourceIpBlacklist"`
	SourceIPWhitelist                  string `json:"sourceIpWhitelist"`
	SslC3D                             string `json:"sslC3d"`
	SslForwardProxy                    string `json:"sslForwardProxy"`
	SslForwardProxyBypass              string `json:"sslForwardProxyBypass"`
	SslForwardProxyVerifiedHandshake   string `json:"sslForwardProxyVerifiedHandshake"`
	SslSignHash                        string `json:"sslSignHash"`
	StrictResume                       string `json:"strictResume"`
	UncleanShutdown                    string `json:"uncleanShutdown"`
	CertKeyChain                       []struct {
		Name          string `json:"name"`
		AppService    string `json:"appService"`
		Cert          string `json:"cert"`
		CertReference struct {
			Link string `json:"link"`
		} `json:"certReference"`
		Chain          string `json:"chain"`
		ChainReference struct {
			Link string `json:"link"`
		} `json:"chainReference"`
		Key          string `json:"key"`
		KeyReference struct {
			Link string `json:"link"`
		} `json:"keyReference"`
		Usage string `json:"usage"`
	} `json:"certKeyChain"`
}

type F5SNI struct {
	Kind                     string `json:"kind"`
	Name                     string `json:"name"`
	Partition                string `json:"partition"`
	FullPath                 string `json:"fullPath"`
	Generation               int    `json:"generation"`
	SelfLink                 string `json:"selfLink"`
	AlertTimeout             string `json:"alertTimeout"`
	AllowDynamicRecordSizing string `json:"allowDynamicRecordSizing"`
	AllowExpiredCrl          string `json:"allowExpiredCrl"`
	AllowNonSsl              string `json:"allowNonSsl"`
	AppService               string `json:"appService"`
	Authenticate             string `json:"authenticate"`
	AuthenticateDepth        int    `json:"authenticateDepth"`
	BypassOnClientCertFail   string `json:"bypassOnClientCertFail"`
	BypassOnHandshakeAlert   string `json:"bypassOnHandshakeAlert"`
	C3DClientFallbackCert    string `json:"c3dClientFallbackCert"`
	C3DDropUnknownOcspStatus string `json:"c3dDropUnknownOcspStatus"`
	C3DOcsp                  string `json:"c3dOcsp"`
	CaFile                   string `json:"caFile"`
	CacheSize                int    `json:"cacheSize"`
	CacheTimeout             int    `json:"cacheTimeout"`
	Cert                     string `json:"cert"`
	CertReference            struct {
		Link string `json:"link"`
	} `json:"certReference"`
	CertExtensionIncludes  []string `json:"certExtensionIncludes"`
	CertLifespan           int      `json:"certLifespan"`
	CertLookupByIpaddrPort string   `json:"certLookupByIpaddrPort"`
	Chain                  string   `json:"chain"`
	ChainReference         struct {
		Link string `json:"link"`
	} `json:"chainReference"`
	CipherGroup           string `json:"cipherGroup"`
	Ciphers               string `json:"ciphers"`
	ClientCertCa          string `json:"clientCertCa"`
	Crl                   string `json:"crl"`
	CrlFile               string `json:"crlFile"`
	Data0Rtt              string `json:"data_0rtt"`
	DefaultsFrom          string `json:"defaultsFrom"`
	DefaultsFromReference struct {
		Link string `json:"link"`
	} `json:"defaultsFromReference"`
	Description                     string        `json:"description"`
	DestinationIPBlacklist          string        `json:"destinationIpBlacklist"`
	DestinationIPWhitelist          string        `json:"destinationIpWhitelist"`
	ForwardProxyBypassDefaultAction string        `json:"forwardProxyBypassDefaultAction"`
	GenericAlert                    string        `json:"genericAlert"`
	HandshakeTimeout                string        `json:"handshakeTimeout"`
	HelloExtensionIncludes          []interface{} `json:"helloExtensionIncludes"`
	HostnameBlacklist               string        `json:"hostnameBlacklist"`
	HostnameWhitelist               string        `json:"hostnameWhitelist"`
	InheritCaCertkeychain           string        `json:"inheritCaCertkeychain"`
	InheritCertkeychain             string        `json:"inheritCertkeychain"`
	Key                             string        `json:"key"`
	KeyReference                    struct {
		Link string `json:"link"`
	} `json:"keyReference"`
	LogPublisher          string `json:"logPublisher"`
	LogPublisherReference struct {
		Link string `json:"link"`
	} `json:"logPublisherReference"`
	LogSslC3DEvents                    string `json:"logSslC3dEvents"`
	LogSslClientAuthenticationEvents   string `json:"logSslClientAuthenticationEvents"`
	LogSslForwardProxyEvents           string `json:"logSslForwardProxyEvents"`
	LogSslHandshakeEvents              string `json:"logSslHandshakeEvents"`
	MaxActiveHandshakes                string `json:"maxActiveHandshakes"`
	MaxAggregateRenegotiationPerMinute string `json:"maxAggregateRenegotiationPerMinute"`
	MaxRenegotiationsPerMinute         int    `json:"maxRenegotiationsPerMinute"`
	MaximumRecordSize                  int    `json:"maximumRecordSize"`
	ModSslMethods                      string `json:"modSslMethods"`
	Mode                               string `json:"mode"`
	NotifyCertStatusToVirtualServer    string `json:"notifyCertStatusToVirtualServer"`
	OcspStapling                       string `json:"ocspStapling"`
	TmOptions                          string `json:"tmOptions"`
	PeerCertMode                       string `json:"peerCertMode"`
	PeerNoRenegotiateTimeout           string `json:"peerNoRenegotiateTimeout"`
	ProxyCaCert                        string `json:"proxyCaCert"`
	ProxyCaKey                         string `json:"proxyCaKey"`
	ProxySsl                           string `json:"proxySsl"`
	ProxySslPassthrough                string `json:"proxySslPassthrough"`
	RenegotiateMaxRecordDelay          string `json:"renegotiateMaxRecordDelay"`
	RenegotiatePeriod                  string `json:"renegotiatePeriod"`
	RenegotiateSize                    string `json:"renegotiateSize"`
	Renegotiation                      string `json:"renegotiation"`
	RetainCertificate                  string `json:"retainCertificate"`
	SecureRenegotiation                string `json:"secureRenegotiation"`
	ServerName                         string `json:"serverName"`
	SessionMirroring                   string `json:"sessionMirroring"`
	SessionTicket                      string `json:"sessionTicket"`
	SessionTicketTimeout               int    `json:"sessionTicketTimeout"`
	SniDefault                         string `json:"sniDefault"`
	SniRequire                         string `json:"sniRequire"`
	SourceIPBlacklist                  string `json:"sourceIpBlacklist"`
	SourceIPWhitelist                  string `json:"sourceIpWhitelist"`
	SslC3D                             string `json:"sslC3d"`
	SslForwardProxy                    string `json:"sslForwardProxy"`
	SslForwardProxyBypass              string `json:"sslForwardProxyBypass"`
	SslForwardProxyVerifiedHandshake   string `json:"sslForwardProxyVerifiedHandshake"`
	SslSignHash                        string `json:"sslSignHash"`
	StrictResume                       string `json:"strictResume"`
	UncleanShutdown                    string `json:"uncleanShutdown"`
	CertKeyChain                       []struct {
		Name          string `json:"name"`
		AppService    string `json:"appService"`
		Cert          string `json:"cert"`
		CertReference struct {
			Link string `json:"link"`
		} `json:"certReference"`
		Chain          string `json:"chain"`
		ChainReference struct {
			Link string `json:"link"`
		} `json:"chainReference"`
		Key          string `json:"key"`
		KeyReference struct {
			Link string `json:"link"`
		} `json:"keyReference"`
		Usage string `json:"usage"`
	} `json:"certKeyChain"`
}
