package services_clients

import (
	"github.com/atrariksa/fastrogos/fenuma/configs"
	apiclient "github.com/atrariksa/fastrogos/rula/client"
	"github.com/atrariksa/fastrogos/rula/client/operations"
	rulamodels "github.com/atrariksa/fastrogos/rula/models"
	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func GetRulaClient() *apiclient.Rula {
	cfg := configs.GetFrom("../../.env")

	// create the transport
	transport := httptransport.New(cfg.Rula.Hostname, "", nil)
	transport.Debug = true

	// create the API client, with the transport
	client := apiclient.New(transport, strfmt.Default)

	// to override the host for the default client
	// apiclient.Default.SetTransport(transport)
	return client
}

func SendCreateUserRequest(client *apiclient.Rula, req *rulamodels.ModelsCreateUserReq) (*operations.CreateUserCreated, error) {

	createUserParam := operations.NewCreateUserParams()
	createUserParam.ModelsCreateUserReq = req

	return client.Operations.CreateUser(createUserParam)
}
