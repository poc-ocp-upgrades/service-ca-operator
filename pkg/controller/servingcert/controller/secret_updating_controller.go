package controller

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"
	"k8s.io/api/core/v1"
	kapierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	informers "k8s.io/client-go/informers/core/v1"
	kcoreclient "k8s.io/client-go/kubernetes/typed/core/v1"
	listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/klog"
	ocontroller "github.com/openshift/library-go/pkg/controller"
	"github.com/openshift/library-go/pkg/crypto"
	"github.com/openshift/service-ca-operator/pkg/boilerplate/controller"
	"github.com/openshift/service-ca-operator/pkg/controller/api"
)

type serviceServingCertUpdateController struct {
	secretClient		kcoreclient.SecretsGetter
	serviceLister		listers.ServiceLister
	secretLister		listers.SecretLister
	ca			*crypto.CA
	dnsSuffix		string
	minTimeLeftForCert	time.Duration
}

func NewServiceServingCertUpdateController(services informers.ServiceInformer, secrets informers.SecretInformer, secretClient kcoreclient.SecretsGetter, ca *crypto.CA, dnsSuffix string) controller.Runner {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	sc := &serviceServingCertUpdateController{secretClient: secretClient, serviceLister: services.Lister(), secretLister: secrets.Lister(), ca: ca, dnsSuffix: dnsSuffix, minTimeLeftForCert: 1 * time.Hour}
	return controller.New("ServiceServingCertUpdateController", sc, controller.WithInformerSynced(services), controller.WithInformer(secrets, controller.FilterFuncs{AddFunc: sc.addSecret, UpdateFunc: sc.updateSecret}))
}
func (sc *serviceServingCertUpdateController) addSecret(obj metav1.Object) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	secret := obj.(*v1.Secret)
	_, ok := toServiceName(secret)
	return ok
}
func (sc *serviceServingCertUpdateController) updateSecret(old, cur metav1.Object) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return sc.addSecret(cur) || sc.addSecret(old)
}
func (sc *serviceServingCertUpdateController) Key(namespace, name string) (metav1.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return sc.secretLister.Secrets(namespace).Get(name)
}
func (sc *serviceServingCertUpdateController) Sync(obj metav1.Object) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	sharedSecret := obj.(*v1.Secret)
	service := sc.getServiceForSecret(sharedSecret)
	if service == nil {
		return nil
	}
	if !isSecretValidForService(service, sharedSecret) {
		return nil
	}
	secretCopy := sharedSecret.DeepCopy()
	if sc.requiresRegeneration(service, sharedSecret, sc.minTimeLeftForCert) {
		if err := toRequiredSecret(sc.dnsSuffix, sc.ca, service, secretCopy); err != nil {
			return err
		}
		_, err := sc.secretClient.Secrets(secretCopy.Namespace).Update(secretCopy)
		return err
	}
	update, err := sc.ensureSecretData(service, secretCopy)
	if err != nil {
		return err
	}
	if update {
		_, err := sc.secretClient.Secrets(secretCopy.Namespace).Update(secretCopy)
		return err
	}
	return nil
}
func isSecretValidForService(sharedService *v1.Service, secret *v1.Secret) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	isValid := true
	if sharedService.Annotations[api.ServingCertSecretAnnotation] != secret.Name && sharedService.Annotations[api.AlphaServingCertSecretAnnotation] != secret.Name {
		isValid = false
	}
	if secret.Annotations[api.ServiceUIDAnnotation] != string(sharedService.UID) && secret.Annotations[api.AlphaServiceUIDAnnotation] != string(sharedService.UID) {
		isValid = false
	}
	return isValid
}
func (sc *serviceServingCertUpdateController) getServiceForSecret(sharedSecret *v1.Secret) *v1.Service {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	serviceName, ok := toServiceName(sharedSecret)
	if !ok {
		return nil
	}
	service, err := sc.serviceLister.Services(sharedSecret.Namespace).Get(serviceName)
	if kapierrors.IsNotFound(err) {
		return nil
	}
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("unable to get service %s/%s: %v", sharedSecret.Namespace, serviceName, err))
		return nil
	}
	return service
}
func (sc *serviceServingCertUpdateController) requiresRegeneration(service *v1.Service, secret *v1.Secret, minTimeLeft time.Duration) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !ocontroller.HasOwnerRef(secret, ownerRef(service)) {
		return true
	}
	expiryString, ok := secret.Annotations[api.ServingCertExpiryAnnotation]
	if !ok {
		expiryString, ok = secret.Annotations[api.AlphaServingCertExpiryAnnotation]
		if !ok {
			return true
		}
	}
	expiry, err := time.Parse(time.RFC3339, expiryString)
	if err != nil {
		return true
	}
	if time.Now().Add(sc.minTimeLeftForCert).After(expiry) {
		return true
	}
	return false
}
func (sc *serviceServingCertUpdateController) ensureSecretData(service *v1.Service, secretCopy *v1.Secret) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	update := false
	tlsCert, ok := secretCopy.Data[v1.TLSCertKey]
	tlsKey, ok2 := secretCopy.Data[v1.TLSPrivateKeyKey]
	if ok && ok2 {
		if len(secretCopy.Data) != 2 {
			secretCopy.Data = map[string][]byte{v1.TLSCertKey: tlsCert, v1.TLSPrivateKeyKey: tlsKey}
			update = true
		}
	} else {
		if err := toRequiredSecret(sc.dnsSuffix, sc.ca, service, secretCopy); err != nil {
			return update, err
		}
		return true, nil
	}
	block, _ := pem.Decode([]byte(tlsCert))
	if block == nil {
		klog.Infof("Error decoding cert bytes %s from secret: %s namespace: %s, replacing cert", v1.TLSCertKey, secretCopy.Name, secretCopy.Namespace)
		if err := toRequiredSecret(sc.dnsSuffix, sc.ca, service, secretCopy); err != nil {
			return update, err
		}
		return true, nil
	}
	_, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		klog.Infof("Error parsing %s from secret: %s namespace: %s, replacing cert", v1.TLSCertKey, secretCopy.Name, secretCopy.Namespace)
		if err := toRequiredSecret(sc.dnsSuffix, sc.ca, service, secretCopy); err != nil {
			return update, err
		}
		return true, nil
	}
	return update, nil
}
func toServiceName(secret *v1.Secret) (string, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	serviceName := secret.Annotations[api.ServiceNameAnnotation]
	if len(serviceName) == 0 {
		serviceName = secret.Annotations[api.AlphaServiceNameAnnotation]
	}
	return serviceName, len(serviceName) != 0
}
