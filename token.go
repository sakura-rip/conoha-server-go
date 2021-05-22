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

type IssueApiTokenResponse struct {
	Access struct {
		Token struct {
			IssuedAt string    `json:"issued_at"`
			Expires  time.Time `json:"expires"`
			ID       string    `json:"id"`
			Tenant   struct {
				DomainID    string `json:"domain_id"`
				Description string `json:"description"`
				Enabled     bool   `json:"enabled"`
				ID          string `json:"id"`
				Name        string `json:"name"`
			} `json:"tenant"`
			AuditIds []string `json:"audit_ids"`
		} `json:"token"`
		Servicecatalog []struct {
			Endpoints []struct {
				Region    string `json:"region"`
				Publicurl string `json:"publicURL"`
			} `json:"endpoints"`
			EndpointsLinks []interface{} `json:"endpoints_links"`
			Type           string        `json:"type"`
			Name           string        `json:"name"`
		} `json:"serviceCatalog"`
		User struct {
			Username   string        `json:"username"`
			RolesLinks []interface{} `json:"roles_links"`
			ID         string        `json:"id"`
			Roles      []struct {
				Name string `json:"name"`
			} `json:"roles"`
			Name string `json:"name"`
		} `json:"user"`
		Metadata struct {
			IsAdmin int      `json:"is_admin"`
			Roles   []string `json:"roles"`
		} `json:"metadata"`
	} `json:"access"`
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
	var res IssueApiTokenResponse
	err = r.ToJSON(&res)
	if err != nil {
		return "", xerrors.Errorf("failed get token: %w", err)
	}
	return res.Access.Token.ID, nil
}
