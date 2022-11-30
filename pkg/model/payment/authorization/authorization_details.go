package authorization

// AuthorizationDetails Данные об авторизации платежа.
type AuthorizationDetails struct {
	// Retrieval Reference Number — уникальный идентификатор транзакции в системе эмитента.
	// Используется при оплате банковской картой.
	RRN string `json:"rrn,omitempty"`
	// Код авторизации банковской карты.
	// Выдается эмитентом и подтверждает проведение авторизации.
	AuthCode string `json:"auth_code,omitempty"`
}
