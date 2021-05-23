package conoha_server_go

import (
	"github.com/imroc/req"
	"golang.org/x/xerrors"
)

type Subnet struct {
	Name            string      `json:"name"`
	EnableDhcp      bool        `json:"enable_dhcp"`
	NetworkID       string      `json:"network_id"`
	TenantID        string      `json:"tenant_id"`
	DNSNameservers  []string    `json:"dns_nameservers"`
	GatewayIP       string      `json:"gateway_ip"`
	Ipv6RaMode      interface{} `json:"ipv6_ra_mode"`
	AllocationPools []struct {
		Start string `json:"start"`
		End   string `json:"end"`
	} `json:"allocation_pools"`
	HostRoutes      []interface{} `json:"host_routes"`
	IPVersion       int           `json:"ip_version"`
	Ipv6AddressMode interface{}   `json:"ipv6_address_mode"`
	Cidr            string        `json:"cidr"`
	ID              string        `json:"id"`
}

type GetSubnetListResponse struct {
	Subnets []Subnet `json:"subnets"`
}

func (c *Conoha) GetSubnetList() ([]Subnet, error) {
	r, err := req.Get(c.endPoint.ToUrl(NetworkService, "subnets"))
	if err != nil {
		return nil, err
	}
	if r.Response().StatusCode != 200 {
		return nil, xerrors.Errorf("wrong status code: %v, message: %v", r.Response().StatusCode, r.String())
	}
	res := GetSubnetListResponse{}
	err = r.ToJSON(&res)
	if err != nil {
		return nil, err
	}
	return res.Subnets, nil
}

type AddSubnetForAdditionalIpResponse struct {
	Subnet *Subnet `json:"subnet"`
}
type AddSubnetForAdditionalIpRequest struct {
	AllocateIp struct {
		Bitmask string `json:"bitmask"`
	} `json:"allocateip"`
}

// AddSubnetForAdditionalIp
// /32 =>1個
// /31 =>2個
// /30 =>4個
// /29 =>8個
// /28 =>16個
func (c *Conoha) AddSubnetForAdditionalIp(count string) (*Subnet, error) {
	data := req.BodyJSON(AddSubnetForAdditionalIpRequest{AllocateIp: struct {
		Bitmask string `json:"bitmask"`
	}(struct{ Bitmask string }{Bitmask: count})})

	r, err := req.Post(c.endPoint.ToUrl(NetworkService, "allocateips"), data)
	if err != nil {
		return nil, err
	}
	if r.Response().StatusCode != 201 {
		return nil, xerrors.Errorf("wrong status code: %v, message: %v", r.Response().StatusCode, r.String())
	}
	res := AddSubnetForAdditionalIpResponse{}
	err = r.ToJSON(&res)
	if err != nil {
		return nil, err
	}
	return res.Subnet, nil
}

func (c *Conoha) DeleteSubnet(subnetId string) error {
	r, err := req.Delete(c.endPoint.ToUrl(NetworkService, "subnets", subnetId))
	if err != nil {
		return err
	}
	if r.Response().StatusCode != 204 {
		return xerrors.Errorf("wrong status code: %v, message: %v", r.Response().StatusCode, r.String())
	}
	return nil
}
