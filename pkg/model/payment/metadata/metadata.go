package metadata

// Metadata Любые дополнительные данные, которые нужны вам для работы (например, номер заказа).
// Передаются в виде набора пар «ключ-значение» и возвращаются в ответе от ЮKassa.
// Ограничения: максимум 16 ключей, имя ключа не больше 32 символов,
// значение ключа не больше 512 символов, тип данных — строка в формате UTF-8.
type Metadata map[string]string

func NewMetadata(metadata map[string]string) *Metadata {
	m := Metadata(metadata)
	return &m
}
