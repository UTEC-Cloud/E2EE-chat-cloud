package e2ee_api

import (
	"encoding/json"

	x3dh_core "tux.tech/x3dh/core"
)

type InboundMessage struct {
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
}

type RequestUserBundle struct {
	UserID string `json:"user_id"`
}

type RequestUploadBundle struct {
	UserID string                     `json:"user_id"`
	Bundle x3dh_core.X3DHClientBundle `json:"bundle"`
}

type RequestUserIsRegistered struct {
	UserID string `json:"user_id"`
}

type RequestSendMsg struct {
	RecipientID string                   `json:"recipient_id"`
	MessageData x3dh_core.InitialMessage `json:"message"`
}

type RequestReceiveMsg struct{}

type OutboundMessage struct {
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
}

type ResponseUserBundle struct {
	Success bool                    `json:"success"`
	Bundle  x3dh_core.X3DHKeyBundle `json:"bundle"`
}

type ResponseUploadBundle struct {
	Success bool `json:"success"`
}

type ResponseUserIsRegistered struct {
	Success bool `json:"success"`
}

type ResponseSendMsg struct {
	Success bool `json:"success"`
}

type ResponseReceiveMsg struct {
	SenderID    string                   `json:"sender_id"`
	MessageData x3dh_core.InitialMessage `json:"message"`
}
