package internal

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type ReflectInterface interface {
	ServiceReflectionFromMessage(msg proto.Message) (ServiceReflectInterface, error)
	ServiceReflectionFromName(name protoreflect.FullName) (ServiceReflectInterface, bool)
	MethodReflectionFromMessage(msg proto.Message) (MethodReflectInterface, error)
	AddService(name protoreflect.FullName)
	AddProtoExtension(*protoimpl.ExtensionInfo)
	Populate() error
}

type ReflectAccessor interface {
	Get(key *protoimpl.ExtensionInfo) (interface{}, bool)
	Has(key *protoimpl.ExtensionInfo) bool
}

type ServiceReflectInterface interface {
	ReflectAccessor
	Descriptor() protoreflect.ServiceDescriptor
	Methods() []MethodReflectInterface
}

type MethodReflectInterface interface {
	ReflectAccessor
	Descriptor() protoreflect.MethodDescriptor
	Request() MessageReflectInterface
	Response() MessageReflectInterface
	FromMessage(proto.Message) MessageReflectInterface
	IsRequest(proto.Message) bool
	IsResponse(proto.Message) bool
	IsValidMessage(proto.Message) bool
}

type MessageReflectInterface interface {
	ReflectAccessor
	Descriptor() protoreflect.MessageDescriptor
	Value() proto.Message
	WithValue(v proto.Message)
	Fields() []FieldReflectInterface
}

type FieldReflectInterface interface {
	ReflectAccessor
	Descriptor() protoreflect.FieldDescriptor
}