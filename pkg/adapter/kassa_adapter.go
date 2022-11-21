package adapter

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/confirmation"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment_method"
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment"
)

type KassaAdapter[
	PaymentMethod payment_method.GenericPaymentMethod,
	Confirmation confirmation.GenericConfirmation,
] struct {
	config *Config
	// HTTP клиент для обработки запросов.
	client *http.Client
	// Адрес, по которому требуется выполнять запросы.
	endpoint string
}

// KassaAdapter создаёт экземпляр структуры Kassa.
func NewKassaAdapter[
	PaymentMethod payment_method.GenericPaymentMethod,
	Confirmation confirmation.GenericConfirmation,
](config *Config, opts ...KassaAdapterOption[PaymentMethod, Confirmation]) *KassaAdapter[PaymentMethod, Confirmation] {
	kassa := &KassaAdapter[PaymentMethod, Confirmation]{
		config:   config,
		client:   http.DefaultClient,
		endpoint: "https://api.yookassa.ru/v3/",
	}
	for i := range opts {
		opts[i](kassa)
	}
	return kassa
}

type KassaAdapterOption[
	PaymentMethod payment_method.GenericPaymentMethod,
	Confirmation confirmation.GenericConfirmation,
] func(*KassaAdapter[PaymentMethod, Confirmation])

func WithKassaHTTPClient[
	PaymentMethod payment_method.GenericPaymentMethod,
	Confirmation confirmation.GenericConfirmation,
](httpClient *http.Client) KassaAdapterOption[PaymentMethod, Confirmation] {
	return func(kassa *KassaAdapter[PaymentMethod, Confirmation]) {
		kassa.client = httpClient
	}
}

// Ping отправляет тестовый запрос для проверки соединения.
func (k KassaAdapter[PaymentMethod, Confirmation]) Ping(ctx context.Context) error {
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

// PaymentRequest отправляет payment.PaymentRequest на сервера ЮКассы
// и получает готовый экземпляр payment.PaymentResponse в ответ.
func (k KassaAdapter[PaymentMethod, Confirmation]) PaymentRequest(
	ctx context.Context,
	paymentRequest *payment.PaymentRequest[PaymentMethod, Confirmation],
) (*payment.PaymentResponse[PaymentMethod, Confirmation], error) {
	paymentBytes, err := json.Marshal(paymentRequest)
	if err != nil {
		return nil, errors.Wrap(err, "encoding payment request")
	}

	result := &payment.PaymentResponse[PaymentMethod, Confirmation]{}
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

// GetPayment получает объект payment.PaymentResponse по ID.
func (k KassaAdapter[PaymentMethod, Confirmation]) GetPayment(
	ctx context.Context,
	id string,
) (*payment.PaymentResponse[PaymentMethod, Confirmation], error) {
	result := &payment.PaymentResponse[PaymentMethod, Confirmation]{}
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
func (k KassaAdapter[PaymentMethod, Confirmation]) AcceptSpending(ctx context.Context, id string) error {
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
func (k KassaAdapter[PaymentMethod, Confirmation]) DeclineSpending(ctx context.Context, id string) error {
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
