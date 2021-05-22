package conoha_server_go

import (
	"github.com/imroc/req"
	"golang.org/x/xerrors"
)

type GetVPSListResponse struct {
	Servers []*GetVPSListResponseServer `json:"servers"`
}
type GetVPSListResponseServer struct {
	ID    string `json:"id"`
	Links []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
	Name string `json:"name"`
}

func (c *Conoha) GetVPSList() ([]*GetVPSListResponseServer, error) {
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
	Server *CreateVPSRequestServer `json:"server"`
}
type CreateVPSRequestServer struct {
	ImageRef  string `json:"imageRef"`
	FlavorRef string `json:"flavorRef"`
	AdminPass string `json:"adminPass"`
	KeyName   string `json:"key_name"`
}

type CreateVPSResponse struct {
	Server struct {
		OsDcfDiskConfig string `json:"OS-DCF:diskConfig"`
		AdminPass       string `json:"adminPass"`
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

func (c *Conoha) DeleteVPS(id string) error {
	r, err := req.Delete(c.endPoint.ToUrl(ComputeService, "servers", id))
	if err != nil {
		return err
	}
	if r.Response().StatusCode != 204 {
		return xerrors.Errorf("wrong status code: %v, message: %v", r.Response().StatusCode, r.String())
	}
	return nil
}

func (c *Conoha) StartVPS(id string) error {
	r, err := req.Delete(c.endPoint.ToUrl(ComputeService, "servers", id, "action"), req.BodyJSON(map[string]interface{}{"on-start": nil}))
	if err != nil {
		return err
	}
	if r.Response().StatusCode != 202 {
		return xerrors.Errorf("wrong status code: %v, message: %v", r.Response().StatusCode, r.String())
	}
	return nil
}

type RebootVPSRequest struct {
	Reboot *RebootVPSRequestReboot `json:"reboot"`
}
type RebootVPSRequestReboot struct {
	Type string `json:"type"`
}

func (c *Conoha) RebootVPS(id, bootType string) error {
	if bootType != "SOFT" && bootType != "HARD" {
		return xerrors.New("boot type must be SOFT or HARD")
	}
	body := req.BodyJSON(RebootVPSRequest{&RebootVPSRequestReboot{Type: bootType}})
	r, err := req.Post(c.endPoint.ToUrl(ComputeService, "servers", id, "action"), body)
	if err != nil {
		return err
	}
	if r.Response().StatusCode != 202 {
		return xerrors.Errorf("wrong status code: %v, message: %v", r.Response().StatusCode, r.String())
	}
	return nil
}

type ShutdownVPSRequest struct {
	OsStop *ShutdownVPSRequestOsStop `json:"os-stop"`
}
type ShutdownVPSRequestOsStop struct {
	ForceShutdown bool `json:"force_shutdown"`
}

func (c *Conoha) ShutdownVPS(id string, force bool) error {
	body := req.BodyJSON(ShutdownVPSRequest{&ShutdownVPSRequestOsStop{force}})
	r, err := req.Post(c.endPoint.ToUrl(ComputeService, "servers", id, "action"), body)
	if err != nil {
		return err
	}
	if r.Response().StatusCode != 202 {
		return xerrors.Errorf("wrong status code: %v, message: %v", r.Response().StatusCode, r.String())
	}
	return nil
}

type FixedIp struct {
	IPAddress string `json:"ip_address"`
	SubnetID  string `json:"subnet_id"`
}

type GetAttachedPortsListResponse struct {
	InterfaceAttachments []*InterfaceAttachment `json:"interfaceAttachments"`
}
type InterfaceAttachment struct {
	FixedIps []struct {
		IPAddress string `json:"ip_address"`
		SubnetID  string `json:"subnet_id"`
	} `json:"fixed_ips"`
	MacAddr   string `json:"mac_addr"`
	NetID     string `json:"net_id"`
	PortID    string `json:"port_id"`
	PortState string `json:"port_state"`
}

func (c *Conoha) GetAttachedPortsList(id string) ([]*InterfaceAttachment, error) {
	r, err := req.Get(c.endPoint.ToUrl(ComputeService, c.tenantId, "servers", id, "os-interface"))
	if err != nil {
		return nil, err
	}
	if r.Response().StatusCode != 200 {
		return nil, xerrors.Errorf("wrong status code: %v, message: %v", r.Response().StatusCode, r.String())
	}
	res := GetAttachedPortsListResponse{}
	err = r.ToJSON(&res)
	if err != nil {
		return nil, err
	}
	return res.InterfaceAttachments, nil
}

type AttachPortToVPSResponse struct {
	InterfaceAttachment *InterfaceAttachment `json:"interfaceAttachment"`
}

func (c *Conoha) AttachPortToVPS(serverId, portId string) (*InterfaceAttachment, error) {
	data := req.BodyJSON(map[string]interface{}{
		"interfaceAttachment": map[string]interface{}{"port_id": portId},
	})
	r, err := req.Post(c.endPoint.ToUrl(ComputeService, c.tenantId, "servers", serverId, "os-interface"), data)
	if err != nil {
		return nil, err
	}
	if r.Response().StatusCode != 200 {
		return nil, xerrors.Errorf("wrong status code: %v, message: %v", r.Response().StatusCode, r.String())
	}
	res := AttachPortToVPSResponse{}
	err = r.ToJSON(&res)
	if err != nil {
		return nil, err
	}
	return res.InterfaceAttachment, nil
}
