package conoha_server_go

import (
	"github.com/imroc/req"
	"golang.org/x/xerrors"
)

type GetVPSListResponse struct {
	Servers []GetVPSListResponseServer `json:"servers"`
}
type GetVPSListResponseServer struct {
	ID    string `json:"id"`
	Links []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
	Name string `json:"name"`
}

func (c *Conoha) GetVPSList() ([]GetVPSListResponseServer, error) {
	r, err := req.Get(c.endPoint.ToUrl(ComputeService, "servers"), c.header)
	if err != nil {
		return nil, err
	}
	if r.Response().StatusCode != 200 {
		return nil, xerrors.Errorf("wrong status code: %v, message: %v", r.Response().StatusCode, r.String())
	}
	res := GetVPSListResponse{}
	err = r.ToJSON(&res)
	if err != nil {
		return nil, err
	}
	return res.Servers, nil
}
