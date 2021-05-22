package conoha_server_go

type Conoha struct {
	header   map[string]string
	endPoint *EndPoint
}

func NewConohaClient(userName, password, tenantId string) *Conoha {
	cl := &Conoha{endPoint: NewEndPoint(tenantId), header: map[string]string{}}
	token, err := cl.issueApiToken(userName, password, tenantId)
	if err != nil {
		panic(err)
	}
	cl.header["X-Auth-Token"] = token
	return cl
}
