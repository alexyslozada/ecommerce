package model

import "encoding/json"

type PayPalRequestValidator struct {
	// Headers
	AuthAlgo         string `json:"auth_algo"`
	CertURL          string `json:"cert_url"`
	TransmissionID   string `json:"transmission_id"`
	TransmissionSig  string `json:"transmission_sig"`
	TransmissionTime string `json:"transmission_time"`

	// Body
	WebhookID    string          `json:"webhook_id"`
	WebhookEvent json.RawMessage `json:"webhook_event"`
}

type PayPalRequestData struct {
	EventType string `json:"event_type"`
	ID        string `json:"id"`
	Resource  struct {
		ID       string `json:"id"`
		Status   string `json:"status"`
		CustomID string `json:"custom_id"`
		Amount   struct {
			Value string `json:"value"`
		} `json:"amount"`
	} `json:"resource"`
}
