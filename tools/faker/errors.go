package faker

import (
	"fmt"
	"github.com/nebtex/hybrids/golang/hybrids"
)

//DecodeError contains the data to fully identify  an error produced by the decoder
type Error struct {
	// field full path
	Path string `json:"path,omitempty"`
	// field underlying type
	HybridType  hybrids.Types `json:"hybrid_type,omitempty"`
	OmniqlType  string        `json:"omniql_type,omitempty"`
	Application string        `json:"application,omitempty"`
	ErrorMsg    string        `json:"error,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("#Application: %s, #Path: %s, #HybridType: %s, #OmniqlType: %s, Error: %s", e.Application, e.Path, e.HybridType.String(), e.OmniqlType, e.ErrorMsg)
}
