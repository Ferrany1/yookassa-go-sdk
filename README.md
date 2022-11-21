# Обёртка для API ЮКассы на go.

## Установка
`go get -u github.com/Ferrany1/yookassa-go-sdk`

## Пример c дженериками
Простая программа для создания платежа и проверки его статуса.
Данный способ вводит ограничение на метод платежа и подтверждение для клиента, но убирает необходимость приведения типов. 
Подойдет если на стороне сервиса будет реализован механизм определена фабрика для создания необходимых наборов методов платежей и подтверждений
```go
package main

import (
	"context"
	"time"

	"github.com/Ferrany1/yookassa-go-sdk-fed/pkg/adapter"
	"github.com/Ferrany1/yookassa-go-sdk-fed/pkg/model/amount"
	"github.com/Ferrany1/yookassa-go-sdk-fed/pkg/model/confirmation"
	"github.com/Ferrany1/yookassa-go-sdk-fed/pkg/model/payment"
	"github.com/Ferrany1/yookassa-go-sdk-fed/pkg/model/payment_method"
	"github.com/Ferrany1/yookassa-go-sdk-fed/pkg/model/status"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()
	kassa := adapter.NewKassaAdapter[payment_method.YooMoneyPaymentMethod, confirmation.RedirectConfirmation](
		&adapter.Config{
			ShopID: "ИДЕНТИФИКАТОР_МАГАЗИНА",
			Token:  "ПРИВАТНЫЙ_КЛЮЧ",
		},
	)

	// Создание и настройка платежа.
	request := payment.NewPaymentRequest[payment_method.YooMoneyPaymentMethod, confirmation.RedirectConfirmation]().
		WithAmount(10, amount.RubCurrency).
		WithDescription("тестовый заказ").
		WithSavePaymentMethod().
		WithCapture().
		WithPaymentMethodData(payment_method.NewYooMoneyPaymentMethod()).
		WithConfirmation(confirmation.NewRedirectConfirmation().WithLocale(confirmation.RussianLocale))
	// Отправка конфига и получение экземпляра платежа в ответ.
	paymentResponse, err := kassa.PaymentRequest(ctx, request)
	if err != nil {
		panic(err)
	}

	// Вывод ссылки для оплаты.
	println(paymentResponse.Confirmation.ConfirmationURL)

	paymentID := paymentResponse.ID
	// Ожидание оплаты и опрос сервера на предмет её совершения.
	// Опрос заканчивается, если истекло время оплаты.

	for i := 0; i < 100; i++ {
		paymentResponse, err = kassa.GetPayment(ctx, paymentID)
		if err != nil {
			panic(err)
		}

		if paymentResponse.Paid {
			println("Оплата прошла успешно.")
			break
		} else if paymentResponse.Status == status.Canceled {
			println("Пользователь отменил оплату...")
		}
		time.Sleep(time.Second * 10)
	}
}
```

