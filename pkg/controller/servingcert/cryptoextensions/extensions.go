package cryptoextensions

import (
	"encoding/asn1"
)

func oid(o asn1.ObjectIdentifier, extra ...int) asn1.ObjectIdentifier {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return asn1.ObjectIdentifier(append(append([]int{}, o...), extra...))
}

var (
	RedHatOID	= asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 2312}
	OpenShiftOID	= oid(RedHatOID, 17)
)
