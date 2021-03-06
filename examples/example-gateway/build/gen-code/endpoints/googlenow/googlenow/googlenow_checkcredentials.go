// Code generated by thriftrw v1.3.0
// @generated

package googlenow

import (
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type GoogleNow_CheckCredentials_Args struct{}

func (v *GoogleNow_CheckCredentials_Args) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *GoogleNow_CheckCredentials_Args) FromWire(w wire.Value) error {
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}
	return nil
}

func (v *GoogleNow_CheckCredentials_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [0]string
	i := 0
	return fmt.Sprintf("GoogleNow_CheckCredentials_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *GoogleNow_CheckCredentials_Args) Equals(rhs *GoogleNow_CheckCredentials_Args) bool {
	return true
}

func (v *GoogleNow_CheckCredentials_Args) MethodName() string {
	return "checkCredentials"
}

func (v *GoogleNow_CheckCredentials_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var GoogleNow_CheckCredentials_Helper = struct {
	Args           func() *GoogleNow_CheckCredentials_Args
	IsException    func(error) bool
	WrapResponse   func(error) (*GoogleNow_CheckCredentials_Result, error)
	UnwrapResponse func(*GoogleNow_CheckCredentials_Result) error
}{}

func init() {
	GoogleNow_CheckCredentials_Helper.Args = func() *GoogleNow_CheckCredentials_Args {
		return &GoogleNow_CheckCredentials_Args{}
	}
	GoogleNow_CheckCredentials_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	GoogleNow_CheckCredentials_Helper.WrapResponse = func(err error) (*GoogleNow_CheckCredentials_Result, error) {
		if err == nil {
			return &GoogleNow_CheckCredentials_Result{}, nil
		}
		return nil, err
	}
	GoogleNow_CheckCredentials_Helper.UnwrapResponse = func(result *GoogleNow_CheckCredentials_Result) (err error) {
		return
	}
}

type GoogleNow_CheckCredentials_Result struct{}

func (v *GoogleNow_CheckCredentials_Result) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *GoogleNow_CheckCredentials_Result) FromWire(w wire.Value) error {
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}
	return nil
}

func (v *GoogleNow_CheckCredentials_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [0]string
	i := 0
	return fmt.Sprintf("GoogleNow_CheckCredentials_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *GoogleNow_CheckCredentials_Result) Equals(rhs *GoogleNow_CheckCredentials_Result) bool {
	return true
}

func (v *GoogleNow_CheckCredentials_Result) MethodName() string {
	return "checkCredentials"
}

func (v *GoogleNow_CheckCredentials_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