## Пример c дженериками рекурентного платежа
```go
package main

import (
	"context"
	"time"

	"github.com/Ferrany1/yookassa-go-sdk-fed/pkg/adapter"
	"github.com/Ferrany1/yookassa-go-sdk-fed/pkg/model/amount"
	"github.com/Ferrany1/yookassa-go-sdk-fed/pkg/model/confirmation"
	"github.com/Ferrany1/yookassa-go-sdk-fed/pkg/model/payment"
	"github.com/Ferrany1/yookassa-go-sdk-fed/pkg/model/payment_method"
	"github.com/Ferrany1/yookassa-go-sdk-fed/pkg/model/status"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()
	kassa := adapter.NewKassaAdapter[payment_method.YooMoneyPaymentMethod, confirmation.RedirectConfirmation](
		&adapter.Config{
			ShopID: "ИДЕНТИФИКАТОР_МАГАЗИНА",
			Token:  "ПРИВАТНЫЙ_КЛЮЧ",
		},
	)

	// Создание и настройка платежа.
	request := payment.NewPaymentRequest[payment_method.YooMoneyPaymentMethod, confirmation.RedirectConfirmation]().
		WithAmount(10, amount.RubCurrency).
		WithDescription("тестовый заказ").
		WithSavePaymentMethod().
		WithCapture().
		WithPaymentMethodData(payment_method.NewYooMoneyPaymentMethod()).
		WithConfirmation(confirmation.NewRedirectConfirmation().WithLocale(confirmation.RussianLocale))
	// Отправка конфига и получение экземпляра платежа в ответ.
	paymentResponse, err := kassa.PaymentRequest(ctx, request)
	if err != nil {
		panic(err)
	}

	// Вывод ссылки для оплаты.
	println(paymentResponse.Confirmation.ConfirmationURL)

	paymentID := paymentResponse.ID
	paymentMethodID := ""

	for i := 0; i < 100; i++ {
		paymentResponse, err = kassa.GetPayment(ctx, paymentID)
		if err != nil {
			panic(err)
		}

		if paymentResponse.Paid {
			paymentMethodID = paymentResponse.PaymentMethod.ID
			println("Оплата прошла успешно 1.")
			break
		} else if paymentResponse.Status == status.Canceled {
			println("Пользователь отменил оплату...")
		}
		time.Sleep(time.Second * 10)
	}

	request = payment.NewPaymentRequest[payment_method.YooMoneyPaymentMethod, confirmation.RedirectConfirmation]().
		WithAmount(10, amount.RubCurrency).
		WithDescription("тестовый заказ 1.1").
		WithSavePaymentMethod().
		WithCapture().
		WithPaymentMethodID(paymentMethodID)

	paymentResponse, err = kassa.PaymentRequest(context.Background(), request)
	if err != nil {
		panic(err)
	}
	paymentID = paymentResponse.ID

	for i := 0; i < 100; i++ {
		paymentResponse, err = kassa.GetPayment(ctx, paymentID)
		if err != nil {
			panic(err)
		}

		if paymentResponse.Paid {
			paymentMethodID = paymentResponse.PaymentMethod.ID
			println("Оплата прошла успешно 1.1")
			break
		} else if paymentResponse.Status == status.Canceled {
			println("Пользователь отменил оплату...")
		}
		time.Sleep(time.Second * 10)
	}
}
```

## Пример без дженериков
Простая программа для создания платежа и проверки его статуса.
```go
package main

import (
	"context"
	"time"

	"github.com/Ferrany1/yookassa-go-sdk-fed/pkg/adapter"
	"github.com/Ferrany1/yookassa-go-sdk-fed/pkg/model/amount"
	"github.com/Ferrany1/yookassa-go-sdk-fed/pkg/model/confirmation"
	"github.com/Ferrany1/yookassa-go-sdk-fed/pkg/model/payment"
	"github.com/Ferrany1/yookassa-go-sdk-fed/pkg/model/status"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()
	kassa := adapter.NewRawKassaAdapter(
		&adapter.Config{
			ShopID: "ИДЕНТИФИКАТОР_МАГАЗИНА",
			Token:  "ПРИВАТНЫЙ_КЛЮЧ",
		},
	)

	// Создание и настройка платежа.
	request := &payment.RawPaymentRequest{
		RawPaymentPart: payment.RawPaymentPart{
			Amount: amount.Amount{
				Value:    "10.00",
				Currency: "RUB",
			},
			Description:       "тестовый заказ",
			SavePaymentMethod: true,
			Capture:           true,
		},
		Confirmation: &map[string]interface{}{
			"type":       string(confirmation.RedirectConfirmationType),
			"locale":     string(confirmation.RussianLocale),
			"return_url": "",
		},
	}

	// Отправка конфига и получение экземпляра платежа в ответ.
	paymentResponse, err := kassa.PaymentRequest(ctx, request)
	if err != nil {
		panic(err)
	}

	// Вывод ссылки для оплаты.
	println(paymentResponse.Confirmation.ConfirmationURL)

	paymentID := paymentResponse.ID
	// Ожидание оплаты и опрос сервера на предмет её совершения.
	// Опрос заканчивается, если истекло время оплаты.

	for i := 0; i < 100; i++ {
		paymentResponse, err = kassa.GetPayment(ctx, paymentID)
		if err != nil {
			panic(err)
		}

		if paymentResponse.Paid {
			println("Оплата прошла успешно.")
			break
		} else if paymentResponse.Status == status.Canceled {
			println("Пользователь отменил оплату...")
		}
		time.Sleep(time.Second * 10)
	}
}
```