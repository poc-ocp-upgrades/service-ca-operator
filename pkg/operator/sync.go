package operator

import (
	"k8s.io/klog"
	operatorv1 "github.com/openshift/api/operator/v1"
)

func syncControllers(c serviceCAOperator, operatorConfig *operatorv1.ServiceCA) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	needsDeploy, err := manageControllerNS(c)
	if err != nil {
		return err
	}
	err = manageSignerControllerResources(c, &needsDeploy)
	if err != nil {
		return err
	}
	err = manageAPIServiceControllerResources(c, &needsDeploy)
	if err != nil {
		return err
	}
	err = manageConfigMapCABundleControllerResources(c, &needsDeploy)
	if err != nil {
		return err
	}
	caModified, err := manageSignerCA(c.corev1Client, c.eventRecorder)
	if err != nil {
		return err
	}
	_, err = manageSignerCABundle(c.corev1Client, c.eventRecorder, caModified)
	if err != nil {
		return err
	}
	configModified, err := manageSignerControllerConfig(c.corev1Client, c.eventRecorder)
	if err != nil {
		return err
	}
	_, err = manageSignerControllerDeployment(c.appsv1Client, c.eventRecorder, operatorConfig, needsDeploy || caModified || configModified)
	if err != nil {
		return err
	}
	configModified, err = manageAPIServiceControllerConfig(c.corev1Client, c.eventRecorder)
	if err != nil {
		return err
	}
	_, err = manageAPIServiceControllerDeployment(c.appsv1Client, c.eventRecorder, operatorConfig, needsDeploy || caModified || configModified)
	if err != nil {
		return err
	}
	configModified, err = manageConfigMapCABundleControllerConfig(c.corev1Client, c.eventRecorder)
	if err != nil {
		return err
	}
	_, err = manageConfigMapCABundleControllerDeployment(c.appsv1Client, c.eventRecorder, operatorConfig, needsDeploy || caModified || configModified)
	if err != nil {
		return err
	}
	klog.V(4).Infof("synced all controller resources")
	return nil
}
