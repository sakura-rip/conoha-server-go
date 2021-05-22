package conoha_server_go

import (
	"fmt"
	"strings"
)

type Path string

const (
	AccountService       Path = "https://account.tyo2.conoha.io/v1/%v"
	ComputeService       Path = "https://compute.tyo2.conoha.io/v2/%v"
	VolumeService        Path = "https://block-storage.tyo2.conoha.io/v2/%v"
	DatabaseService      Path = "https://database-hosting.tyo2.conoha.io/v1"
	ImageService         Path = "https://image-service.tyo2.conoha.io"
	DNSService           Path = "https://dns-service.tyo2.conoha.io"
	ObjectStorageService Path = "https://object-storage.tyo2.conoha.io/v1/nc_%v"
	MailService          Path = "https://mail-hosting.tyo2.conoha.io/v1"
	IdentityService      Path = "https://identity.tyo2.conoha.io/v2.0"
	NetworkService       Path = "https://networking.tyo2.conoha.io"
)

type EndPoint struct {
	tenantId string
}

func NewEndPoint(tenantId string) *EndPoint {
	return &EndPoint{tenantId: tenantId}
}

func (e *EndPoint) ToUrl(path Path) string {
	if strings.Contains(string(path), "%v") {
		return fmt.Sprintf(string(path), e.tenantId)
	}
	return string(path)
}
