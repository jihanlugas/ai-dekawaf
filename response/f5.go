package response

import (
	"encoding/json"
	"time"
)

type F5Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	// Referer         string      `json:"referer"`
	// RestOperationID int         `json:"restOperationId"`
	// Kind            string      `json:"kind"`
	// ErrorStack      interface{} `json:"errorStack"`
	// ApiError        int         `json:"apiError"`
}

type F5Reference struct {
	Link            string `json:"link"`
	IsSubcollection bool   `json:"isSubcollection,omitempty"`
}

type F5VS struct {
	Kind              string       `json:"kind"`
	Name              string       `json:"name"`
	Partition         string       `json:"partition"`
	FullPath          string       `json:"fullPath"`
	Generation        int          `json:"generation"`
	SelfLink          string       `json:"selfLink"`
	AddressStatus     string       `json:"addressStatus"`
	AutoDiscovery     string       `json:"autoDiscovery"`
	AutoLasthop       string       `json:"autoLasthop"`
	CmpEnabled        string       `json:"cmpEnabled"`
	ConnectionLimit   int          `json:"connectionLimit"`
	CreationTime      string       `json:"creationTime"`
	Description       string       `json:"description,omitempty"`
	Destination       string       `json:"destination"`
	Enabled           bool         `json:"enabled"`
	EvictionProtected string       `json:"evictionProtected"`
	GtmScore          int          `json:"gtmScore"`
	IPProtocol        string       `json:"ipProtocol"`
	LastModifiedTime  string       `json:"lastModifiedTime"`
	Mask              string       `json:"mask"`
	Mirror            string       `json:"mirror"`
	MobileAppTunnel   string       `json:"mobileAppTunnel"`
	Nat64             string       `json:"nat64"`
	Pool              string       `json:"pool,omitempty"`
	PoolReference     *F5Reference `json:"poolReference,omitempty"`
	RateLimit         string       `json:"rateLimit"`
	RateLimitDstMask  int          `json:"rateLimitDstMask"`
	RateLimitMode     string       `json:"rateLimitMode"`
	RateLimitSrcMask  int          `json:"rateLimitSrcMask"`
	SecurityNatPolicy struct {
		UseDevicePolicy      string `json:"useDevicePolicy"`
		UseRouteDomainPolicy string `json:"useRouteDomainPolicy"`
	} `json:"securityNatPolicy"`
	ServersslUseSni            string `json:"serversslUseSni"`
	ServiceDownImmediateAction string `json:"serviceDownImmediateAction"`
	Source                     string `json:"source"`
	SourceAddressTranslation   struct {
		Type string `json:"type"`
	} `json:"sourceAddressTranslation"`
	SourcePort                   string        `json:"sourcePort"`
	SynCookieStatus              string        `json:"synCookieStatus"`
	ThroughputCapacity           string        `json:"throughputCapacity"`
	TranslateAddress             string        `json:"translateAddress"`
	TranslatePort                string        `json:"translatePort"`
	VlansDisabled                bool          `json:"vlansDisabled"`
	VsIndex                      int           `json:"vsIndex"`
	Rules                        []string      `json:"rules,omitempty"`
	RulesReference               []F5Reference `json:"rulesReference,omitempty"`
	SecurityLogProfiles          []string      `json:"securityLogProfiles,omitempty"`
	SecurityLogProfilesReference []F5Reference `json:"securityLogProfilesReference,omitempty"`
	PoliciesReference            F5Reference   `json:"policiesReference"`
	ProfilesReference            F5Reference   `json:"profilesReference"`
}

type F5ListVS struct {
	Kind     string `json:"kind"`
	SelfLink string `json:"selfLink"`
	Items    []F5VS `json:"items"`
}

type F5VSProfile struct {
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Partition  string `json:"partition"`
	FullPath   string `json:"fullPath"`
	Generation int    `json:"generation"`
	SelfLink   string `json:"selfLink"`
	Context    string `json:"context"`
}

type F5VerifySSLInVS struct {
	Kind     string        `json:"kind"`
	SelfLink string        `json:"selfLink"`
	Items    []F5VSProfile `json:"items"`
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

type F5MemberPool struct {
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
		Autopopulate string `json:"autopopulate"`
	} `json:"fqdn"`
	InheritProfile string `json:"inheritProfile"`
	Logging        string `json:"logging"`
	Monitor        string `json:"monitor"`
	PriorityGroup  int    `json:"priorityGroup"`
	RateLimit      string `json:"rateLimit"`
	Ratio          int    `json:"ratio"`
	Session        string `json:"session"`
	State          string `json:"state"`
}

type F5ListMemberPool struct {
	Kind     string         `json:"kind"`
	SelfLink string         `json:"selfLink"`
	Items    []F5MemberPool `json:"items"`
}

type F5ASMReference struct {
	Link            string `json:"link"`
	IsSubCollection bool   `json:"isSubCollection,omitempty"`
	Title           string `json:"title,omitempty"`
}

