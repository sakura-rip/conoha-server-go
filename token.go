package conoha_server_go

import (
	"github.com/imroc/req"
	"golang.org/x/xerrors"
	"time"
)

// GetTokenRequest request
type GetTokenRequest struct {
	Auth Auth `json:"auth"`
}
type PasswordCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Auth struct {
	PasswordCredentials PasswordCredentials `json:"passwordCredentials"`
	TenantId            string              `json:"tenantId"`
}

// GetTokenResponse response
type GetTokenResponse struct {
	Access Access `json:"access"`
}
type Tenant struct {
	DomainID    string `json:"domain_id"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	ID          string `json:"id"`
	Name        string `json:"name"`
}
type Token struct {
	IssuedAt string    `json:"issued_at"`
	Expires  time.Time `json:"expires"`
	ID       string    `json:"id"`
	Tenant   Tenant    `json:"tenant"`
	AuditIds []string  `json:"audit_ids"`
}
type Endpoints struct {
	Region    string `json:"region"`
	PublicUrl string `json:"publicURL"`
}
type ServiceCatalog struct {
	Endpoints      []Endpoints   `json:"endpoints"`
	EndpointsLinks []interface{} `json:"endpoints_links"`
	Type           string        `json:"type"`
	Name           string        `json:"name"`
}
type Roles struct {
	Name string `json:"name"`
}
type User struct {
	Username   string        `json:"username"`
	RolesLinks []interface{} `json:"roles_links"`
	ID         string        `json:"id"`
	Roles      []Roles       `json:"roles"`
	Name       string        `json:"name"`
}
type Metadata struct {
	IsAdmin int      `json:"is_admin"`
	Roles   []string `json:"roles"`
}
type Access struct {
	Token          Token            `json:"token"`
	ServiceCatalog []ServiceCatalog `json:"serviceCatalog"`
	User           User             `json:"user"`
	Metadata       Metadata         `json:"metadata"`
}

func (c *Conoha) issueApiToken(userName, password, tenantId string) (string, error) {
	body := req.BodyJSON(GetTokenRequest{Auth: Auth{
		PasswordCredentials: PasswordCredentials{
			Username: userName,
			Password: password,
		},
		TenantId: tenantId,
	}})
	r, err := req.Post(c.endPoint.ToUrl(IdentityService, "tokens"), body)
	if err != nil {
		return "", xerrors.Errorf("failed get token: %w", err)
	}
	if r.Response().StatusCode != 200 {
		return "", xerrors.Errorf("wrong status code: %v, message: %v", r.Response().StatusCode, r.String())
	}
	var res GetTokenResponse
	err = r.ToJSON(&res)
	if err != nil {
		return "", xerrors.Errorf("failed get token: %w", err)
	}
	return res.Access.Token.ID, nil
}
