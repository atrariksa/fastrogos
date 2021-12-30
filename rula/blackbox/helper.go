package blackbox

import (
	"fmt"

	apiclient "github.com/atrariksa/fastrogos/rula/client"
	"github.com/atrariksa/fastrogos/rula/configs"
	"github.com/cucumber/godog"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

func GetSliceOfMapFromTable(data *godog.Table) []map[string]interface{} {
	var som = make([]map[string]interface{}, 0)
	var mKeys = make(map[int]string)
	for k, v := range data.Rows[0].Cells {
		mKeys[k] = v.Value
	}
	for rk, rv := range data.Rows {
		if rk == 0 {
			continue
		}
		m := make(map[string]interface{})
		for k, v := range rv.Cells {
			m[mKeys[k]] = v.Value
		}
		som = append(som, m)
	}
	return som
}

func ToString(input interface{}) string {
	return fmt.Sprintf("%v", input)
}

func ErrGotWant(got, want interface{}) error {
	return fmt.Errorf("got: %v, want: %v", got, want)
}

func GetRulaClient() *apiclient.Rula {
	cfg := configs.GetFrom("../../.env")

	// create the transport
	transport := httptransport.New(cfg.App.Hostname, "", nil)
	transport.Debug = true

	// create the API client, with the transport
	client := apiclient.New(transport, strfmt.Default)

	// to override the host for the default client
	// apiclient.Default.SetTransport(transport)
	return client
}
