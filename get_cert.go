package gokalkan

import "github.com/snurbol/gokalkan/ckalkan"

// GetCertFromCMS обеспечивает получение сертификата из CMS.
func (cli *Client) GetCertFromCMS(cmsInBase64 string, signID int) (string, error) {
	flags := ckalkan.FlagInBase64 | ckalkan.FlagNoCheckCertTime
	return cli.kc.GetCertFromCMS(cmsInBase64, signID, flags)
}
