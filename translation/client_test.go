package translation_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/domicmeia/gcp_practice/translation"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestHelloClientSuite(t *testing.T) {
	suite.Run(t, new(HelloClientSuite))
}

type HelloClientSuite struct {
	suite.Suite
	mockServerService *MockService
	server            *httptest.Server
	underTest         translation.HelloClient
}

type MockService struct {
	mock.Mock
}

func (m *MockService) Translate(word, language string) (string, error) {
	args := m.Called(word, language)
	return args.String(0), args.Error(1)
}

func (suite *HelloClientSuite) SetupSuite() {
	suite.mockServerService = new(MockService)
	handler := func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		defer r.Body.Close()

		var m map[string]interface{}
		_ = json.Unmarshal(b, &m)

		word := m["word"].(string)
		language := m["language"].(string)

		resp, err := suite.mockServerService.Translate(word, language)

		if err != nil {
			http.Error(w, "error", 500)
		}

		if resp == "" {
			http.Error(w, "missing", 404)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, resp)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	suite.server = httptest.NewServer(mux)
}

func (suite *HelloClientSuite) TearDownSuite() {
	suite.server.Close()
}
