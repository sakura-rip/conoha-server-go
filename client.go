package conoha_server_go

import "golang.org/x/xerrors"

type Conoha struct {
	header   map[string]string
	endPoint *EndPoint
}

func NewConohaClient(userName, password, tenantId string) (*Conoha, error) {
	cl := &Conoha{endPoint: NewEndPoint(tenantId), header: map[string]string{}}
	token, err := cl.issueApiToken(userName, password, tenantId)
	if err != nil {
		return nil, xerrors.Errorf("failed create conoha client: %w", err)
	}
	cl.header["X-Auth-Token"] = token
	return cl, nil
}
