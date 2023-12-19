package gokalkan

import "github.com/gokalkan/gokalkan/ckalkan"

func (cli *Client) SignWSSE2(x509, xmlData, id string, flags ckalkan.Flag) (string, error) {
	soapEnvelope := WrapWithWSSESoapEnvelope(xmlData, id)
	return cli.kc.SignWSSE(soapEnvelope, x509, flags, id)
}