type F5CreateSecurityPolicy struct {
	SignatureRequirementReference               F5ASMReference `json:"signatureRequirementReference"`
	PlainTextProfileReference                   F5ASMReference `json:"plainTextProfileReference"`
	IdlFileReference                            F5ASMReference `json:"idlFileReference"`
	EnablePassiveMode                           bool           `json:"enablePassiveMode"`
	BehavioralEnforcementReference              F5ASMReference `json:"behavioralEnforcementReference"`
	DataGuardReference                          F5ASMReference `json:"dataGuardReference"`
	CreatedDatetime                             string         `json:"createdDatetime"`
	DatabaseProtectionReference                 F5ASMReference `json:"databaseProtectionReference"`
	CookieSettingsReference                     F5ASMReference `json:"cookieSettingsReference"`
	CsrfURLReference                            F5ASMReference `json:"csrfUrlReference"`
	VersionLastChange                           string         `json:"versionLastChange"`
	Name                                        string         `json:"name"`
	CaseInsensitive                             bool           `json:"caseInsensitive"`
	HeaderSettingsReference                     F5ASMReference `json:"headerSettingsReference"`
	SectionReference                            F5ASMReference `json:"sectionReference"`
	FlowReference                               F5ASMReference `json:"flowReference"`
	LoginPageReference                          F5ASMReference `json:"loginPageReference"`
	Description                                 string         `json:"description"`
	FullPath                                    string         `json:"fullPath"`
	PolicyBuilderParameterReference             F5ASMReference `json:"policyBuilderParameterReference"`
	HasParent                                   bool           `json:"hasParent"`
	ThreatCampaignReference                     F5ASMReference `json:"threatCampaignReference"`
	Partition                                   string         `json:"partition"`
	ManagedByBewaf                              bool           `json:"managedByBewaf"`
	CsrfProtectionReference                     F5ASMReference `json:"csrfProtectionReference"`
	GraphqlProfileReference                     F5ASMReference `json:"graphqlProfileReference"`
	PolicyAntivirusReference                    F5ASMReference `json:"policyAntivirusReference"`
	Kind                                        string         `json:"kind"`
	PolicyBuilderCookieReference                F5ASMReference `json:"policyBuilderCookieReference"`
	IPIntelligenceReference                     F5ASMReference `json:"ipIntelligenceReference"`
	ProtocolIndependent                         bool           `json:"protocolIndependent"`
	SessionAwarenessSettingsReference           F5ASMReference `json:"sessionAwarenessSettingsReference"`
	PolicyBuilderURLReference                   F5ASMReference `json:"policyBuilderUrlReference"`
	PolicyBuilderServerTechnologiesReference    F5ASMReference `json:"policyBuilderServerTechnologiesReference"`
	PolicyBuilderFiletypeReference              F5ASMReference `json:"policyBuilderFiletypeReference"`
	SignatureSetReference                       F5ASMReference `json:"signatureSetReference"`
	ParameterReference                          F5ASMReference `json:"parameterReference"`
	ApplicationLanguage                         string         `json:"applicationLanguage"`
	EnforcementMode                             string         `json:"enforcementMode"`
	LoginEnforcementReference                   F5ASMReference `json:"loginEnforcementReference"`
	OpenAPIFileReference                        F5ASMReference `json:"openApiFileReference"`
	NavigationParameterReference                F5ASMReference `json:"navigationParameterReference"`
	GwtProfileReference                         F5ASMReference `json:"gwtProfileReference"`
	WebhookReference                            F5ASMReference `json:"webhookReference"`
	WhitelistIPReference                        F5ASMReference `json:"whitelistIpReference"`
	HistoryRevisionReference                    F5ASMReference `json:"historyRevisionReference"`
	PolicyBuilderReference                      F5ASMReference `json:"policyBuilderReference"`
	ResponsePageReference                       F5ASMReference `json:"responsePageReference"`
	VulnerabilityAssessmentReference            F5ASMReference `json:"vulnerabilityAssessmentReference"`
	ServerTechnologyReference                   F5ASMReference `json:"serverTechnologyReference"`
	CookieReference                             F5ASMReference `json:"cookieReference"`
	BlockingSettingReference                    F5ASMReference `json:"blockingSettingReference"`
	HostNameReference                           F5ASMReference `json:"hostNameReference"`
	VersionDeviceName                           string         `json:"versionDeviceName"`
	SelfLink                                    string         `json:"selfLink"`
	ThreatCampaignSettingReference              F5ASMReference `json:"threatCampaignSettingReference"`
	SignatureReference                          F5ASMReference `json:"signatureReference"`
	PolicyBuilderRedirectionProtectionReference F5ASMReference `json:"policyBuilderRedirectionProtectionReference"`
	FiletypeReference                           F5ASMReference `json:"filetypeReference"`
	ID                                          string         `json:"id"`
	ModifierName                                string         `json:"modifierName"`
	VersionDatetime                             string         `json:"versionDatetime"`
	SsrfHostReference                           F5ASMReference `json:"ssrfHostReference"`
	SessionTrackingStatusReference              F5ASMReference `json:"sessionTrackingStatusReference"`
	Active                                      bool           `json:"active"`
	AuditLogReference                           F5ASMReference `json:"auditLogReference"`
	DisallowedGeolocationReference              F5ASMReference `json:"disallowedGeolocationReference"`
	RedirectionProtectionDomainReference        F5ASMReference `json:"redirectionProtectionDomainReference"`
	Type                                        string         `json:"type"`
	SignatureSettingReference                   F5ASMReference `json:"signatureSettingReference"`
	DeceptionResponsePageReference              F5ASMReference `json:"deceptionResponsePageReference"`
	WebsocketURLReference                       F5ASMReference `json:"websocketUrlReference"`
	XMLProfileReference                         F5ASMReference `json:"xmlProfileReference"`
	MethodReference                             F5ASMReference `json:"methodReference"`
	VulnerabilityReference                      F5ASMReference `json:"vulnerabilityReference"`
	RedirectionProtectionReference              F5ASMReference `json:"redirectionProtectionReference"`
	PolicyBuilderSessionsAndLoginsReference     F5ASMReference `json:"policyBuilderSessionsAndLoginsReference"`
	TemplateReference                           F5ASMReference `json:"templateReference"`
	PolicyBuilderHeaderReference                F5ASMReference `json:"policyBuilderHeaderReference"`
	CreatorName                                 string         `json:"creatorName"`
	URLReference                                F5ASMReference `json:"urlReference"`
	HeaderReference                             F5ASMReference `json:"headerReference"`
	ActionItemReference                         F5ASMReference `json:"actionItemReference"`
	MicroserviceReference                       F5ASMReference `json:"microserviceReference"`
	XMLValidationFileReference                  F5ASMReference `json:"xmlValidationFileReference"`
	LastUpdateMicros                            int64          `json:"lastUpdateMicros"`
	JSONProfileReference                        F5ASMReference `json:"jsonProfileReference"`
	BruteForceAttackPreventionReference         F5ASMReference `json:"bruteForceAttackPreventionReference"`
	DisabledActionItemReference                 F5ASMReference `json:"disabledActionItemReference"`
	JSONValidationFileReference                 F5ASMReference `json:"jsonValidationFileReference"`
	ExtractionReference                         F5ASMReference `json:"extractionReference"`
	CharacterSetReference                       F5ASMReference `json:"characterSetReference"`
	SuggestionReference                         F5ASMReference `json:"suggestionReference"`
	DeceptionSettingsReference                  F5ASMReference `json:"deceptionSettingsReference"`
	IsModified                                  bool           `json:"isModified"`
	SensitiveParameterReference                 F5ASMReference `json:"sensitiveParameterReference"`
	GeneralReference                            F5ASMReference `json:"generalReference"`
	VersionPolicyName                           string         `json:"versionPolicyName"`
	PolicyBuilderCentralConfigurationReference  F5ASMReference `json:"policyBuilderCentralConfigurationReference"`
}

