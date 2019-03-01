package saml

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"github.com/autom8ter/identify/certificates"
	"github.com/autom8ter/identify/saml/driver"
	"github.com/nu7hatch/gouuid"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"
)

type AuthorizationRequest struct {
	Id                 string
	IssueInstant       string
	AppSettings        *driver.AppSettings
	AccountSettings    *driver.AccountSettings
	Base64             int
}

func NewAuthorizationRequest(settingsFn ...driver.SamlSettingsFunc) *AuthorizationRequest {
	myIdUUID, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Error is UUID Generation:", err)
	}
	//yyyy-MM-dd'T'H:mm:ss
	layout := "2006-01-02T15:04:05"
	t := time.Now().UTC().Format(layout)
	sett := driver.NewSamlSettings(settingsFn...)
	return &AuthorizationRequest{
		AccountSettings: sett.AccountSettings,
		AppSettings:     sett.AppSettings,
		Id:              "_" + myIdUUID.String(),
		IssueInstant:    t,
	}
}

func NewAuthorizationRequestFromEnv() *AuthorizationRequest {
	myIdUUID, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Error is UUID Generation:", err)
	}
	//yyyy-MM-dd'T'H:mm:ss
	layout := "2006-01-02T15:04:05"
	t := time.Now().UTC().Format(layout)
	sett := driver.NewSamlSettings(driver.NewSettingsFromEnv())
	return &AuthorizationRequest{
		AccountSettings: sett.AccountSettings,
		AppSettings:     sett.AppSettings,
		Id:              "_" + myIdUUID.String(),
		IssueInstant:    t,
	}
}

// GetRequest returns a string formatted XML document that represents the SAML document
// TODO: parameterize more parts of the request
func (ar AuthorizationRequest) GetRequest(base64Encode bool) (string, error) {
	d := driver.AuthnRequest{
		XMLName: xml.Name{
			Local: "samlp:AuthnRequest",
		},
		SAMLP:                       "urn:oasis:names:tc:SAML:2.0:protocol",
		SAML:                        "urn:oasis:names:tc:SAML:2.0:assertion",
		ID:                          ar.Id,
		ProtocolBinding:             "urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST",
		Version:                     "2.0",
		AssertionConsumerServiceURL: ar.AppSettings.AssertionConsumerServiceURL,
		Issuer: driver.Issuer{
			XMLName: xml.Name{
				Local: "saml:Issuer",
			},
			Url: ar.AppSettings.Issuer,
		},
		IssueInstant: ar.IssueInstant,
		NameIDPolicy: driver.NameIDPolicy{
			XMLName: xml.Name{
				Local: "samlp:NameIDPolicy",
			},
			AllowCreate: true,
			Format:      "urn:oasis:names:tc:SAML:2.0:nameid-format:transient",
		},
		RequestedAuthnContext: driver.RequestedAuthnContext{
			XMLName: xml.Name{
				Local: "samlp:RequestedAuthnContext",
			},
			SAMLP:      "urn:oasis:names:tc:SAML:2.0:protocol",
			Comparison: "exact",

			AuthnContextClassRef: driver.AuthnContextClassRef{
				XMLName: xml.Name{
					Local: "saml:AuthnContextClassRef",
				},
				SAML:      "urn:oasis:names:tc:SAML:2.0:assertion",
				Transport: "urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport",
			},
		},
	}
	b, err := xml.MarshalIndent(d, "", "    ")
	if err != nil {
		return "", err
	}

	xmlAuthnRequest := fmt.Sprintf("<?xml version='1.0' encoding='UTF-8'?>\n%s", b)

	if base64Encode {
		data := []byte(xmlAuthnRequest)
		return base64.StdEncoding.EncodeToString(data), nil
	} else {
		return string(xmlAuthnRequest), nil
	}
}

