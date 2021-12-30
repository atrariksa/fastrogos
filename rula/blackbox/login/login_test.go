package login_test

//go:generate go run github.com/cucumber/godog/cmd/godog login.feature
import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/atrariksa/fastrogos/rula/blackbox"
	apiclient "github.com/atrariksa/fastrogos/rula/client"
	"github.com/atrariksa/fastrogos/rula/client/operations"
	"github.com/atrariksa/fastrogos/rula/models"
	"github.com/cucumber/godog"
)

type apiFeature struct {
}

var (
	rulaClient *apiclient.Rula
)

type loginCtxKey struct{}

func (a *apiFeature) ISendLoginRequestWithDataByMap(ctx context.Context, data *godog.Table) (context.Context, error) {
	var som = blackbox.GetSliceOfMapFromTable(data)
	resp, err := sendloginRequest(rulaClient, som[0])
	if err != nil {
		return ctx, err
	}
	ctx = context.WithValue(ctx, loginCtxKey{}, resp)
	return ctx, nil
}

func (a *apiFeature) theResponseMustMatchDataByMap(ctx context.Context, data *godog.Table) error {
	var som = blackbox.GetSliceOfMapFromTable(data)
	resp := ctx.Value(loginCtxKey{}).(*operations.LoginOK)
	if resp.Payload.Code != som[0]["code"] {
		return blackbox.ErrGotWant(resp.Payload.Code, som[0]["code"])
	}
	if resp.Payload.Message != som[0]["message"] {
		return blackbox.ErrGotWant(resp.Payload.Message, som[0]["message"])
	}
	return nil
}

func (a *apiFeature) theResponseDataShouldMatchJson(ctx context.Context, data *godog.DocString) error {
	var expected, actual models.UserData

	resp := ctx.Value(loginCtxKey{}).(*operations.LoginOK)
	bActual, _ := json.Marshal(resp.Payload.Data)
	// re-encode expected response
	if err := json.Unmarshal([]byte(data.Content), &expected); err != nil {
		return err
	}

	// re-encode actual response too
	if err := json.Unmarshal(bActual, &actual); err != nil {
		return err
	}

	// the matching may be adapted per different requirements.
	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("expected JSON does not match actual, %v vs. %v", expected, actual)
	}
	return nil
}

func InitializeTestSuite(s *godog.TestSuiteContext) {
	s.BeforeSuite(func() {
		rulaClient = blackbox.GetRulaClient()
	})
}

func InitializeScenario(s *godog.ScenarioContext) {
	api := &apiFeature{}
	s.Step(`^I send login request with data by map$`, api.ISendLoginRequestWithDataByMap)
	s.Step(`^the response must match data by map$`, api.theResponseMustMatchDataByMap)
	s.Step(`^the response data should match json:$`, api.theResponseDataShouldMatchJson)
}

func sendloginRequest(client *apiclient.Rula, params map[string]interface{}) (*operations.LoginOK, error) {

	loginParam := operations.NewLoginParams()
	loginParam.ModelsLoginReq = &models.ModelsLoginReq{
		Password: blackbox.ToString(params["password"]),
		Username: blackbox.ToString(params["username"]),
	}

	return client.Operations.Login(loginParam)
}
