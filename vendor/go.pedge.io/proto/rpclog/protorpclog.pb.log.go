// Code generated by protoc-gen-protolog
// source: rpclog/protorpclog.proto
// DO NOT EDIT!

package protorpclog

import "go.pedge.io/protolog"

func init() {
	protolog.Register("protorpclog.Call", protolog.MessageType_MESSAGE_TYPE_EVENT, func() protolog.Message { return &Call{} })
}

func (m *Call) ProtologName() string {
	return "protorpclog.Call"
}