package blackbox

//go:generate go run github.com/cucumber/godog/cmd/godog create_user.feature
import (
	"context"

	"github.com/go-openapi/strfmt"

	apiclient "github.com/atrariksa/fastrogos/rula/client"
	"github.com/atrariksa/fastrogos/rula/client/operations"
	"github.com/atrariksa/fastrogos/rula/configs"
	"github.com/atrariksa/fastrogos/rula/models"
	"github.com/cucumber/godog"
	httptransport "github.com/go-openapi/runtime/client"
)

type apiFeature struct {
}

var (
	rulaClient *apiclient.Rula
)

type createUserCtxKey struct{}

func (a *apiFeature) iSendCreateUserRequestWithDataByMap(ctx context.Context, data *godog.Table) (context.Context, error) {
	var som = GetSliceOfMapFromTable(data)
	resp, err := sendCreateUserRequest(rulaClient, som[0])
	if err != nil {
		return ctx, err
	}
	ctx = context.WithValue(ctx, createUserCtxKey{}, resp)
	return ctx, nil
}

func (a *apiFeature) theResponseMustMatchDataByMap(ctx context.Context, data *godog.Table) error {
	var som = GetSliceOfMapFromTable(data)
	resp := ctx.Value(createUserCtxKey{}).(*operations.CreateUserCreated)
	if resp.Payload.Code != som[0]["code"] {
		return ErrGotWant(resp.Payload.Code, som[0]["code"])
	}
	if resp.Payload.Message != som[0]["message"] {
		return ErrGotWant(resp.Payload.Message, som[0]["message"])
	}
	return nil
}

func InitializeTestSuite(s *godog.TestSuiteContext) {
	s.BeforeSuite(func() {
		rulaClient = getRulaClient()
	})
}

func InitializeScenario(s *godog.ScenarioContext) {
	api := &apiFeature{}
	s.Step(`^I send create user request with data by map$`, api.iSendCreateUserRequestWithDataByMap)
	s.Step(`^the response must match data by map$`, api.theResponseMustMatchDataByMap)
}

func getRulaClient() *apiclient.Rula {
	cfg := configs.GetFrom("../.env")

	// create the transport
	transport := httptransport.New(cfg.App.Hostname, "", nil)
	transport.Debug = true

	// create the API client, with the transport
	client := apiclient.New(transport, strfmt.Default)

	// to override the host for the default client
	// apiclient.Default.SetTransport(transport)
	return client
}

func sendCreateUserRequest(client *apiclient.Rula, params map[string]interface{}) (*operations.CreateUserCreated, error) {

	createUserParam := operations.NewCreateUserParams()
	createUserParam.ModelsCreateUserReq = &models.ModelsCreateUserReq{
		Role:     ToString(params["role"]),
		Email:    ToString(params["email"]),
		Password: ToString(params["password"]),
		Username: ToString(params["username"]),
	}

	return client.Operations.CreateUser(createUserParam)
}
