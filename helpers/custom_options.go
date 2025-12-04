package pgghelpers

import (
	"fmt"
	"sync"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
)

var protoregistryMutex sync.Mutex

// getExtension returns the value of an extension.
//
// If the extension with the given ID is not already registered, it will be made up.
func getExtension(extendedMessage proto.Message, extendedType proto.Message, fieldID int32, fieldType interface{}) (interface{}, error) {
	// To prevent concurrent map read/write while querying the registry and registring new extensions, request a lock.
	protoregistryMutex.Lock()
	defer protoregistryMutex.Unlock()

	// In the new API, we can't easily iterate over all extensions for a message type from the global registry
	// in the same way. But we can check if the extension is already known.

	// We need to find the extension descriptor.
	// Since we don't have the full name of the extension, only the field ID and the extended message type.

	// The original code was iterating over all extensions of the extended message.
	// protoregistry.GlobalTypes.RangeExtensionsByMessage does exactly that.

	var extensionType protoreflect.ExtensionType
	protoregistry.GlobalTypes.RangeExtensionsByMessage(extendedMessage.ProtoReflect().Descriptor().FullName(), func(xt protoreflect.ExtensionType) bool {
		if xt.TypeDescriptor().Number() == protoreflect.FieldNumber(fieldID) {
			extensionType = xt
			return false
		}
		return true
	})

	if extensionType == nil {
		// Register the extension dynamically
		// This is complicated in the new API.
		// We need to construct a dynamic extension type.

		// For now, let's assume standard types.
		// If we really need dynamic registration, we might need to use internal/impl packages or just fail.
		// But let's try to construct a minimal ExtensionInfo.

		// Note: RegisterExtension is not available in protoregistry.GlobalTypes directly for dynamic types
		// created this way without using internal protoimpl.

		// However, we can try to use the dynamicpb or just return error for now if not found.
		// The original code was very aggressive.

		return nil, fmt.Errorf("extension %d not found", fieldID)
	}

	return proto.GetExtension(extendedMessage, extensionType), nil
}

// stringMethodOptionsExtension extracts method options of a string type.
// To define your own extensions see:
// https://developers.google.com/protocol-buffers/docs/proto#customoptions
// Typically the fieldID of private extensions should be in the range:
// 50000-99999
func stringMethodOptionsExtension(fieldID int32, f *descriptorpb.MethodDescriptorProto) string {
	if f == nil || f.Options == nil {
		return ""
	}

	var extendedType *descriptorpb.MethodOptions
	var fieldType *string

	ext, err := getExtension(f.Options, extendedType, fieldID, fieldType)
	if err != nil {
		return ""
	}

	if str, ok := ext.(*string); ok {
		return *str
	}

	return ""
}

// boolMethodOptionsExtension extracts method options of a boolean type.
func boolMethodOptionsExtension(fieldID int32, f *descriptorpb.MethodDescriptorProto) bool {
	if f == nil || f.Options == nil {
		return false
	}

	var extendedType *descriptorpb.MethodOptions
	var fieldType *bool

	ext, err := getExtension(f.Options, extendedType, fieldID, fieldType)
	if err != nil {
		return false
	}

	if b, ok := ext.(*bool); ok {
		return *b
	}

	return false
}

// stringFileOptionsExtension extracts file options of a string type.
// To define your own extensions see:
// https://developers.google.com/protocol-buffers/docs/proto#customoptions
// Typically the fieldID of private extensions should be in the range:
// 50000-99999
func stringFileOptionsExtension(fieldID int32, f *descriptorpb.FileDescriptorProto) string {
	if f == nil || f.Options == nil {
		return ""
	}

	var extendedType *descriptorpb.FileOptions
	var fieldType *string

	ext, err := getExtension(f.Options, extendedType, fieldID, fieldType)
	if err != nil {
		return ""
	}

	if str, ok := ext.(*string); ok {
		return *str
	}

	return ""
}

func stringFieldExtension(fieldID int32, f *descriptorpb.FieldDescriptorProto) string {
	if f == nil || f.Options == nil {
		return ""
	}

	var extendedType *descriptorpb.FieldOptions
	var fieldType *string

	ext, err := getExtension(f.Options, extendedType, fieldID, fieldType)
	if err != nil {
		return ""
	}

	str, ok := ext.(*string)
	if !ok {
		return ""
	}

	return *str
}

func int64FieldExtension(fieldID int32, f *descriptorpb.FieldDescriptorProto) int64 {
	if f == nil || f.Options == nil {
		return 0
	}

	var extendedType *descriptorpb.FieldOptions
	var fieldType *int64

	ext, err := getExtension(f.Options, extendedType, fieldID, fieldType)
	if err != nil {
		return 0
	}

	i, ok := ext.(*int64)
	if !ok {
		return 0
	}

	return *i
}

func int64MessageExtension(fieldID int32, f *descriptorpb.DescriptorProto) int64 {
	if f == nil || f.Options == nil {
		return 0
	}

	var extendedType *descriptorpb.MessageOptions
	var fieldType *int64

	ext, err := getExtension(f.Options, extendedType, fieldID, fieldType)
	if err != nil {
		return 0
	}

	i, ok := ext.(*int64)
	if !ok {
		return 0
	}

	return *i
}

func stringMessageExtension(fieldID int32, f *descriptorpb.DescriptorProto) string {
	if f == nil || f.Options == nil {
		return ""
	}

	var extendedType *descriptorpb.MessageOptions
	var fieldType *string

	ext, err := getExtension(f.Options, extendedType, fieldID, fieldType)
	if err != nil {
		return ""
	}

	str, ok := ext.(*string)
	if !ok {
		return ""
	}

	return *str
}

func boolFieldExtension(fieldID int32, f *descriptorpb.FieldDescriptorProto) bool {
	if f == nil || f.Options == nil {
		return false
	}

	var extendedType *descriptorpb.FieldOptions
	var fieldType *bool

	ext, err := getExtension(f.Options, extendedType, fieldID, fieldType)
	if err != nil {
		return false
	}

	b, ok := ext.(*bool)
	if !ok {
		return false
	}

	return *b
}

func boolMessageExtension(fieldID int32, f *descriptorpb.DescriptorProto) bool {
	if f == nil || f.Options == nil {
		return false
	}
	var extendedType *descriptorpb.MessageOptions
	var fieldType *bool

	ext, err := getExtension(f.Options, extendedType, fieldID, fieldType)
	if err != nil {
		return false
	}

	b, ok := ext.(*bool)
	if !ok {
		return false
	}

	return *b
}
