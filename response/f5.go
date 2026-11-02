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
