package main

import (
	"fmt"
	"github.com/autom8ter/identify/saml"
	"github.com/autom8ter/identify/saml/driver"
)

func main() {
	example := NewExampleFunc()

	// Construct an AuthnRequest
	authRequest := saml.NewAuthorizationRequest(example)

	// Return a SAML AuthnRequest as a string
	saml, err := authRequest.GetSignedRequest(false, "/path/to/publickey.cer", "/path/to/privatekey.pem")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(saml)
}

func NewExampleFunc() driver.SamlSettingsFunc {
	return func(settings *driver.SamlSettings) {
		settings.AppSettings.Issuer = "issuer"
		settings.AppSettings.AssertionConsumerServiceURL = "http://www.onelogin.net"
		settings.AccountSettings.Certificate = "cert"
		settings.AccountSettings.IDP_SSO_Target_URL = "http://www.onelogin.net"
	}
}
