package gokalkan

import (
	"encoding/base64"
	"encoding/xml"
	"strings"

	"github.com/snurbol/gokalkan/ckalkan"
)

// Sign подписывает данные и возвращает CMS с подписью.
func (cli *Client) Sign(data []byte, isDetached, withTSP bool) (signature []byte, err error) {
	dataB64 := base64.StdEncoding.EncodeToString(data)
	flags := ckalkan.FlagSignCMS | ckalkan.FlagInBase64 | ckalkan.FlagOutBase64

	if withTSP {
		flags |= ckalkan.FlagWithTimestamp
	}

	if isDetached {
		flags |= ckalkan.FlagDetachedData
	}

	signatureB64, err := cli.kc.SignData("", dataB64, "", flags)
	if err != nil {
		return nil, err
	}

	return base64.StdEncoding.DecodeString(signatureB64)
}

// SignXML подписывает данные в формате XML.
func (cli *Client) SignXML(xmlData string) (string, error) {
	return cli.kc.SignXML(xmlData, "", 0, "", "", "")
}

func (cli *Client) SignWSSE(xmlData, id string) (string, error) {
	soapEnvelope := WrapWithWSSESoapEnvelope(xmlData, id)
	return cli.kc.SignWSSE(soapEnvelope, "", 16777232, id)
}

const (
	xmlnsSOAP    = "http://schemas.xmlsoap.org/soap/envelope/"
	xmlnsSOAPENV = "http://schemas.xmlsoap.org/soap/envelope/"
	xmlnsBodyWsu = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
	xmlnsBodyNs2 = "http://bip.bee.kz/SyncChannel/v10/Types"

	replaceKey = "replace-this"
)

// soapEnvelope представляет soap:Envelope
type soapEnvelope struct {
	XMLName xml.Name `xml:"S:Envelope"`
	SOAP    string   `xml:"xmlns:S,attr"`
	SOAPENV string   `xml:"xmlns:SOAP-ENV,attr"`
	Body    soapBody `xml:"S:Body"`
}

// soapBody представляет soap:Body
type soapBody struct {
	ID      string `xml:"wsu:Id,attr"`
	Wsu     string `xml:"xmlns:wsu,attr"`
	Ns2     string `xml:"xmlns:ns2,attr"`
	Content string `xml:",chardata"`
}

// WrapWithWSSESoapEnvelope оборачивает XML документ в SOAP формат, а точнее записывает
// содержимое под тегом soap:Body
func WrapWithWSSESoapEnvelope(dataXML, id string) (result string) {
	envelope := soapEnvelope{
		SOAP:    xmlnsSOAP,
		SOAPENV: xmlnsSOAPENV,
		Body: soapBody{
			ID:      id,
			Wsu:     xmlnsBodyWsu,
			Ns2:     xmlnsBodyNs2,
			Content: replaceKey,
		},
	}

	b, err := xml.Marshal(envelope)
	if err != nil {
		return result
	}

	result = strings.Replace(string(b), replaceKey, dataXML, 1)

	return result
}
