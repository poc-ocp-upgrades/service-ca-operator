package operator

import (
	"fmt"
	"os"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	operatorv1 "github.com/openshift/api/operator/v1"
	"github.com/openshift/library-go/pkg/operator/v1helpers"
)

func setFailingTrue(operatorConfig *operatorv1.ServiceCA, reason, message string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	v1helpers.SetOperatorCondition(&operatorConfig.Status.Conditions, operatorv1.OperatorCondition{Type: operatorv1.OperatorStatusTypeFailing, Status: operatorv1.ConditionTrue, Reason: reason, Message: message})
}
func setFailingFalse(operatorConfig *operatorv1.ServiceCA, reason string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	v1helpers.SetOperatorCondition(&operatorConfig.Status.Conditions, operatorv1.OperatorCondition{Type: operatorv1.OperatorStatusTypeFailing, Status: operatorv1.ConditionFalse, Reason: reason})
}
func setProgressingTrue(operatorConfig *operatorv1.ServiceCA, reason, message string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	v1helpers.SetOperatorCondition(&operatorConfig.Status.Conditions, operatorv1.OperatorCondition{Type: operatorv1.OperatorStatusTypeProgressing, Status: operatorv1.ConditionTrue, Reason: reason, Message: message})
}
func setAvailableTrue(operatorConfig *operatorv1.ServiceCA, reason string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	v1helpers.SetOperatorCondition(&operatorConfig.Status.Conditions, operatorv1.OperatorCondition{Type: operatorv1.OperatorStatusTypeAvailable, Status: operatorv1.ConditionTrue, Reason: reason})
}
func setProgressingFalse(operatorConfig *operatorv1.ServiceCA, reason, message string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	v1helpers.SetOperatorCondition(&operatorConfig.Status.Conditions, operatorv1.OperatorCondition{Type: operatorv1.OperatorStatusTypeProgressing, Status: operatorv1.ConditionFalse, Reason: reason, Message: message})
}
func setAvailableFalse(operatorConfig *operatorv1.ServiceCA, reason, message string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	v1helpers.SetOperatorCondition(&operatorConfig.Status.Conditions, operatorv1.OperatorCondition{Type: operatorv1.OperatorStatusTypeAvailable, Status: operatorv1.ConditionFalse, Reason: reason, Message: message})
}
func isDeploymentStatusAvailable(deploy appsv1.Deployment) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return deploy.Status.AvailableReplicas > 0
}
func isDeploymentStatusAvailableAndUpdated(deploy appsv1.Deployment) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return deploy.Status.AvailableReplicas > 0 && deploy.Status.ObservedGeneration >= deploy.Generation && deploy.Status.UpdatedReplicas == deploy.Status.Replicas
}
func isDeploymentStatusComplete(deploy appsv1.Deployment) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	replicas := int32(1)
	if deploy.Spec.Replicas != nil {
		replicas = *(deploy.Spec.Replicas)
	}
	return deploy.Status.UpdatedReplicas == replicas && deploy.Status.Replicas == replicas && deploy.Status.AvailableReplicas == replicas && deploy.Status.ObservedGeneration >= deploy.Generation
}
func (c *serviceCAOperator) syncStatus(operatorConfigCopy *operatorv1.ServiceCA, existingDeployments *appsv1.DeploymentList, targetDeploymentNames sets.String) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	versionUpdatable := true
	versionUpdatableAndDeploymentsComplete := true
	statusMsg := ""
	existingDeploymentNames := sets.String{}
	for _, dep := range existingDeployments.Items {
		existingDeploymentNames.Insert(dep.Name)
		reason := "ManagedDeploymentsNotReady"
		if dep.DeletionTimestamp != nil {
			statusMsg += fmt.Sprintf("\n%s deleting", dep.Name)
			setProgressingTrue(operatorConfigCopy, reason, statusMsg)
			setAvailableFalse(operatorConfigCopy, reason, statusMsg)
			versionUpdatable = false
			versionUpdatableAndDeploymentsComplete = false
			continue
		}
		if !isDeploymentStatusAvailable(dep) {
			statusMsg += fmt.Sprintf("\n%s does not have available replicas", dep.Name)
			setProgressingTrue(operatorConfigCopy, reason, statusMsg)
			setAvailableFalse(operatorConfigCopy, reason, statusMsg)
			versionUpdatable = false
			versionUpdatableAndDeploymentsComplete = false
			continue
		}
		if !isDeploymentStatusAvailableAndUpdated(dep) {
			statusMsg += fmt.Sprintf("\n%s is updating", dep.Name)
			versionUpdatable = false
			versionUpdatableAndDeploymentsComplete = false
			continue
		} else if !isDeploymentStatusComplete(dep) {
			versionUpdatableAndDeploymentsComplete = false
			statusMsg += fmt.Sprintf("\n%s is creating replicas.", dep.Name)
			continue
		}
	}
	missing := targetDeploymentNames.Difference(existingDeploymentNames)
	if len(missing) > 0 {
		reason := "ManagedDeploymentsNotFound"
		statusMsg = fmt.Sprintf("Deployments %v not found", missing)
		setProgressingTrue(operatorConfigCopy, reason, statusMsg)
		setAvailableFalse(operatorConfigCopy, reason, statusMsg)
		return
	}
	if versionUpdatableAndDeploymentsComplete {
		reason := "ManagedDeploymentsCompleteAndUpdated"
		setAvailableTrue(operatorConfigCopy, reason)
		setProgressingFalse(operatorConfigCopy, reason, "All service-ca-operator deployments updated")
		c.setVersion()
		return
	}
	if versionUpdatable {
		reason := "ManagedDeploymentsAvailableAndUpdated"
		setAvailableTrue(operatorConfigCopy, reason)
		setProgressingTrue(operatorConfigCopy, reason, statusMsg)
		c.setVersion()
		return
	}
	reason := "ManagedDeploymentsAvailable"
	setAvailableTrue(operatorConfigCopy, reason)
	setProgressingTrue(operatorConfigCopy, reason, statusMsg)
}
func (c *serviceCAOperator) setVersion() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	version := os.Getenv(operatorVersionEnvName)
	if c.versionGetter.GetVersions()["operator"] != version {
		c.versionGetter.SetVersion("operator", version)
	}
}
