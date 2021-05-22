package conoha_server_go

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func init() {
	_ = godotenv.Load()
}

func TestConoha_GetPaymentSummary(t *testing.T) {
	cl, err := NewConohaClient(os.Getenv("c_user_name"), os.Getenv("c_password"), os.Getenv("c_tenant_id"))
	if err != nil {
		t.Fatal(err)
	}
	summary, err := cl.GetPaymentSummary()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n", summary)
}
