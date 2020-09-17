// Code generated by 'github.com/traefik/yaegi/extract encoding/xml'. DO NOT EDIT.

// +build go1.14,!go1.15

package stdlib

import (
	"encoding/xml"
	"go/constant"
	"go/token"
	"reflect"
)

func init() {
	Symbols["encoding/xml"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"CopyToken":       reflect.ValueOf(xml.CopyToken),
		"Escape":          reflect.ValueOf(xml.Escape),
		"EscapeText":      reflect.ValueOf(xml.EscapeText),
		"HTMLAutoClose":   reflect.ValueOf(&xml.HTMLAutoClose).Elem(),
		"HTMLEntity":      reflect.ValueOf(&xml.HTMLEntity).Elem(),
		"Header":          reflect.ValueOf(constant.MakeFromLiteral("\"<?xml version=\\\"1.0\\\" encoding=\\\"UTF-8\\\"?>\\n\"", token.STRING, 0)),
		"Marshal":         reflect.ValueOf(xml.Marshal),
		"MarshalIndent":   reflect.ValueOf(xml.MarshalIndent),
		"NewDecoder":      reflect.ValueOf(xml.NewDecoder),
		"NewEncoder":      reflect.ValueOf(xml.NewEncoder),
		"NewTokenDecoder": reflect.ValueOf(xml.NewTokenDecoder),
		"Unmarshal":       reflect.ValueOf(xml.Unmarshal),

		// type definitions
		"Attr":                 reflect.ValueOf((*xml.Attr)(nil)),
		"CharData":             reflect.ValueOf((*xml.CharData)(nil)),
		"Comment":              reflect.ValueOf((*xml.Comment)(nil)),
		"Decoder":              reflect.ValueOf((*xml.Decoder)(nil)),
		"Directive":            reflect.ValueOf((*xml.Directive)(nil)),
		"Encoder":              reflect.ValueOf((*xml.Encoder)(nil)),
		"EndElement":           reflect.ValueOf((*xml.EndElement)(nil)),
		"Marshaler":            reflect.ValueOf((*xml.Marshaler)(nil)),
		"MarshalerAttr":        reflect.ValueOf((*xml.MarshalerAttr)(nil)),
		"Name":                 reflect.ValueOf((*xml.Name)(nil)),
		"ProcInst":             reflect.ValueOf((*xml.ProcInst)(nil)),
		"StartElement":         reflect.ValueOf((*xml.StartElement)(nil)),
		"SyntaxError":          reflect.ValueOf((*xml.SyntaxError)(nil)),
		"TagPathError":         reflect.ValueOf((*xml.TagPathError)(nil)),
		"Token":                reflect.ValueOf((*xml.Token)(nil)),
		"TokenReader":          reflect.ValueOf((*xml.TokenReader)(nil)),
		"UnmarshalError":       reflect.ValueOf((*xml.UnmarshalError)(nil)),
		"Unmarshaler":          reflect.ValueOf((*xml.Unmarshaler)(nil)),
		"UnmarshalerAttr":      reflect.ValueOf((*xml.UnmarshalerAttr)(nil)),
		"UnsupportedTypeError": reflect.ValueOf((*xml.UnsupportedTypeError)(nil)),

		// interface wrapper definitions
		"_Marshaler":       reflect.ValueOf((*_encoding_xml_Marshaler)(nil)),
		"_MarshalerAttr":   reflect.ValueOf((*_encoding_xml_MarshalerAttr)(nil)),
		"_Token":           reflect.ValueOf((*_encoding_xml_Token)(nil)),
		"_TokenReader":     reflect.ValueOf((*_encoding_xml_TokenReader)(nil)),
		"_Unmarshaler":     reflect.ValueOf((*_encoding_xml_Unmarshaler)(nil)),
		"_UnmarshalerAttr": reflect.ValueOf((*_encoding_xml_UnmarshalerAttr)(nil)),
	}
}

// _encoding_xml_Marshaler is an interface wrapper for Marshaler type
type _encoding_xml_Marshaler struct {
	WMarshalXML func(e *xml.Encoder, start xml.StartElement) error
}

func (W _encoding_xml_Marshaler) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return W.WMarshalXML(e, start)
}

// _encoding_xml_MarshalerAttr is an interface wrapper for MarshalerAttr type
type _encoding_xml_MarshalerAttr struct {
	WMarshalXMLAttr func(name xml.Name) (xml.Attr, error)
}

func (W _encoding_xml_MarshalerAttr) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return W.WMarshalXMLAttr(name)
}

// _encoding_xml_Token is an interface wrapper for Token type
type _encoding_xml_Token struct {
}

// _encoding_xml_TokenReader is an interface wrapper for TokenReader type
type _encoding_xml_TokenReader struct {
	WToken func() (xml.Token, error)
}

func (W _encoding_xml_TokenReader) Token() (xml.Token, error) { return W.WToken() }

// _encoding_xml_Unmarshaler is an interface wrapper for Unmarshaler type
type _encoding_xml_Unmarshaler struct {
	WUnmarshalXML func(d *xml.Decoder, start xml.StartElement) error
}

func (W _encoding_xml_Unmarshaler) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return W.WUnmarshalXML(d, start)
}

// _encoding_xml_UnmarshalerAttr is an interface wrapper for UnmarshalerAttr type
type _encoding_xml_UnmarshalerAttr struct {
	WUnmarshalXMLAttr func(attr xml.Attr) error
}

func (W _encoding_xml_UnmarshalerAttr) UnmarshalXMLAttr(attr xml.Attr) error {
	return W.WUnmarshalXMLAttr(attr)
}
