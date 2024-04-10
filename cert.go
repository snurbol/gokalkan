package gokalkan

// GetCertificateFromStore получает открытую часть ЭЦП(сертификат) из хранилища
func (cli *Client) GetCertificateFromStore(alias string) (string, error) {
	cert, err := cli.kc.X509ExportCertificateFromStore(alias)
	if err != nil {
		return "", err
	}

	return cert, nil
}
