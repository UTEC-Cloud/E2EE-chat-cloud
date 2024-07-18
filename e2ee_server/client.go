package main

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	api "tux.tech/e2ee/api"
)

type WsClient struct {
	username string
	server   *WsServer
	conn     *websocket.Conn
	send     chan []byte
}

func NewWsClient(username string, server *WsServer, conn *websocket.Conn) *WsClient {
	return &WsClient{
		username: username,
		server:   server,
		conn:     conn,
		send:     make(chan []byte),
	}
}

func (client *WsClient) WritePump() {
	for message := range client.send {
		client.conn.WriteMessage(websocket.TextMessage, message)
	}
}

func (client *WsClient) ReadPump() {
	for {
		mt, message, err := client.conn.ReadMessage()
		if err != nil || mt == websocket.CloseMessage {
			break // Exit loop
		}
		if mt == websocket.TextMessage {
			client.HandleMessage(message)
		}
	}
	client.Disconnect()
}

func (client *WsClient) HandleMessage(rawMessage []byte) {
	// Log
	fmt.Println("Received message from", client.username, ":", string(rawMessage))
	// Parse JSON
	message := &api.InboundMessage{}
	err := json.Unmarshal(rawMessage, message)
	if err != nil {
		return
	}
	switch message.Method {
	case "get_bundle":
		client.HandleGetUserBundle(message.Params)
	case "upload_bundle":
		client.HandleUploadBundle(message.Params)
	case "send_message":
		client.HandleSendMessage(message.Params)
	case "receive_message":
		client.HandleReceiveMessage(message.Params)
	case "user_is_registered":
		client.HandleUserIsRegistered(message.Params)
	}
}

func (client *WsClient) Disconnect() {
	client.server.UnsetClient(client)
	client.conn.Close()
	close(client.send) // Close write pump
}

func (client *WsClient) HandleGetUserBundle(rawParams json.RawMessage) {
	params := &api.RequestUserBundle{}
	err := json.Unmarshal(rawParams, params)
	if err != nil {
		return
	}

	// Get the bundle
	bundle, ok := client.server.X3DHServer.GetClientBundle(params.UserID)
	fmt.Println("User", client.username, "requested bundle for user", params.UserID, ":", ok)
	response := &api.ResponseUserBundle{
		Success: ok,
		Bundle:  bundle,
	}
	api_reply := &api.OutboundMessage{
		Method: "get_bundle",
		Params: response,
	}
	responseBytes, err := json.Marshal(api_reply)
	if err != nil {
		fmt.Println("Error marshalling fail response to get_bundle")
		return
	}
	client.send <- responseBytes
}

func (client *WsClient) HandleUploadBundle(rawParams json.RawMessage) {
	params := &api.RequestUploadBundle{}
	err := json.Unmarshal(rawParams, params)
	if err != nil {
		return
	}
	// Check if the user is the same
	if client.username != params.UserID {
		fmt.Println("User", client.username, "attempted to upload bundle for user", params.UserID)
		response := &api.ResponseUploadBundle{
			Success: false,
		}
		api_reply := &api.OutboundMessage{
			Method: "upload_bundle",
			Params: response,
		}
		responseBytes, err := json.Marshal(api_reply)
		if err != nil {
			fmt.Println("Error marshalling fail response to upload_bundle")
			return
		}
		client.send <- responseBytes
	}
	// Register
	client.server.X3DHServer.RegisterClient(params.UserID, params.Bundle)
	fmt.Println("User", client.username, "uploaded bundle for user", params.UserID)
	// Send response
	response := &api.ResponseUploadBundle{
		Success: true,
	}
	api_reply := &api.OutboundMessage{
		Method: "upload_bundle",
		Params: response,
	}
	responseBytes, err := json.Marshal(api_reply)
	if err != nil {
		fmt.Println("Error marshalling success response to upload_bundle")
		return
	}
	client.send <- responseBytes
}

func (client *WsClient) HandleSendMessage(rawParams json.RawMessage) {
	params := &api.RequestSendMsg{}
	err := json.Unmarshal(rawParams, params)
	if err != nil {
		return
	}
	// TODO: Implement
}

func (client *WsClient) HandleReceiveMessage(rawParams json.RawMessage) {
	params := &api.RequestReceiveMsg{}
	err := json.Unmarshal(rawParams, params)
	if err != nil {
		return
	}
	// TODO: Implement
}

func (client *WsClient) HandleUserIsRegistered(rawParams json.RawMessage) {
	params := &api.RequestUserIsRegistered{}
	err := json.Unmarshal(rawParams, params)
	if err != nil {
		return
	}
	// Check if the user is registered
	registered := client.server.X3DHServer.IsClientRegistered(params.UserID)
	fmt.Println("User", client.username, "checked if user", params.UserID, "is registered")
	// Send response
	response := &api.ResponseUserIsRegistered{
		Success: registered,
	}
	api_reply := &api.OutboundMessage{
		Method: "user_is_registered",
		Params: response,
	}
	responseBytes, err := json.Marshal(api_reply)
	if err != nil {
		fmt.Println("Error marshalling success response to user_is_registered")
		return
	}
	client.send <- responseBytes
}
