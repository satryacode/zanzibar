// Code generated by zanzibar
// @generated
// AUTOGENERATED FILE: easyjson marshaller/unmarshallers.

package googlenowClient

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson5cf1a88aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildClientsGooglenow(in *jlexer.Lexer, out *AddCredentialsHTTPRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "authCode":
			out.AuthCode = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson5cf1a88aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildClientsGooglenow(out *jwriter.Writer, in AddCredentialsHTTPRequest) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"authCode\":")
	out.String(string(in.AuthCode))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AddCredentialsHTTPRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5cf1a88aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildClientsGooglenow(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AddCredentialsHTTPRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5cf1a88aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildClientsGooglenow(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AddCredentialsHTTPRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5cf1a88aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildClientsGooglenow(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AddCredentialsHTTPRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5cf1a88aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildClientsGooglenow(l, v)
}
