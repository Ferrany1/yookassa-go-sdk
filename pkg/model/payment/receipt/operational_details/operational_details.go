package operational_details //nolint:revive // better package readability

import "time"

type OperationalDetail struct {
	OperationID string    `json:"operation_id"`
	Value       string    `json:"value"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewOperationalDetail() *OperationalDetail {
	return &OperationalDetail{}
}

func (d *OperationalDetail) WithOperationID(operationID string) *OperationalDetail {
	d.OperationID = operationID
	return d
}

func (d *OperationalDetail) WithValue(value string) *OperationalDetail {
	d.Value = value
	return d
}

func (d *OperationalDetail) WithCreatedAt(createdAt time.Time) *OperationalDetail {
	d.CreatedAt = createdAt
	return d
}

type OperationalDetails []*OperationalDetail
