package adapter

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment"
)

type RawKassaAdapter struct {
	config *Config
	// HTTP клиент для обработки запросов.
	client *http.Client
	// Адрес, по которому требуется выполнять запросы.
	endpoint string
}

// RawKassaAdapter создаёт экземпляр структуры Kassa.
func NewRawKassaAdapter(config *Config, opts ...RawKassaAdapterOption) *RawKassaAdapter {
	kassa := &RawKassaAdapter{
		config:   config,
		client:   http.DefaultClient,
		endpoint: "https://api.yookassa.ru/v3/",
	}
	for i := range opts {
		opts[i](kassa)
	}
	return kassa
}

type RawKassaAdapterOption func(*RawKassaAdapter)

func WithRawKassaHTTPClient(httpClient *http.Client) RawKassaAdapterOption {
	return func(kassa *RawKassaAdapter) {
		kassa.client = httpClient
	}
}

// Ping отправляет тестовый запрос для проверки соединения.
func (k RawKassaAdapter) Ping(ctx context.Context) error {
	if err := doRequestWithDecode(
		ctx,
		http.MethodGet,
		k.endpoint+paymentsEndpoint,
		k.client,
		withBasicAuth(k.config.ShopID, k.config.Token),
	); err != nil {
		return errors.Wrap(err, "pinging yookassa")
	}
	return nil
}

// PaymentRequest отправляет payment.RawPaymentRequest на сервера ЮКассы
// и получает готовый экземпляр payment.RawPaymentResponse в ответ.
func (k RawKassaAdapter) PaymentRequest(
	ctx context.Context,
	paymentRequest *payment.RawPaymentRequest,
) (*payment.RawPaymentResponse, error) {
	paymentBytes, err := json.Marshal(paymentRequest)
	if err != nil {
		return nil, errors.Wrap(err, "encoding payment request")
	}

	result := &payment.RawPaymentResponse{}
	if err = doRequestWithDecode(
		ctx,
		http.MethodPost,
		k.endpoint+paymentsEndpoint,
		k.client,
		withBasicAuth(k.config.ShopID, k.config.Token),
		withRequestHeaderParams(
			map[string]string{
				"Content-Type":    "application/json",
				"Idempotence-Key": uuid.New().String(),
			},
		),
		withRequestBody(bytes.NewBuffer(paymentBytes)),
		withResponseTarget(result),
	); err != nil {
		return nil, errors.Wrap(err, "payment request")
	}

	if result.Type == errorType {
		return nil, errors.New(result.Description)
	}
	return result, nil
}

// GetPayment получает объект payment.RawPaymentResponse по ID.
func (k RawKassaAdapter) GetPayment(ctx context.Context, id string) (*payment.RawPaymentResponse, error) {
	result := &payment.RawPaymentResponse{}
	if err := doRequestWithDecode(
		ctx,
		http.MethodGet,
		k.endpoint+paymentsEndpoint+id,
		k.client,
		withBasicAuth(k.config.ShopID, k.config.Token),
		withResponseTarget(result),
	); err != nil {
		return nil, errors.Wrap(err, "getting payment")
	}

	if result.Type == errorType {
		return nil, errors.New(result.Description)
	}
	return result, nil
}

// AcceptSpending совершает списание средств, когда платёж перешёл в статус waiting_for_capture.
func (k RawKassaAdapter) AcceptSpending(ctx context.Context, id string) error {
	if err := doRequestWithDecode(
		ctx,
		http.MethodPost,
		k.endpoint+paymentsEndpoint+id+captureEndpoint,
		k.client,
		withBasicAuth(k.config.ShopID, k.config.Token),
		withRequestHeaderParams(
			map[string]string{
				"Content-Type":    "application/json",
				"Idempotence-Key": uuid.New().String(),
			},
		),
	); err != nil {
		return errors.Wrap(err, "accepting spending")
	}
	return nil
}

// DeclineSpending совершает отмену списания средств, когда платёж перешёл в статус waiting_for_capture.
func (k RawKassaAdapter) DeclineSpending(ctx context.Context, id string) error {
	if err := doRequestWithDecode(
		ctx,
		http.MethodPost,
		k.endpoint+paymentsEndpoint+id+declineEndpoint,
		k.client,
		withBasicAuth(k.config.ShopID, k.config.Token),
		withRequestHeaderParams(
			map[string]string{
				"Content-Type":    "application/json",
				"Idempotence-Key": uuid.New().String(),
			},
		),
	); err != nil {
		return errors.Wrap(err, "declining spending")
	}
	return nil
}
