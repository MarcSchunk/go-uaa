package uaa

// IdentityZonesEndpoint is the path to the users resource.
const IdentityZonesEndpoint string = "/identity-zones"

// IdentityZone is a UAA identity zone.
// http://docs.cloudfoundry.org/api/uaa/version/4.14.0/index.html#identity-zones
type IdentityZone struct {
	ID           string             `json:"id,omitempty"`
	Subdomain    string             `json:"subdomain"`
	Config       IdentityZoneConfig `json:"config"`
	Name         string             `json:"name"`
	Version      int                `json:"version,omitempty"`
	Description  string             `json:"description,omitempty"`
	Created      int                `json:"created,omitempty"`
	LastModified int                `json:"last_modified,omitempty"`
}

type ClientSecretPolicy struct {
	MinLength                 int `json:"minLength,omitempty"`
	MaxLength                 int `json:"maxLength,omitempty"`
	RequireUpperCaseCharacter int `json:"requireUpperCaseCharacter,omitempty"`
	RequireLowerCaseCharacter int `json:"requireLowerCaseCharacter,omitempty"`
	RequireDigit              int `json:"requireDigit,omitempty"`
	RequireSpecialCharacter   int `json:"requireSpecialCharacter,omitempty"`
}

type TokenPolicy struct {
	AccessTokenValidity  int    `json:"accessTokenValidity,omitempty"`
	RefreshTokenValidity int    `json:"refreshTokenValidity,omitempty"`
	JwtRevocable         bool   `json:"jwtRevocable,omitempty"`
	RefreshTokenUnique   bool   `json:"refreshTokenUnique,omitempty"`
	RefreshTokenFormat   string `json:"refreshTokenFormat,omitempty"`
	ActiveKeyID          string `json:"activeKeyId,omitempty"`
}

type SAMLConfig struct {
	AssertionSigned            bool   `json:"assertionSigned,omitempty"`
	RequestSigned              bool   `json:"requestSigned,omitempty"`
	WantAssertionSigned        bool   `json:"wantAssertionSigned,omitempty"`
	WantAuthnRequestSigned     bool   `json:"wantAuthnRequestSigned,omitempty"`
	AssertionTimeToLiveSeconds int    `json:"assertionTimeToLiveSeconds,omitempty"`
	ActiveKeyID                string `json:"activeKeyId,omitempty"`
	Keys                       struct {
		Key1 struct {
			Certificate string `json:"certificate,omitempty"`
		} `json:"key1,omitempty"`
	} `json:"keys,omitempty"`
	DisableInResponseToCheck bool `json:"disableInResponseToCheck,omitempty"`
}

type CORSPolicy struct {
	XhrConfiguration struct {
		AllowedOrigins        []string      `json:"allowedOrigins,omitempty"`
		AllowedOriginPatterns []interface{} `json:"allowedOriginPatterns,omitempty"`
		AllowedUris           []string      `json:"allowedUris,omitempty"`
		AllowedURIPatterns    []interface{} `json:"allowedUriPatterns,omitempty"`
		AllowedHeaders        []string      `json:"allowedHeaders,omitempty"`
		AllowedMethods        []string      `json:"allowedMethods,omitempty"`
		AllowedCredentials    bool          `json:"allowedCredentials,omitempty"`
		MaxAge                int           `json:"maxAge,omitempty"`
	} `json:"xhrConfiguration,omitempty"`
	DefaultConfiguration struct {
		AllowedOrigins        []string      `json:"allowedOrigins,omitempty"`
		AllowedOriginPatterns []interface{} `json:"allowedOriginPatterns,omitempty"`
		AllowedUris           []string      `json:"allowedUris,omitempty"`
		AllowedURIPatterns    []interface{} `json:"allowedUriPatterns,omitempty"`
		AllowedHeaders        []string      `json:"allowedHeaders,omitempty"`
		AllowedMethods        []string      `json:"allowedMethods,omitempty"`
		AllowedCredentials    bool          `json:"allowedCredentials,omitempty"`
		MaxAge                int           `json:"maxAge,omitempty"`
	} `json:"defaultConfiguration,omitempty"`
}

type IdentityZoneLinks struct {
	Logout struct {
		RedirectURL              string   `json:"redirectUrl,omitempty"`
		RedirectParameterName    string   `json:"redirectParameterName,omitempty"`
		DisableRedirectParameter bool     `json:"disableRedirectParameter,omitempty"`
		Whitelist                []string `json:"whitelist,omitempty"`
	} `json:"logout,omitempty"`
	HomeRedirect string `json:"homeRedirect,omitempty"`
	SelfService  struct {
		SelfServiceLinksEnabled bool   `json:"selfServiceLinksEnabled,omitempty"`
		Signup                  string `json:"signup,omitempty"`
		Passwd                  string `json:"passwd,omitempty"`
	} `json:"selfService,omitempty"`
}

type Prompt struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

type Branding struct {
	CompanyName string `json:"companyName,omitempty"`
	ProductLogo string `json:"productLogo,omitempty"`
	SquareLogo  string `json:"squareLogo,omitempty"`
}

type IdentityZoneUserConfig struct {
	DefaultGroups []string `json:"defaultGroups,omitempty"`
}

type IdentityZoneMFAConfig struct {
	Enabled      *bool  `json:"enabled,omitempty"`
	ProviderName string `json:"providerName,omitempty"`
}

type IdentityZoneConfig struct {
	ClientSecretPolicy    *ClientSecretPolicy     `json:"clientSecretPolicy,omitempty"`
	TokenPolicy           *TokenPolicy            `json:"tokenPolicy,omitempty"`
	SAMLConfig            *SAMLConfig             `json:"samlConfig,omitempty"`
	CORSPolicy            *CORSPolicy             `json:"corsPolicy,omitempty"`
	Links                 *IdentityZoneLinks      `json:"links,omitempty"`
	Prompts               []Prompt                `json:"prompts,omitempty"`
	IDPDiscoveryEnabled   *bool                   `json:"idpDiscoveryEnabled,omitempty"`
	Branding              *Branding               `json:"branding,omitempty"`
	AccountChooserEnabled *bool                   `json:"accountChooserEnabled,omitempty"`
	UserConfig            *IdentityZoneUserConfig `json:"userConfig,omitempty"`
	MFAConfig             *IdentityZoneMFAConfig  `json:"mfaConfig,omitempty"`
}
