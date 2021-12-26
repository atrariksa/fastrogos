package client_services

import (
	"github.com/atrariksa/fastrogos/fenuma/configs"
	"github.com/atrariksa/fastrogos/fenuma/models"
	apiclient "github.com/atrariksa/fastrogos/rula/client"
	"github.com/atrariksa/fastrogos/rula/client/operations"
	rulamodels "github.com/atrariksa/fastrogos/rula/models"
	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

var rulaClient *apiclient.Rula

func GetRulaClient() *apiclient.Rula {
	if rulaClient != nil {
		return rulaClient
	}
	cfg := configs.Get()

	// create the transport
	transport := httptransport.New(cfg.Rula.Hostname, "", nil)
	transport.Debug = true

	// create the API client, with the transport
	client := apiclient.New(transport, strfmt.Default)

	// to override the host for the default client
	// apiclient.Default.SetTransport(transport)
	return client
}

func SendCreateUserRequest(req *rulamodels.ModelsCreateUserReq) (*operations.CreateUserCreated, error) {

	createUserParam := operations.NewCreateUserParams()
	createUserParam.ModelsCreateUserReq = req

	return GetRulaClient().Operations.CreateUser(createUserParam)
}

func Login(req models.LoginReq) (*operations.LoginOK, error) {

	loginParam := operations.NewLoginParams()
	loginParam.ModelsLoginReq = &rulamodels.ModelsLoginReq{
		Username: req.Username,
		Password: req.Password,
	}

	return GetRulaClient().Operations.Login(loginParam)
}
