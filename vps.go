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

type CreateVPSRequest struct {
	Server CreateVPSRequestServer `json:"server"`
}
type CreateVPSRequestServer struct {
	Imageref  string `json:"imageRef"`
	Flavorref string `json:"flavorRef"`
	Adminpass string `json:"adminPass"`
	KeyName   string `json:"key_name"`
}

type CreateVPSResponse struct {
	Server struct {
		OsDcfDiskconfig string `json:"OS-DCF:diskConfig"`
		Adminpass       string `json:"adminPass"`
		ID              string `json:"id"`
		Links           []struct {
			Href string `json:"href"`
			Rel  string `json:"rel"`
		} `json:"links"`
		SecurityGroups []struct {
			Name string `json:"name"`
		} `json:"security_groups"`
	} `json:"server"`
}

func (c *Conoha) CreateVPS(imageRef, flovorRef, adminPassword, sshKeyName string) (string, error) {
	data := req.BodyJSON(CreateVPSRequest{Server: CreateVPSRequestServer{imageRef, flovorRef, adminPassword, sshKeyName}})
	r, err := req.Post(c.endPoint.ToUrl(ComputeService, "servers"), data)
	if err != nil {
		return "", err
	}
	res := CreateVPSResponse{}
	if r.Response().StatusCode != 200 {
		return "", xerrors.Errorf("wrong status code: %v, message: %v", r.Response().StatusCode, r.String())
	}
	err = r.ToJSON(&res)
	if err != nil {
		return "", err
	}
	return res.Server.ID, nil
}