type F5ListSecurityPolicy struct {
	Kind       string             `json:"kind"`
	SelfLink   string             `json:"selfLink"`
	TotalItems int                `json:"totalItems"`
	Items      []F5SecurityPolicy `json:"items"`
}

type F5SecurityPolicy struct {
	F5CreateSecurityPolicy
	// List responses may encode this timestamp in scientific notation.
	LastUpdateMicros     json.Number `json:"lastUpdateMicros"`
	VirtualServers       []string    `json:"virtualServers"`
	ManualVirtualServers []string    `json:"manualVirtualServers"`
	SubPath              string      `json:"subPath"`
}

type F5ServerTechnologyReference struct {
	Link                 string `json:"link"`
	ServerTechnologyName string `json:"serverTechnologyName"`
}

type F5ServerTechnology struct {
	ServerTechnologyDisplayName string                        `json:"serverTechnologyDisplayName"`
	ServerTechnologyName        string                        `json:"serverTechnologyName"`
	LogoFileName                string                        `json:"logoFileName"`
	LastUpdateMicros            int64                         `json:"lastUpdateMicros"`
	Description                 string                        `json:"description"`
	Kind                        string                        `json:"kind"`
	SelfLink                    string                        `json:"selfLink"`
	ServerTechnologyReferences  []F5ServerTechnologyReference `json:"serverTechnologyReferences"`
	ID                          string                        `json:"id"`
	IsLearnable                 bool                          `json:"isLearnable"`
	HasDeceptionResponses       bool                          `json:"hasDeceptionResponses"`
}

type F5ListServerTecnology struct {
	Kind       string               `json:"kind"`
	SelfLink   string               `json:"selfLink"`
	TotalItems int                  `json:"totalItems"`
	Items      []F5ServerTechnology `json:"items"`
}

type F5ServerTechnologyInPolicy struct {
	Kind                      string                      `json:"kind"`
	SelfLink                  string                      `json:"selfLink"`
	ID                        string                      `json:"id"`
	LastUpdateMicros          int64                       `json:"lastUpdateMicros"`
	ServerTechnologyReference F5ServerTechnologyReference `json:"serverTechnologyReference"`
}

type F5ListServerTecnologyInPolicy struct {
	Kind       string                       `json:"kind"`
	SelfLink   string                       `json:"selfLink"`
	TotalItems int                          `json:"totalItems"`
	Items      []F5ServerTechnologyInPolicy `json:"items"`
}

