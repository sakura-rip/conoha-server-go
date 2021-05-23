package conoha_server_go

import (
	"github.com/imroc/req"
	"golang.org/x/xerrors"
)

type Port struct {
	Status              string        `json:"status"`
	Name                string        `json:"name"`
	AllowedAddressPairs []interface{} `json:"allowed_address_pairs"`
	AdminStateUp        bool          `json:"admin_state_up"`
	NetworkID           string        `json:"network_id"`
	TenantID            string        `json:"tenant_id"`
	ExtraDhcpOpts       []interface{} `json:"extra_dhcp_opts"`
	BindingVnicType     string        `json:"binding:vnic_type"`
	DeviceOwner         string        `json:"device_owner"`
	MacAddress          string        `json:"mac_address"`
	FixedIps            []struct {
		SubnetID  string `json:"subnet_id"`
		IPAddress string `json:"ip_address"`
	} `json:"fixed_ips"`
	ID             string   `json:"id"`
	SecurityGroups []string `json:"security_groups"`
	DeviceID       string   `json:"device_id"`
}
type GetPortListResponse struct {
	Ports []Port `json:"ports"`
}

func (c *Conoha) GetPortList() ([]Port, error) {
	r, err := req.Get(c.endPoint.ToUrl(NetworkService, "ports"))
	if err != nil {
		return nil, err
	}
	if r.Response().StatusCode != 200 {
		return nil, xerrors.Errorf("wrong status code: %v, message: %v", r.Response().StatusCode, r.String())
	}
	res := GetPortListResponse{}
	err = r.ToJSON(&res)
	if err != nil {
		return nil, err
	}
	return res.Ports, nil
}

type AddPortRequest struct {
	Port AddPortRequestPort `json:"port"`
}
type AddPortRequestFixedIps struct {
	IPAddress string `json:"ip_address"`
	SubnetID  string `json:"subnet_id"`
}
type AddPortRequestPort struct {
	NetworkID string                   `json:"network_id"`
	FixedIps  []AddPortRequestFixedIps `json:"fixed_ips"`
}
type AddPortResponse struct {
	Port *Port `json:"port"`
}

func (c *Conoha) AddPort(networkID string) (*Port, error) {
	data := req.BodyJSON(AddPortRequest{AddPortRequestPort{NetworkID: networkID, FixedIps: []AddPortRequestFixedIps{}}})
	r, err := req.Get(c.endPoint.ToUrl(NetworkService), data)
	if err != nil {
		return nil, err
	}
	if r.Response().StatusCode != 201 {
		return nil, xerrors.Errorf("wrong status code: %v, message: %v", r.Response().StatusCode, r.String())
	}
	res := AddPortResponse{}
	err = r.ToJSON(&res)
	if err != nil {
		return nil, err
	}
	return res.Port, nil
}
