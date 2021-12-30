package create_user_test

//go:generate go run github.com/cucumber/godog/cmd/godog create_user.feature
import (
	"context"

	"github.com/atrariksa/fastrogos/rula/blackbox"
	apiclient "github.com/atrariksa/fastrogos/rula/client"
	"github.com/atrariksa/fastrogos/rula/client/operations"
	"github.com/atrariksa/fastrogos/rula/models"
	"github.com/cucumber/godog"
)

type apiCreateUserFeature struct {
}

var (
	rulaClient *apiclient.Rula
)

type createUserCtxKey struct{}

func (a *apiCreateUserFeature) iSendCreateUserRequestWithDataByMap(ctx context.Context, data *godog.Table) (context.Context, error) {
	var som = blackbox.GetSliceOfMapFromTable(data)
	resp, err := sendCreateUserRequest(rulaClient, som[0])
	if err != nil {
		return ctx, err
	}
	ctx = context.WithValue(ctx, createUserCtxKey{}, resp)
	return ctx, nil
}

func (a *apiCreateUserFeature) theResponseMustMatchDataByMap(ctx context.Context, data *godog.Table) error {
	var som = blackbox.GetSliceOfMapFromTable(data)
	resp := ctx.Value(createUserCtxKey{}).(*operations.CreateUserCreated)
	if resp.Payload.Code != som[0]["code"] {
		return blackbox.ErrGotWant(resp.Payload.Code, som[0]["code"])
	}
	if resp.Payload.Message != som[0]["message"] {
		return blackbox.ErrGotWant(resp.Payload.Message, som[0]["message"])
	}
	return nil
}

func InitializeTestSuite(s *godog.TestSuiteContext) {
	s.BeforeSuite(func() {
		rulaClient = blackbox.GetRulaClient()
	})
}

func InitializeScenario(s *godog.ScenarioContext) {
	api := &apiCreateUserFeature{}
	s.Step(`^I send create user request with data by map$`, api.iSendCreateUserRequestWithDataByMap)
	s.Step(`^the response must match data by map$`, api.theResponseMustMatchDataByMap)
}

func sendCreateUserRequest(client *apiclient.Rula, params map[string]interface{}) (*operations.CreateUserCreated, error) {

	createUserParam := operations.NewCreateUserParams()
	createUserParam.ModelsCreateUserReq = &models.ModelsCreateUserReq{
		Role:     blackbox.ToString(params["role"]),
		Email:    blackbox.ToString(params["email"]),
		Password: blackbox.ToString(params["password"]),
		Username: blackbox.ToString(params["username"]),
	}

	return client.Operations.CreateUser(createUserParam)
}
