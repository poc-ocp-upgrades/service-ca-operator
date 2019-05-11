package cryptoextensions

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	kapiv1 "k8s.io/api/core/v1"
	"github.com/openshift/library-go/pkg/crypto"
)

var (
	OpenShiftServerSigningOID			= oid(OpenShiftOID, 100)
	OpenShiftServerSigningServiceOID	= oid(OpenShiftServerSigningOID, 2)
	OpenShiftServerSigningServiceUIDOID	= oid(OpenShiftServerSigningServiceOID, 1)
)

func ServiceServerCertificateExtensionV1(svc *kapiv1.Service) crypto.CertificateExtensionFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(cert *x509.Certificate) error {
		uid, err := asn1.Marshal(svc.UID)
		if err != nil {
			return err
		}
		cert.ExtraExtensions = append(cert.ExtraExtensions, pkix.Extension{Id: OpenShiftServerSigningServiceUIDOID, Critical: false, Value: uid})
		return nil
	}
}
