// Code generated by zanzibar
// @generated
// AUTOGENERATED FILE: easyjson marshaller/unmarshallers.

package bar

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	bar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/bar/bar"
	foo "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/foo/foo"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonCa6a3ed2DecodeGithubComUberZanzibarExamplesExampleGatewayBuildEndpointsBar(in *jlexer.Lexer, out *TooManyArgsHTTPRequest) {
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
		case "request":
			if in.IsNull() {
				in.Skip()
				out.Request = nil
			} else {
				if out.Request == nil {
					out.Request = new(bar.BarRequest)
				}
				(*out.Request).UnmarshalEasyJSON(in)
			}
		case "foo":
			if in.IsNull() {
				in.Skip()
				out.Foo = nil
			} else {
				if out.Foo == nil {
					out.Foo = new(foo.FooStruct)
				}
				(*out.Foo).UnmarshalEasyJSON(in)
			}
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
func easyjsonCa6a3ed2EncodeGithubComUberZanzibarExamplesExampleGatewayBuildEndpointsBar(out *jwriter.Writer, in TooManyArgsHTTPRequest) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"request\":")
	if in.Request == nil {
		out.RawString("null")
	} else {
		(*in.Request).MarshalEasyJSON(out)
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"foo\":")
	if in.Foo == nil {
		out.RawString("null")
	} else {
		(*in.Foo).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TooManyArgsHTTPRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCa6a3ed2EncodeGithubComUberZanzibarExamplesExampleGatewayBuildEndpointsBar(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TooManyArgsHTTPRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCa6a3ed2EncodeGithubComUberZanzibarExamplesExampleGatewayBuildEndpointsBar(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TooManyArgsHTTPRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCa6a3ed2DecodeGithubComUberZanzibarExamplesExampleGatewayBuildEndpointsBar(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TooManyArgsHTTPRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCa6a3ed2DecodeGithubComUberZanzibarExamplesExampleGatewayBuildEndpointsBar(l, v)
}
func easyjsonCa6a3ed2DecodeGithubComUberZanzibarExamplesExampleGatewayBuildEndpointsBar1(in *jlexer.Lexer, out *NormalHTTPRequest) {
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
		case "request":
			if in.IsNull() {
				in.Skip()
				out.Request = nil
			} else {
				if out.Request == nil {
					out.Request = new(bar.BarRequest)
				}
				(*out.Request).UnmarshalEasyJSON(in)
			}
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
func easyjsonCa6a3ed2EncodeGithubComUberZanzibarExamplesExampleGatewayBuildEndpointsBar1(out *jwriter.Writer, in NormalHTTPRequest) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"request\":")
	if in.Request == nil {
		out.RawString("null")
	} else {
		(*in.Request).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NormalHTTPRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCa6a3ed2EncodeGithubComUberZanzibarExamplesExampleGatewayBuildEndpointsBar1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NormalHTTPRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCa6a3ed2EncodeGithubComUberZanzibarExamplesExampleGatewayBuildEndpointsBar1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NormalHTTPRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCa6a3ed2DecodeGithubComUberZanzibarExamplesExampleGatewayBuildEndpointsBar1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NormalHTTPRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCa6a3ed2DecodeGithubComUberZanzibarExamplesExampleGatewayBuildEndpointsBar1(l, v)
}
func easyjsonCa6a3ed2DecodeGithubComUberZanzibarExamplesExampleGatewayBuildEndpointsBar2(in *jlexer.Lexer, out *ArgNotStructHTTPRequest) {
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
		case "request":
			out.Request = string(in.String())
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
func easyjsonCa6a3ed2EncodeGithubComUberZanzibarExamplesExampleGatewayBuildEndpointsBar2(out *jwriter.Writer, in ArgNotStructHTTPRequest) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"request\":")
	out.String(string(in.Request))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ArgNotStructHTTPRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCa6a3ed2EncodeGithubComUberZanzibarExamplesExampleGatewayBuildEndpointsBar2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ArgNotStructHTTPRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCa6a3ed2EncodeGithubComUberZanzibarExamplesExampleGatewayBuildEndpointsBar2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ArgNotStructHTTPRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCa6a3ed2DecodeGithubComUberZanzibarExamplesExampleGatewayBuildEndpointsBar2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ArgNotStructHTTPRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCa6a3ed2DecodeGithubComUberZanzibarExamplesExampleGatewayBuildEndpointsBar2(l, v)
}
