package models

type QueryParams struct {
	Offset uint32 `form:"offset"`
	Limit  uint32 `form:"limit"`
	Order  string `form:"order"`
}

type BodyAssignBot struct {
	// The order id of delivery pending to assign
	// example: "QgDz4iJONRfbHGSvqmSl"
	// required: true
	OrderId string `json:"order_id" validate:"required"`
} // @name BodyAssignBot

type ErrorResponse struct {
	Error string `json:"error"`
} // @name ErrorResponse

type CreateBotResponse struct {
	Message string `json:"message"`
	Bot     Bot    `json:"bot"`
} // @name CreateBotResponse

type CreateBotRequest struct {
	// Location coordinates of the bot
	Location Location `json:"location"`
	// ZoneId of the zone where the bot is located
	ZoneId string `json:"zone_id"`
} // @name CreateBotRequest

type GetBotsByZoneResponse struct {
	Total       int    `json:"total"`
	PreviousUrl string `json:"previous_url"`
	NextUrl     string `json:"next_url"`
	Results     []Bot  `json:"results"`
} // @name GetBotsByZoneResponse

type MessageNotFound struct {
	Message string `json:"message"`
} // @name MessageNotFound

type AssignBotToOrderResponse struct {
	Message string   `json:"message"`
	Bot     Bot      `json:"bot"`
	Order   Delivery `json:"order"`
} // @name AssignBotToOrderResponse

type CreateDeliveryRequest struct {
	Pickup  Pickup  `json:"pickup"`
	Dropoff Dropoff `json:"dropoff"`
	ZoneId  string  `json:"zone_id" validate:"required"`
} // @name CreateDeliveryRequest

type CreateDeliveryResponse struct {
	Message  string   `json:"message"`
	Delivery Delivery `json:"delivery"`
} // @name CreateDeliveryResponse

type GetDeliveriesResponse struct {
	Total       int        `json:"total"`
	PreviousUrl string     `json:"previous_url"`
	NextUrl     string     `json:"next_url"`
	Results     []Delivery `json:"results"`
} // @name GetDeliveriesResponse

type AssignBotsToAllPendingOrdersResponse struct {
	Message          string     `json:"message"`
	OrdersUnassigned []Delivery `json:"orders_unassigned"`
} // @name AssignBotsToAllPendingOrdersResponse
