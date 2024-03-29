package accrual_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/sergeizaitcev/gophermart/internal/gophermart/clients/accrual"
	"github.com/sergeizaitcev/gophermart/internal/gophermart/domain"
)

type TransportMock struct {
	mock.Mock
	header http.Header
}

func NewTransportMock() *TransportMock {
	return &TransportMock{
		header: make(http.Header),
	}
}

func (m *TransportMock) RoundTrip(req *http.Request) (*http.Response, error) {
	args := m.Called(req.Method, req.URL.EscapedPath())

	err := args.Error(2)
	if err != nil {
		return nil, err
	}

	statusCode := args.Int(0)
	data := args.Get(1)

	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	res := &http.Response{
		StatusCode: statusCode,
		Status:     http.StatusText(statusCode),
		Header:     m.header.Clone(),
		Body:       io.NopCloser(bytes.NewReader(b)),
		Request:    req,
	}

	return res, nil
}

type ClientSuite struct {
	suite.Suite

	transport *TransportMock
	client    *accrual.Client
}

func TestClient(t *testing.T) {
	suite.Run(t, new(ClientSuite))
}

func (suite *ClientSuite) SetupTest() {
	suite.transport = NewTransportMock()
	suite.client = accrual.NewClient("//localhost", &accrual.ClientOption{
		Transport: suite.transport,
	})
}

func (suite *ClientSuite) TestOK() {
	want := map[string]any{
		"Order":   "49927398716",
		"Status":  domain.AccrualStatusRegistered.String(),
		"Accrual": 1000.0,
	}

	suite.transport.On("RoundTrip", "GET", "/api/orders/1").
		Return(http.StatusOK, want, nil)

	got, err := suite.client.GetAccrualInfo(context.Background(), "1")
	if suite.NoError(err) {
		suite.EqualValues(want["Order"], got.OrderNumber)
		suite.EqualValues(want["Status"], got.Status.String())
		suite.EqualValues(want["Accrual"], got.Accrual.Float64())
	}
}

func (suite *ClientSuite) TestNotRegistered() {
	suite.transport.On("RoundTrip", "GET", "/api/orders/2").
		Return(http.StatusNoContent, struct{}{}, nil)

	_, err := suite.client.GetAccrualInfo(context.Background(), "2")
	suite.ErrorIs(err, domain.ErrOrderNotRegistered)
}

func (suite *ClientSuite) TestTooManyRequest() {
	suite.transport.header.Add("Retry-After", "60")
	suite.transport.On("RoundTrip", "GET", "/api/orders/3").
		Return(http.StatusTooManyRequests, struct{}{}, nil)

	_, err := suite.client.GetAccrualInfo(context.Background(), "3")

	var exhausted *domain.ResourceExhaustedError
	if suite.ErrorAs(err, &exhausted) {
		suite.Equal(60*time.Second, exhausted.RetryAfter)
	}
}

func (suite *ClientSuite) TestInternalServerError() {
	suite.transport.On("RoundTrip", "GET", "/api/orders/4").
		Return(http.StatusInternalServerError, struct{}{}, nil)

	_, err := suite.client.GetAccrualInfo(context.Background(), "4")
	suite.ErrorIs(err, domain.ErrInternalServerError)
}
