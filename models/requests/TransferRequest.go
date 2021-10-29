package requests

type TransferRequest struct {
	ReceiverID int     `json:"receiver_id"`
	Amount     int     `json:"amount"`
	Message    *string `json:"message"`
}