// GetSignedRequest returns a string formatted XML document that represents the SAML document
// TODO: parameterize more parts of the request
func (ar AuthorizationRequest) GetSignedRequest(base64Encode bool, publicCert string, privateCert string) (string, error) {
	cert, err := certificates.LoadCertificate(publicCert)
	if err != nil {
		return "", err
	}

	d := driver.AuthnSignedRequest{
		XMLName: xml.Name{
			Local: "samlp:AuthnRequest",
		},
		SAMLP:                       "urn:oasis:names:tc:SAML:2.0:protocol",
		SAML:                        "urn:oasis:names:tc:SAML:2.0:assertion",
		SAMLSIG:                     "http://www.w3.org/2000/09/xmldsig#",
		ID:                          ar.Id,
		ProtocolBinding:             "urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST",
		Version:                     "2.0",
		AssertionConsumerServiceURL: ar.AppSettings.AssertionConsumerServiceURL,
		Issuer: driver.Issuer{
			XMLName: xml.Name{
				Local: "saml:Issuer",
			},
			Url: ar.AppSettings.Issuer,
		},
		IssueInstant: ar.IssueInstant,
		NameIDPolicy: driver.NameIDPolicy{
			XMLName: xml.Name{
				Local: "samlp:NameIDPolicy",
			},
			AllowCreate: true,
			Format:      "urn:oasis:names:tc:SAML:2.0:nameid-format:transient",
		},
		RequestedAuthnContext: driver.RequestedAuthnContext{
			XMLName: xml.Name{
				Local: "samlp:RequestedAuthnContext",
			},
			SAMLP:      "urn:oasis:names:tc:SAML:2.0:protocol",
			Comparison: "exact",
		},
		AuthnContextClassRef: driver.AuthnContextClassRef{
			XMLName: xml.Name{
				Local: "saml:AuthnContextClassRef",
			},
			SAML:      "urn:oasis:names:tc:SAML:2.0:assertion",
			Transport: "urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport",
		},
		Signature: driver.Signature{
			XMLName: xml.Name{
				Local: "samlsig:Signature",
			},
			Id: "Signature1",
			SignedInfo: driver.SignedInfo{
				XMLName: xml.Name{
					Local: "samlsig:SignedInfo",
				},
				CanonicalizationMethod: driver.CanonicalizationMethod{
					XMLName: xml.Name{
						Local: "samlsig:CanonicalizationMethod",
					},
					Algorithm: "http://www.w3.org/2001/10/xml-exc-c14n#",
				},
				SignatureMethod: driver.SignatureMethod{
					XMLName: xml.Name{
						Local: "samlsig:SignatureMethod",
					},
					Algorithm: "http://www.w3.org/2000/09/xmldsig#rsa-sha1",
				},
				SamlsigReference: driver.SamlsigReference{
					XMLName: xml.Name{
						Local: "samlsig:Reference",
					},
					URI: "#" + ar.Id,
					Transforms: driver.Transforms{
						XMLName: xml.Name{
							Local: "samlsig:Transforms",
						},
						Transform: driver.Transform{
							XMLName: xml.Name{
								Local: "samlsig:Transform",
							},
							Algorithm: "http://www.w3.org/2000/09/xmldsig#enveloped-signature",
						},
					},
					DigestMethod: driver.DigestMethod{
						XMLName: xml.Name{
							Local: "samlsig:DigestMethod",
						},
						Algorithm: "http://www.w3.org/2000/09/xmldsig#sha1",
					},
					DigestValue: driver.DigestValue{
						XMLName: xml.Name{
							Local: "samlsig:DigestValue",
						},
					},
				},
			},
			SignatureValue: driver.SignatureValue{
				XMLName: xml.Name{
					Local: "samlsig:SignatureValue",
				},
			},
			KeyInfo: driver.KeyInfo{
				XMLName: xml.Name{
					Local: "samlsig:KeyInfo",
				},
				X509Data: driver.X509Data{
					XMLName: xml.Name{
						Local: "samlsig:X509Data",
					},
					X509Certificate: driver.X509Certificate{
						XMLName: xml.Name{
							Local: "samlsig:X509Certificate",
						},
						Cert: cert,
					},
				},
			},
		},
	}
	b, err := xml.MarshalIndent(d, "", "    ")
	if err != nil {
		return "", err
	}

	samlAuthnRequest := string(b)
	// Write the SAML to a file.

	samlXmlsecInput, err := ioutil.TempFile(os.TempDir(), "tmpgs")
	if err != nil {
		return "", err
	}
	samlXmlsecOutput, err := ioutil.TempFile(os.TempDir(), "tmpgs")
	if err != nil {
		return "", err
	}

	samlXmlsecOutput.Close()

	samlXmlsecInput.WriteString("<?xml version='1.0' encoding='UTF-8'?>\n")
	samlXmlsecInput.WriteString(samlAuthnRequest)
	samlXmlsecInput.Close()

	_, errOut := exec.Command("xmlsec1", "--sign", "--privkey-pem", privateCert,
		"--id-attr:ID", "urn:oasis:names:tc:SAML:2.0:protocol:AuthnRequest",
		"--output", samlXmlsecOutput.Name(), samlXmlsecInput.Name()).Output()
	if errOut != nil {
		return "", errOut
	}

	samlSignedRequest, err := ioutil.ReadFile(samlXmlsecOutput.Name())
	if err != nil {
		return "", err
	}
	samlSignedRequestXml := strings.Trim(string(samlSignedRequest), "\n")

	if base64Encode {
		data := []byte(samlSignedRequestXml)
		return base64.StdEncoding.EncodeToString(data), nil
	} else {
		return string(samlSignedRequestXml), nil
	}
}

func (ar AuthorizationRequest) GetRequestUrl() (string, error) {
	u, err := url.Parse(ar.AccountSettings.IDP_SSO_Target_URL)
	if err != nil {
		return "", err
	}
	base64EncodedUTF8SamlRequest, err := ar.GetRequest(true)
	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Add("SAMLRequest", base64EncodedUTF8SamlRequest)

	u.RawQuery = q.Encode()
	return u.String(), nil
}

type Response struct {
	XmlDoc      string
	Settings    driver.AccountSettings
	certificate x509.Certificate
}

type ModifyRequestFunc func(req *AuthorizationRequest) *AuthorizationRequest
