package conoha_server_go

import (
	"os"
	"testing"
)

var cl *Conoha

func init() {
	cl, _ = NewConohaClient(os.Getenv("c_user_name"), os.Getenv("c_password"), os.Getenv("c_tenant_id"))
}

func TestConoha_GetSubnetList(t *testing.T) {
	_, err := cl.GetSubnetList()
	if err != nil {
		t.Error(err)
	}
	//fmt.Printf("%+v\n", list)
}
