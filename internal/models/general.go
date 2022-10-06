package models

type QueryParams struct {
	Offset uint32 `form:"offset"`
	Limit  uint32 `form:"limit"`
	Order  string `form:"order"`
}

type BodyAssignBot struct {
	OrderId string `json:"order_id" validate:"required"`
}
