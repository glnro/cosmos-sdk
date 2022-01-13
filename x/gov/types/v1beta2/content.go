package v1beta2

import (
	"fmt"

	"github.com/gogo/protobuf/proto"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

// Copied over from /x/gov/types/keys.go to avoid circular imports
const (
	moduleName = "gov"

	routerKey = moduleName
)

// NewLegacyContent creates a new MsgExecLegacyContent from a legacy Content
// interface.
func NewLegacyContent(content v1beta1.Content, authority string) (*MsgExecLegacyContent, error) {
	msg, ok := content.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("%T does not implement proto.Message", content)
	}

	any, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		return nil, err
	}

	return NewMsgExecLegacyContent(any, authority), nil
}

// LegacyContentFromMessage extracts the legacy Content interface from a
// MsgExecLegacyContent.
func LegacyContentFromMessage(msg *MsgExecLegacyContent) v1beta1.Content {
	content, ok := msg.Content.GetCachedValue().(v1beta1.Content)
	if !ok {
		return nil
	}
	return content
}