type F5AssignSecurityPolicyToVS struct {
	SignatureRequirementReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"signatureRequirementReference"`
	PlainTextProfileReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"plainTextProfileReference"`
	IdlFileReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"idlFileReference"`
	EnablePassiveMode              bool `json:"enablePassiveMode"`
	BehavioralEnforcementReference struct {
		Link string `json:"link"`
	} `json:"behavioralEnforcementReference"`
	DataGuardReference struct {
		Link string `json:"link"`
	} `json:"dataGuardReference"`
	CreatedDatetime             time.Time `json:"createdDatetime"`
	DatabaseProtectionReference struct {
		Link string `json:"link"`
	} `json:"databaseProtectionReference"`
	CookieSettingsReference struct {
		Link string `json:"link"`
	} `json:"cookieSettingsReference"`
	CsrfURLReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"csrfUrlReference"`
	VersionLastChange       string `json:"versionLastChange"`
	Name                    string `json:"name"`
	CaseInsensitive         bool   `json:"caseInsensitive"`
	HeaderSettingsReference struct {
		Link string `json:"link"`
	} `json:"headerSettingsReference"`
	SectionReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"sectionReference"`
	FlowReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"flowReference"`
	LoginPageReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"loginPageReference"`
	Description                     string `json:"description"`
	FullPath                        string `json:"fullPath"`
	PolicyBuilderParameterReference struct {
		Link string `json:"link"`
	} `json:"policyBuilderParameterReference"`
	HasParent               bool `json:"hasParent"`
	ThreatCampaignReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"threatCampaignReference"`
	Partition               string `json:"partition"`
	ManagedByBewaf          bool   `json:"managedByBewaf"`
	CsrfProtectionReference struct {
		Link string `json:"link"`
	} `json:"csrfProtectionReference"`
	GraphqlProfileReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"graphqlProfileReference"`
	PolicyAntivirusReference struct {
		Link string `json:"link"`
	} `json:"policyAntivirusReference"`
	Kind                         string   `json:"kind"`
	VirtualServers               []string `json:"virtualServers"`
	PolicyBuilderCookieReference struct {
		Link string `json:"link"`
	} `json:"policyBuilderCookieReference"`
	IPIntelligenceReference struct {
		Link string `json:"link"`
	} `json:"ipIntelligenceReference"`
	ProtocolIndependent               bool `json:"protocolIndependent"`
	SessionAwarenessSettingsReference struct {
		Link string `json:"link"`
	} `json:"sessionAwarenessSettingsReference"`
	PolicyBuilderURLReference struct {
		Link string `json:"link"`
	} `json:"policyBuilderUrlReference"`
	PolicyBuilderServerTechnologiesReference struct {
		Link string `json:"link"`
	} `json:"policyBuilderServerTechnologiesReference"`
	PolicyBuilderFiletypeReference struct {
		Link string `json:"link"`
	} `json:"policyBuilderFiletypeReference"`
	SignatureSetReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"signatureSetReference"`
	ParameterReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"parameterReference"`
	ApplicationLanguage       string `json:"applicationLanguage"`
	EnforcementMode           string `json:"enforcementMode"`
	LoginEnforcementReference struct {
		Link string `json:"link"`
	} `json:"loginEnforcementReference"`
	OpenAPIFileReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"openApiFileReference"`
	NavigationParameterReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"navigationParameterReference"`
	ApplicationService  string `json:"applicationService"`
	GwtProfileReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"gwtProfileReference"`
	WebhookReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"webhookReference"`
	WhitelistIPReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"whitelistIpReference"`
	HistoryRevisionReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"historyRevisionReference"`
	PolicyBuilderReference struct {
		Link string `json:"link"`
	} `json:"policyBuilderReference"`
	ResponsePageReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"responsePageReference"`
	VulnerabilityAssessmentReference struct {
		Link string `json:"link"`
	} `json:"vulnerabilityAssessmentReference"`
	ServerTechnologyReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"serverTechnologyReference"`
	CookieReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"cookieReference"`
	BlockingSettingReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"blockingSettingReference"`
	HostNameReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"hostNameReference"`
	VersionDeviceName              string `json:"versionDeviceName"`
	SelfLink                       string `json:"selfLink"`
	ThreatCampaignSettingReference struct {
		Link string `json:"link"`
	} `json:"threatCampaignSettingReference"`
	SignatureReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"signatureReference"`
	PolicyBuilderRedirectionProtectionReference struct {
		Link string `json:"link"`
	} `json:"policyBuilderRedirectionProtectionReference"`
	FiletypeReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"filetypeReference"`
	ID                   string        `json:"id"`
	ModifierName         string        `json:"modifierName"`
	ManualVirtualServers []interface{} `json:"manualVirtualServers"`
	VersionDatetime      time.Time     `json:"versionDatetime"`
	SsrfHostReference    struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"ssrfHostReference"`
	SubPath                        string `json:"subPath"`
	SessionTrackingStatusReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"sessionTrackingStatusReference"`
	Active            bool `json:"active"`
	AuditLogReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"auditLogReference"`
	DisallowedGeolocationReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"disallowedGeolocationReference"`
	RedirectionProtectionDomainReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"redirectionProtectionDomainReference"`
	ApplicationServiceManagedUpdatesOnly bool   `json:"applicationServiceManagedUpdatesOnly"`
	Type                                 string `json:"type"`
	SignatureSettingReference            struct {
		Link string `json:"link"`
	} `json:"signatureSettingReference"`
	DeceptionResponsePageReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"deceptionResponsePageReference"`
	WebsocketURLReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"websocketUrlReference"`
	XMLProfileReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"xmlProfileReference"`
	MethodReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"methodReference"`
	VulnerabilityReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"vulnerabilityReference"`
	RedirectionProtectionReference struct {
		Link string `json:"link"`
	} `json:"redirectionProtectionReference"`
	PolicyBuilderSessionsAndLoginsReference struct {
		Link string `json:"link"`
	} `json:"policyBuilderSessionsAndLoginsReference"`
	TemplateReference struct {
		Link  string `json:"link"`
		Title string `json:"title"`
	} `json:"templateReference"`
	PolicyBuilderHeaderReference struct {
		Link string `json:"link"`
	} `json:"policyBuilderHeaderReference"`
	CreatorName  string `json:"creatorName"`
	URLReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"urlReference"`
	HeaderReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"headerReference"`
	ActionItemReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"actionItemReference"`
	MicroserviceReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"microserviceReference"`
	XMLValidationFileReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"xmlValidationFileReference"`
	LastUpdateMicros     int64 `json:"lastUpdateMicros"`
	JSONProfileReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"jsonProfileReference"`
	BruteForceAttackPreventionReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"bruteForceAttackPreventionReference"`
	DisabledActionItemReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"disabledActionItemReference"`
	JSONValidationFileReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"jsonValidationFileReference"`
	ExtractionReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"extractionReference"`
	CharacterSetReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"characterSetReference"`
	SuggestionReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"suggestionReference"`
	DeceptionSettingsReference struct {
		Link string `json:"link"`
	} `json:"deceptionSettingsReference"`
	IsModified                  bool `json:"isModified"`
	SensitiveParameterReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"sensitiveParameterReference"`
	GeneralReference struct {
		Link string `json:"link"`
	} `json:"generalReference"`
	VersionPolicyName                          string `json:"versionPolicyName"`
	PolicyBuilderCentralConfigurationReference struct {
		Link string `json:"link"`
	} `json:"policyBuilderCentralConfigurationReference"`
}

type F5AssignServerTecnologyToSecurityPolicy struct {
	Kind                      string `json:"kind"`
	SelfLink                  string `json:"selfLink"`
	ID                        string `json:"id"`
	LastUpdateMicros          int64  `json:"lastUpdateMicros"`
	ServerTechnologyReference struct {
		Link                 string `json:"link"`
		ServerTechnologyName string `json:"serverTechnologyName"`
	} `json:"serverTechnologyReference"`
}

type F5SetLearningModePolicy struct {
	ResponseStatusCodes                []string `json:"responseStatusCodes"`
	LearningMode                       string   `json:"learningMode"`
	UntrustedTrafficSiteChangeTracking struct {
		MaxDaysBetweenSamples    int `json:"maxDaysBetweenSamples"`
		MinMinutesBetweenSamples int `json:"minMinutesBetweenSamples"`
		DifferentSources         int `json:"differentSources"`
	} `json:"untrustedTrafficSiteChangeTracking"`
	Kind                                   string `json:"kind"`
	SelfLink                               string `json:"selfLink"`
	InactiveEntityInactivityDurationInDays int    `json:"inactiveEntityInactivityDurationInDays"`
	AutoApply                              struct {
		Frequency string `json:"frequency"`
	} `json:"autoApply"`
	LearnInactiveEntities                  bool `json:"learnInactiveEntities"`
	EnableTrustedTrafficSiteChangeTracking bool `json:"enableTrustedTrafficSiteChangeTracking"`
	TrustedTrafficLoosen                   struct {
		MaxDaysBetweenSamples  int `json:"maxDaysBetweenSamples"`
		DifferentSources       int `json:"differentSources"`
		MinHoursBetweenSamples int `json:"minHoursBetweenSamples"`
	} `json:"trustedTrafficLoosen"`
	TrafficTighten struct {
		MinDaysBetweenSamples          int `json:"minDaysBetweenSamples"`
		TotalRequests                  int `json:"totalRequests"`
		MaxModificationSuggestionScore int `json:"maxModificationSuggestionScore"`
	} `json:"trafficTighten"`
	LearnOnlyFromNonBotTraffic       bool `json:"learnOnlyFromNonBotTraffic"`
	TrustedTrafficSiteChangeTracking struct {
		MaxDaysBetweenSamples    int `json:"maxDaysBetweenSamples"`
		MinMinutesBetweenSamples int `json:"minMinutesBetweenSamples"`
		DifferentSources         int `json:"differentSources"`
	} `json:"trustedTrafficSiteChangeTracking"`
	TrustAllIps                              bool `json:"trustAllIps"`
	FullyAutomatic                           bool `json:"fullyAutomatic"`
	LearnFromResponses                       bool `json:"learnFromResponses"`
	EnableUntrustedTrafficSiteChangeTracking bool `json:"enableUntrustedTrafficSiteChangeTracking"`
	UntrustedTrafficLoosen                   struct {
		MaxDaysBetweenSamples  int `json:"maxDaysBetweenSamples"`
		MinHoursBetweenSamples int `json:"minHoursBetweenSamples"`
		DifferentSources       int `json:"differentSources"`
	} `json:"untrustedTrafficLoosen"`
	EnableFullPolicyInspection bool `json:"enableFullPolicyInspection"`
}

type F5GetLearningModePolicy struct {
	ResponseStatusCodes                []string `json:"responseStatusCodes"`
	LearningMode                       string   `json:"learningMode"`
	UntrustedTrafficSiteChangeTracking struct {
		MaxDaysBetweenSamples    int `json:"maxDaysBetweenSamples"`
		MinMinutesBetweenSamples int `json:"minMinutesBetweenSamples"`
		DifferentSources         int `json:"differentSources"`
	} `json:"untrustedTrafficSiteChangeTracking"`
	Kind                                   string `json:"kind"`
	SelfLink                               string `json:"selfLink"`
	InactiveEntityInactivityDurationInDays int    `json:"inactiveEntityInactivityDurationInDays"`
	AutoApply                              struct {
		Frequency string `json:"frequency"`
	} `json:"autoApply"`
	LearnInactiveEntities                  bool `json:"learnInactiveEntities"`
	EnableTrustedTrafficSiteChangeTracking bool `json:"enableTrustedTrafficSiteChangeTracking"`
	TrustedTrafficLoosen                   struct {
		MaxDaysBetweenSamples  int `json:"maxDaysBetweenSamples"`
		DifferentSources       int `json:"differentSources"`
		MinHoursBetweenSamples int `json:"minHoursBetweenSamples"`
	} `json:"trustedTrafficLoosen"`
	TrafficTighten struct {
		MinDaysBetweenSamples          int `json:"minDaysBetweenSamples"`
		TotalRequests                  int `json:"totalRequests"`
		MaxModificationSuggestionScore int `json:"maxModificationSuggestionScore"`
	} `json:"trafficTighten"`
	LearnOnlyFromNonBotTraffic       bool `json:"learnOnlyFromNonBotTraffic"`
	TrustedTrafficSiteChangeTracking struct {
		MaxDaysBetweenSamples    int `json:"maxDaysBetweenSamples"`
		MinMinutesBetweenSamples int `json:"minMinutesBetweenSamples"`
		DifferentSources         int `json:"differentSources"`
	} `json:"trustedTrafficSiteChangeTracking"`
	TrustAllIps                              bool `json:"trustAllIps"`
	FullyAutomatic                           bool `json:"fullyAutomatic"`
	LearnFromResponses                       bool `json:"learnFromResponses"`
	EnableUntrustedTrafficSiteChangeTracking bool `json:"enableUntrustedTrafficSiteChangeTracking"`
	UntrustedTrafficLoosen                   struct {
		MaxDaysBetweenSamples  int `json:"maxDaysBetweenSamples"`
		MinHoursBetweenSamples int `json:"minHoursBetweenSamples"`
		DifferentSources       int `json:"differentSources"`
	} `json:"untrustedTrafficLoosen"`
	EnableFullPolicyInspection bool `json:"enableFullPolicyInspection"`
}

type F5GetSecurityPolicy struct {
	SignatureRequirementReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"signatureRequirementReference"`
	PlainTextProfileReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"plainTextProfileReference"`
	IdlFileReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"idlFileReference"`
	EnablePassiveMode              bool `json:"enablePassiveMode"`
	BehavioralEnforcementReference struct {
		Link string `json:"link"`
	} `json:"behavioralEnforcementReference"`
	DataGuardReference struct {
		Link string `json:"link"`
	} `json:"dataGuardReference"`
	CreatedDatetime             time.Time `json:"createdDatetime"`
	DatabaseProtectionReference struct {
		Link string `json:"link"`
	} `json:"databaseProtectionReference"`
	CookieSettingsReference struct {
		Link string `json:"link"`
	} `json:"cookieSettingsReference"`
	CsrfURLReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"csrfUrlReference"`
	VersionLastChange       string `json:"versionLastChange"`
	Name                    string `json:"name"`
	CaseInsensitive         bool   `json:"caseInsensitive"`
	HeaderSettingsReference struct {
		Link string `json:"link"`
	} `json:"headerSettingsReference"`
	SectionReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"sectionReference"`
	FlowReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"flowReference"`
	LoginPageReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"loginPageReference"`
	Description                     string `json:"description"`
	FullPath                        string `json:"fullPath"`
	PolicyBuilderParameterReference struct {
		Link string `json:"link"`
	} `json:"policyBuilderParameterReference"`
	HasParent               bool `json:"hasParent"`
	ThreatCampaignReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"threatCampaignReference"`
	Partition               string `json:"partition"`
	ManagedByBewaf          bool   `json:"managedByBewaf"`
	CsrfProtectionReference struct {
		Link string `json:"link"`
	} `json:"csrfProtectionReference"`
	GraphqlProfileReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"graphqlProfileReference"`
	PolicyAntivirusReference struct {
		Link string `json:"link"`
	} `json:"policyAntivirusReference"`
	Kind                         string   `json:"kind"`
	VirtualServers               []string `json:"virtualServers"`
	PolicyBuilderCookieReference struct {
		Link string `json:"link"`
	} `json:"policyBuilderCookieReference"`
	IPIntelligenceReference struct {
		Link string `json:"link"`
	} `json:"ipIntelligenceReference"`
	ProtocolIndependent               bool `json:"protocolIndependent"`
	SessionAwarenessSettingsReference struct {
		Link string `json:"link"`
	} `json:"sessionAwarenessSettingsReference"`
	PolicyBuilderURLReference struct {
		Link string `json:"link"`
	} `json:"policyBuilderUrlReference"`
	PolicyBuilderServerTechnologiesReference struct {
		Link string `json:"link"`
	} `json:"policyBuilderServerTechnologiesReference"`
	PolicyBuilderFiletypeReference struct {
		Link string `json:"link"`
	} `json:"policyBuilderFiletypeReference"`
	SignatureSetReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"signatureSetReference"`
	ParameterReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"parameterReference"`
	ApplicationLanguage       string `json:"applicationLanguage"`
	EnforcementMode           string `json:"enforcementMode"`
	LoginEnforcementReference struct {
		Link string `json:"link"`
	} `json:"loginEnforcementReference"`
	OpenAPIFileReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"openApiFileReference"`
	NavigationParameterReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"navigationParameterReference"`
	GwtProfileReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"gwtProfileReference"`
	WebhookReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"webhookReference"`
	WhitelistIPReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"whitelistIpReference"`
	HistoryRevisionReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"historyRevisionReference"`
	PolicyBuilderReference struct {
		Link string `json:"link"`
	} `json:"policyBuilderReference"`
	ResponsePageReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"responsePageReference"`
	VulnerabilityAssessmentReference struct {
		Link string `json:"link"`
	} `json:"vulnerabilityAssessmentReference"`
	ServerTechnologyReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"serverTechnologyReference"`
	CookieReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"cookieReference"`
	BlockingSettingReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"blockingSettingReference"`
	HostNameReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"hostNameReference"`
	VersionDeviceName              string `json:"versionDeviceName"`
	SelfLink                       string `json:"selfLink"`
	ThreatCampaignSettingReference struct {
		Link string `json:"link"`
	} `json:"threatCampaignSettingReference"`
	SignatureReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"signatureReference"`
	PolicyBuilderRedirectionProtectionReference struct {
		Link string `json:"link"`
	} `json:"policyBuilderRedirectionProtectionReference"`
	FiletypeReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"filetypeReference"`
	ID                   string        `json:"id"`
	ModifierName         string        `json:"modifierName"`
	ManualVirtualServers []interface{} `json:"manualVirtualServers"`
	VersionDatetime      time.Time     `json:"versionDatetime"`
	SsrfHostReference    struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"ssrfHostReference"`
	SubPath                        string `json:"subPath"`
	SessionTrackingStatusReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"sessionTrackingStatusReference"`
	Active            bool `json:"active"`
	AuditLogReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"auditLogReference"`
	DisallowedGeolocationReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"disallowedGeolocationReference"`
	RedirectionProtectionDomainReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"redirectionProtectionDomainReference"`
	Type                      string `json:"type"`
	SignatureSettingReference struct {
		Link string `json:"link"`
	} `json:"signatureSettingReference"`
	DeceptionResponsePageReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"deceptionResponsePageReference"`
	WebsocketURLReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"websocketUrlReference"`
	XMLProfileReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"xmlProfileReference"`
	MethodReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"methodReference"`
	VulnerabilityReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"vulnerabilityReference"`
	RedirectionProtectionReference struct {
		Link string `json:"link"`
	} `json:"redirectionProtectionReference"`
	PolicyBuilderSessionsAndLoginsReference struct {
		Link string `json:"link"`
	} `json:"policyBuilderSessionsAndLoginsReference"`
	TemplateReference struct {
		Link  string `json:"link"`
		Title string `json:"title"`
	} `json:"templateReference"`
	PolicyBuilderHeaderReference struct {
		Link string `json:"link"`
	} `json:"policyBuilderHeaderReference"`
	CreatorName  string `json:"creatorName"`
	URLReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"urlReference"`
	HeaderReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"headerReference"`
	ActionItemReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"actionItemReference"`
	MicroserviceReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"microserviceReference"`
	XMLValidationFileReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"xmlValidationFileReference"`
	LastUpdateMicros     int64 `json:"lastUpdateMicros"`
	JSONProfileReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"jsonProfileReference"`
	BruteForceAttackPreventionReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"bruteForceAttackPreventionReference"`
	DisabledActionItemReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"disabledActionItemReference"`
	JSONValidationFileReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"jsonValidationFileReference"`
	ExtractionReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"extractionReference"`
	CharacterSetReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"characterSetReference"`
	SuggestionReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"suggestionReference"`
	DeceptionSettingsReference struct {
		Link string `json:"link"`
	} `json:"deceptionSettingsReference"`
	IsModified                  bool `json:"isModified"`
	SensitiveParameterReference struct {
		Link            string `json:"link"`
		IsSubCollection bool   `json:"isSubCollection"`
	} `json:"sensitiveParameterReference"`
	GeneralReference struct {
		Link string `json:"link"`
	} `json:"generalReference"`
	VersionPolicyName                          string `json:"versionPolicyName"`
	PolicyBuilderCentralConfigurationReference struct {
		Link string `json:"link"`
	} `json:"policyBuilderCentralConfigurationReference"`
}

type F5BotDefense struct {
	Kind                      string `json:"kind"`
	Name                      string `json:"name"`
	Partition                 string `json:"partition"`
	FullPath                  string `json:"fullPath"`
	Generation                int    `json:"generation"`
	SelfLink                  string `json:"selfLink"`
	AllowBrowserAccess        string `json:"allowBrowserAccess"`
	APIAccessStrictMitigation string `json:"apiAccessStrictMitigation"`
	BlockingPage              struct {
		Body       string `json:"body"`
		Headers    string `json:"headers"`
		StatusCode int    `json:"statusCode"`
		Type       string `json:"type"`
	} `json:"blockingPage"`
	BrowserMitigationAction string `json:"browserMitigationAction"`
	CaptchaResponse         struct {
		Failure struct {
			Body string `json:"body"`
			Type string `json:"type"`
		} `json:"failure"`
		First struct {
			Body string `json:"body"`
			Type string `json:"type"`
		} `json:"first"`
	} `json:"captchaResponse"`
	CrossDomainRequests        string `json:"crossDomainRequests"`
	DeviceidMode               string `json:"deviceidMode"`
	DosAttackStrictMitigation  string `json:"dosAttackStrictMitigation"`
	EnforcementMode            string `json:"enforcementMode"`
	EnforcementReadinessPeriod int    `json:"enforcementReadinessPeriod"`
	GracePeriod                int    `json:"gracePeriod"`
	HoneypotPage               struct {
		Body       string `json:"body"`
		Headers    string `json:"headers"`
		StatusCode int    `json:"statusCode"`
		Type       string `json:"type"`
	} `json:"honeypotPage"`
	MobileDetection struct {
		AllowAndroidRootedDevice   string `json:"allowAndroidRootedDevice"`
		AllowAnyAndroidPackage     string `json:"allowAnyAndroidPackage"`
		AllowAnyIosPackage         string `json:"allowAnyIosPackage"`
		AllowEmulators             string `json:"allowEmulators"`
		AllowJailbrokenDevices     string `json:"allowJailbrokenDevices"`
		BlockDebuggerEnabledDevice string `json:"blockDebuggerEnabledDevice"`
		ClientSideChallengeMode    string `json:"clientSideChallengeMode"`
	} `json:"mobileDetection"`
	PerformChallengeInTransparent     string `json:"performChallengeInTransparent"`
	SignatureStagingUponUpdate        string `json:"signatureStagingUponUpdate"`
	SinglePageApplication             string `json:"singlePageApplication"`
	Template                          string `json:"template"`
	AnomalyCategoryOverridesReference struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"anomalyCategoryOverridesReference"`
	AnomalyOverridesReference struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"anomalyOverridesReference"`
	ClassOverridesReference struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"classOverridesReference"`
	ExternalDomainsReference struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"externalDomainsReference"`
	MicroServicesReference struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"microServicesReference"`
	SignatureCategoryOverridesReference struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"signatureCategoryOverridesReference"`
	SignatureOverridesReference struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"signatureOverridesReference"`
	SiteDomainsReference struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"siteDomainsReference"`
	StagedSignaturesReference struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"stagedSignaturesReference"`
	WhitelistReference struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"whitelistReference"`
}

type F5ListBotDefense struct {
	Kind     string         `json:"kind"`
	SelfLink string         `json:"selfLink"`
	Items    []F5BotDefense `json:"items"`
}

type F5ListLogProfile struct {
	Kind     string         `json:"kind"`
	SelfLink string         `json:"selfLink"`
	Items    []F5LogProfile `json:"items"`
}

type F5LogProfile struct {
	Kind          string `json:"kind"`
	Name          string `json:"name"`
	Partition     string `json:"partition"`
	FullPath      string `json:"fullPath"`
	Generation    int    `json:"generation"`
	SelfLink      string `json:"selfLink"`
	AutoDiscovery struct {
	} `json:"autoDiscovery"`
	BuiltIn        string `json:"builtIn"`
	Classification struct {
		LogAllClassificationMatches string `json:"logAllClassificationMatches"`
	} `json:"classification"`
	Description string `json:"description"`
	Flowspec    struct {
	} `json:"flowspec"`
	Hidden         string `json:"hidden"`
	IPIntelligence struct {
		AggregateRate        int64  `json:"aggregateRate"`
		LogGeo               string `json:"logGeo"`
		LogRtbh              string `json:"logRtbh"`
		LogScrubber          string `json:"logScrubber"`
		LogShun              string `json:"logShun"`
		LogTranslationFields string `json:"logTranslationFields"`
	} `json:"ipIntelligence"`
	Nat struct {
		EndInboundSession  string `json:"endInboundSession"`
		EndOutboundSession struct {
			Action string `json:"action"`
		} `json:"endOutboundSession"`
		Errors string `json:"errors"`
		Format struct {
			EndInboundSession struct {
				FieldListDelimiter string `json:"fieldListDelimiter"`
				Type               string `json:"type"`
			} `json:"endInboundSession"`
			EndOutboundSession struct {
				FieldListDelimiter string `json:"fieldListDelimiter"`
				Type               string `json:"type"`
			} `json:"endOutboundSession"`
			Errors struct {
				FieldListDelimiter string `json:"fieldListDelimiter"`
				Type               string `json:"type"`
			} `json:"errors"`
			QuotaExceeded struct {
				FieldListDelimiter string `json:"fieldListDelimiter"`
				Type               string `json:"type"`
			} `json:"quotaExceeded"`
			StartInboundSession struct {
				FieldListDelimiter string `json:"fieldListDelimiter"`
				Type               string `json:"type"`
			} `json:"startInboundSession"`
			StartOutboundSession struct {
				FieldListDelimiter string `json:"fieldListDelimiter"`
				Type               string `json:"type"`
			} `json:"startOutboundSession"`
		} `json:"format"`
		LogSubscriberID string `json:"logSubscriberId"`
		LsnLegacyMode   string `json:"lsnLegacyMode"`
		QuotaExceeded   string `json:"quotaExceeded"`
		RateLimit       struct {
			AggregateRate        int64 `json:"aggregateRate"`
			EndInboundSession    int64 `json:"endInboundSession"`
			EndOutboundSession   int64 `json:"endOutboundSession"`
			Errors               int64 `json:"errors"`
			QuotaExceeded        int64 `json:"quotaExceeded"`
			StartInboundSession  int64 `json:"startInboundSession"`
			StartOutboundSession int64 `json:"startOutboundSession"`
		} `json:"rateLimit"`
		StartInboundSession  string `json:"startInboundSession"`
		StartOutboundSession struct {
			Action string `json:"action"`
		} `json:"startOutboundSession"`
	} `json:"nat"`
	Netflow struct {
	} `json:"netflow"`
	PacketFilter struct {
		AggregateRate int `json:"aggregateRate"`
	} `json:"packetFilter"`
	PortMisuse struct {
		AggregateRate int64 `json:"aggregateRate"`
	} `json:"portMisuse"`
	ProtocolInspection struct {
		LogPacket string `json:"logPacket"`
	} `json:"protocolInspection"`
	TrafficStatistics struct {
		ActiveFlows         string `json:"activeFlows"`
		MissedFlows         string `json:"missedFlows"`
		ReapedFlows         string `json:"reapedFlows"`
		Syncookies          string `json:"syncookies"`
		SyncookiesWhitelist string `json:"syncookiesWhitelist"`
	} `json:"trafficStatistics"`
	ApplicationReference struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"applicationReference"`
	NetworkReference struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"networkReference"`
	ProtocolDNSReference struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"protocolDnsReference"`
	ProtocolSipReference struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"protocolSipReference"`
}

type F5AttachBotDefenseandLogProfileToVS struct {
	Kind              string    `json:"kind"`
	Name              string    `json:"name"`
	Partition         string    `json:"partition"`
	FullPath          string    `json:"fullPath"`
	Generation        int       `json:"generation"`
	SelfLink          string    `json:"selfLink"`
	AddressStatus     string    `json:"addressStatus"`
	AutoDiscovery     string    `json:"autoDiscovery"`
	AutoLasthop       string    `json:"autoLasthop"`
	CmpEnabled        string    `json:"cmpEnabled"`
	ConnectionLimit   int       `json:"connectionLimit"`
	CreationTime      time.Time `json:"creationTime"`
	Destination       string    `json:"destination"`
	Enabled           bool      `json:"enabled"`
	EvictionProtected string    `json:"evictionProtected"`
	GtmScore          int       `json:"gtmScore"`
	IPProtocol        string    `json:"ipProtocol"`
	LastModifiedTime  time.Time `json:"lastModifiedTime"`
	Mask              string    `json:"mask"`
	Mirror            string    `json:"mirror"`
	MobileAppTunnel   string    `json:"mobileAppTunnel"`
	Nat64             string    `json:"nat64"`
	Pool              string    `json:"pool"`
	PoolReference     struct {
		Link string `json:"link"`
	} `json:"poolReference"`
	RateLimit         string `json:"rateLimit"`
	RateLimitDstMask  int    `json:"rateLimitDstMask"`
	RateLimitMode     string `json:"rateLimitMode"`
	RateLimitSrcMask  int    `json:"rateLimitSrcMask"`
	SecurityNatPolicy struct {
		UseDevicePolicy      string `json:"useDevicePolicy"`
		UseRouteDomainPolicy string `json:"useRouteDomainPolicy"`
	} `json:"securityNatPolicy"`
	ServersslUseSni            string `json:"serversslUseSni"`
	ServiceDownImmediateAction string `json:"serviceDownImmediateAction"`
	Source                     string `json:"source"`
	SourceAddressTranslation   struct {
		Type string `json:"type"`
	} `json:"sourceAddressTranslation"`
	SourcePort                   string   `json:"sourcePort"`
	SynCookieStatus              string   `json:"synCookieStatus"`
	ThroughputCapacity           string   `json:"throughputCapacity"`
	TranslateAddress             string   `json:"translateAddress"`
	TranslatePort                string   `json:"translatePort"`
	VlansDisabled                bool     `json:"vlansDisabled"`
	VsIndex                      int      `json:"vsIndex"`
	SecurityLogProfiles          []string `json:"securityLogProfiles"`
	SecurityLogProfilesReference []struct {
		Link string `json:"link"`
	} `json:"securityLogProfilesReference"`
	PoliciesReference struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"policiesReference"`
	ProfilesReference struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"profilesReference"`
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

type F5ListCertAndChain struct {
	Kind     string           `json:"kind"`
	SelfLink string           `json:"selfLink"`
	Items    []F5CertAndChain `json:"items"`
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

type F5ListKey struct {
	Kind     string  `json:"kind"`
	SelfLink string  `json:"selfLink"`
	Items    []F5Key `json:"items"`
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

type F5ListSNI struct {
	Kind     string  `json:"kind"`
	SelfLink string  `json:"selfLink"`
	Items    []F5SNI `json:"items"`
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
