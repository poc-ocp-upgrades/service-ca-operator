package api

const (
	OperatorConfigInstanceName			= "cluster"
	SignerControllerConfigMapName		= "service-serving-cert-signer-config"
	APIServiceInjectorConfigMapName		= "apiservice-cabundle-injector-config"
	ConfigMapInjectorConfigMapName		= "configmap-cabundle-injector-config"
	SigningCABundleConfigMapName		= "signing-cabundle"
	SignerControllerSAName				= "service-serving-cert-signer-sa"
	APIServiceInjectorSAName			= "apiservice-cabundle-injector-sa"
	ConfigMapInjectorSAName				= "configmap-cabundle-injector-sa"
	SignerControllerServiceName			= "service-serving-cert-signer"
	SignerControllerDeploymentName		= "service-serving-cert-signer"
	APIServiceInjectorDeploymentName	= "apiservice-cabundle-injector"
	ConfigMapInjectorDeploymentName		= "configmap-cabundle-injector"
	SignerControllerSecretName			= "service-serving-cert-signer-signing-key"
)
