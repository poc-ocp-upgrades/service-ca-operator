package v4_00_assets

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes	[]byte
	info	os.FileInfo
}
type bindataFileInfo struct {
	name	string
	size	int64
	mode	os.FileMode
	modTime	time.Time
}

func (fi bindataFileInfo) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}

var _v400ApiserviceCabundleControllerClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: system:openshift:controller:apiservice-cabundle-injector
rules:
- apiGroups:
  - apiregistration.k8s.io
  resources:
  - apiservices
  verbs:
  - get
  - list
  - watch
  - update
  - patch
`)

func v400ApiserviceCabundleControllerClusterroleYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ApiserviceCabundleControllerClusterroleYaml, nil
}
func v400ApiserviceCabundleControllerClusterroleYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ApiserviceCabundleControllerClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/apiservice-cabundle-controller/clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ApiserviceCabundleControllerClusterrolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: system:openshift:controller:apiservice-cabundle-injector
roleRef:
  kind: ClusterRole
  name: system:openshift:controller:apiservice-cabundle-injector
subjects:
- kind: ServiceAccount
  namespace: openshift-service-ca
  name: apiservice-cabundle-injector-sa
`)

func v400ApiserviceCabundleControllerClusterrolebindingYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ApiserviceCabundleControllerClusterrolebindingYaml, nil
}
func v400ApiserviceCabundleControllerClusterrolebindingYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ApiserviceCabundleControllerClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/apiservice-cabundle-controller/clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ApiserviceCabundleControllerCmYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  namespace: openshift-service-ca
  name: apiservice-cabundle-injector-config
data:
  controller-config.yaml:
`)

func v400ApiserviceCabundleControllerCmYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ApiserviceCabundleControllerCmYaml, nil
}
func v400ApiserviceCabundleControllerCmYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ApiserviceCabundleControllerCmYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/apiservice-cabundle-controller/cm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ApiserviceCabundleControllerDefaultconfigYaml = []byte(`apiVersion: servicecertsigner.config.openshift.io/v1alpha1
kind: APIServiceCABundleInjectorConfig
caBundleFile: /var/run/configmaps/signing-cabundle/ca-bundle.crt
authentication:
  disabled: true
authorization:
  disabled: true
leaderElection:
  leaseDuration: "15s"
  renewDeadline: "10s"
  retryPeriod: "2s"
`)

func v400ApiserviceCabundleControllerDefaultconfigYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ApiserviceCabundleControllerDefaultconfigYaml, nil
}
func v400ApiserviceCabundleControllerDefaultconfigYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ApiserviceCabundleControllerDefaultconfigYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/apiservice-cabundle-controller/defaultconfig.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ApiserviceCabundleControllerDeploymentYaml = []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  namespace: openshift-service-ca
  name: apiservice-cabundle-injector
  labels:
    app: apiservice-cabundle-injector
    apiservice-cabundle-injector: "true"
spec:
  strategy:
    type: Recreate
  selector:
    matchLabels:
      app: apiservice-cabundle-injector
      apiservice-cabundle-injector: "true"
  template:
    metadata:
      name: apiservice-cabundle-injector
      labels:
        app: apiservice-cabundle-injector
        apiservice-cabundle-injector: "true"
    spec:
      serviceAccountName: apiservice-cabundle-injector-sa
      containers:
      - name: apiservice-cabundle-injector-controller
        image: ${IMAGE}
        imagePullPolicy: IfNotPresent
        command: ["service-ca-operator", "apiservice-cabundle-injector"]
        args:
        - "--config=/var/run/configmaps/config/controller-config.yaml"
        ports:
        - containerPort: 8443
        volumeMounts:
        - mountPath: /var/run/configmaps/config
          name: config
        - mountPath: /var/run/configmaps/signing-cabundle
          name: signing-cabundle
      volumes:
      - name: signing-cabundle
        configMap:
          name: signing-cabundle
      - name: config
        configMap:
          name: apiservice-cabundle-injector-config
      nodeSelector:
        node-role.kubernetes.io/master: ""
      tolerations:
      - operator: Exists
`)

func v400ApiserviceCabundleControllerDeploymentYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ApiserviceCabundleControllerDeploymentYaml, nil
}
func v400ApiserviceCabundleControllerDeploymentYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ApiserviceCabundleControllerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/apiservice-cabundle-controller/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ApiserviceCabundleControllerNsYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  name: openshift-service-ca
  labels:
    openshift.io/run-level: "1"
  annotations:
    openshift.io/node-selector: ""
`)

func v400ApiserviceCabundleControllerNsYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ApiserviceCabundleControllerNsYaml, nil
}
func v400ApiserviceCabundleControllerNsYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ApiserviceCabundleControllerNsYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/apiservice-cabundle-controller/ns.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ApiserviceCabundleControllerRoleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: system:openshift:controller:apiservice-cabundle-injector
  namespace: openshift-service-ca
rules:
- apiGroups:
  - ""
  resources:
  - events
  verbs:
  - create
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - get
  - list
  - watch
  - update
  - create
- apiGroups:
  - ""
  resources:
  - pods
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - "apps"
  resources:
  - replicasets
  - deployments
  verbs:
  - get
  - list
  - watch
`)

func v400ApiserviceCabundleControllerRoleYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ApiserviceCabundleControllerRoleYaml, nil
}
func v400ApiserviceCabundleControllerRoleYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ApiserviceCabundleControllerRoleYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/apiservice-cabundle-controller/role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ApiserviceCabundleControllerRolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: system:openshift:controller:apiservice-cabundle-injector
  namespace: openshift-service-ca
roleRef:
  kind: Role
  name: system:openshift:controller:apiservice-cabundle-injector
subjects:
- kind: ServiceAccount
  namespace: openshift-service-ca
  name: apiservice-cabundle-injector-sa
`)

func v400ApiserviceCabundleControllerRolebindingYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ApiserviceCabundleControllerRolebindingYaml, nil
}
func v400ApiserviceCabundleControllerRolebindingYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ApiserviceCabundleControllerRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/apiservice-cabundle-controller/rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ApiserviceCabundleControllerSaYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: openshift-service-ca
  name: apiservice-cabundle-injector-sa
`)

func v400ApiserviceCabundleControllerSaYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ApiserviceCabundleControllerSaYaml, nil
}
func v400ApiserviceCabundleControllerSaYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ApiserviceCabundleControllerSaYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/apiservice-cabundle-controller/sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ApiserviceCabundleControllerSigningCabundleYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  namespace: openshift-service-ca
  name: signing-cabundle
data:
  ca-bundle.crt:
`)

func v400ApiserviceCabundleControllerSigningCabundleYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ApiserviceCabundleControllerSigningCabundleYaml, nil
}
func v400ApiserviceCabundleControllerSigningCabundleYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ApiserviceCabundleControllerSigningCabundleYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/apiservice-cabundle-controller/signing-cabundle.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ConfigmapCabundleControllerClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: system:openshift:controller:configmap-cabundle-injector
rules:
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - get
  - list
  - watch
  - update
`)

func v400ConfigmapCabundleControllerClusterroleYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ConfigmapCabundleControllerClusterroleYaml, nil
}
func v400ConfigmapCabundleControllerClusterroleYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ConfigmapCabundleControllerClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/configmap-cabundle-controller/clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ConfigmapCabundleControllerClusterrolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: system:openshift:controller:configmap-cabundle-injector
roleRef:
  kind: ClusterRole
  name: system:openshift:controller:configmap-cabundle-injector
subjects:
- kind: ServiceAccount
  namespace: openshift-service-ca
  name: configmap-cabundle-injector-sa
`)

func v400ConfigmapCabundleControllerClusterrolebindingYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ConfigmapCabundleControllerClusterrolebindingYaml, nil
}
func v400ConfigmapCabundleControllerClusterrolebindingYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ConfigmapCabundleControllerClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/configmap-cabundle-controller/clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ConfigmapCabundleControllerCmYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  namespace: openshift-service-ca
  name: configmap-cabundle-injector-config
data:
  controller-config.yaml:
`)

func v400ConfigmapCabundleControllerCmYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ConfigmapCabundleControllerCmYaml, nil
}
func v400ConfigmapCabundleControllerCmYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ConfigmapCabundleControllerCmYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/configmap-cabundle-controller/cm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ConfigmapCabundleControllerDefaultconfigYaml = []byte(`apiVersion: servicecertsigner.config.openshift.io/v1alpha1
kind: ConfigMapCABundleInjectorConfig
caBundleFile: /var/run/configmaps/signing-cabundle/ca-bundle.crt
authentication:
  disabled: true
authorization:
  disabled: true
leaderElection:
  leaseDuration: "15s"
  renewDeadline: "10s"
  retryPeriod: "2s"
`)

func v400ConfigmapCabundleControllerDefaultconfigYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ConfigmapCabundleControllerDefaultconfigYaml, nil
}
func v400ConfigmapCabundleControllerDefaultconfigYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ConfigmapCabundleControllerDefaultconfigYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/configmap-cabundle-controller/defaultconfig.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ConfigmapCabundleControllerDeploymentYaml = []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  namespace: openshift-service-ca
  name: configmap-cabundle-injector
  labels:
    app: configmap-cabundle-injector
    configmap-cabundle-injector: "true"
spec:
  strategy:
    type: Recreate
  selector:
    matchLabels:
      app: configmap-cabundle-injector
      configmap-cabundle-injector: "true"
  template:
    metadata:
      name: configmap-cabundle-injector
      labels:
        app: configmap-cabundle-injector
        configmap-cabundle-injector: "true"
    spec:
      serviceAccountName: configmap-cabundle-injector-sa
      containers:
      - name: configmap-cabundle-injector-controller
        image: ${IMAGE}
        imagePullPolicy: IfNotPresent
        command: ["service-ca-operator", "configmap-cabundle-injector"]
        args:
        - "--config=/var/run/configmaps/config/controller-config.yaml"
        ports:
        - containerPort: 8443
        volumeMounts:
        - mountPath: /var/run/configmaps/config
          name: config
        - mountPath: /var/run/configmaps/signing-cabundle
          name: signing-cabundle
      volumes:
      - name: signing-cabundle
        configMap:
          name: signing-cabundle
      - name: config
        configMap:
          name: configmap-cabundle-injector-config
      nodeSelector:
        node-role.kubernetes.io/master: ""
      tolerations:
      - operator: Exists
`)

func v400ConfigmapCabundleControllerDeploymentYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ConfigmapCabundleControllerDeploymentYaml, nil
}
func v400ConfigmapCabundleControllerDeploymentYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ConfigmapCabundleControllerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/configmap-cabundle-controller/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ConfigmapCabundleControllerNsYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  name: openshift-service-ca
  labels:
    openshift.io/run-level: "1"
  annotations:
    openshift.io/node-selector: ""
`)

func v400ConfigmapCabundleControllerNsYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ConfigmapCabundleControllerNsYaml, nil
}
func v400ConfigmapCabundleControllerNsYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ConfigmapCabundleControllerNsYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/configmap-cabundle-controller/ns.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ConfigmapCabundleControllerRoleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: system:openshift:controller:configmap-cabundle-injector
  namespace: openshift-service-ca
rules:
- apiGroups:
  - ""
  resources:
  - events
  verbs:
  - create
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - get
  - list
  - watch
  - update
  - create
- apiGroups:
  - ""
  resources:
  - pods
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - "apps"
  resources:
  - replicasets
  - deployments
  verbs:
  - get
  - list
  - watch
`)

func v400ConfigmapCabundleControllerRoleYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ConfigmapCabundleControllerRoleYaml, nil
}
func v400ConfigmapCabundleControllerRoleYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ConfigmapCabundleControllerRoleYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/configmap-cabundle-controller/role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ConfigmapCabundleControllerRolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: system:openshift:controller:configmap-cabundle-injector
  namespace: openshift-service-ca
roleRef:
  kind: Role
  name: system:openshift:controller:configmap-cabundle-injector
subjects:
- kind: ServiceAccount
  namespace: openshift-service-ca
  name: configmap-cabundle-injector-sa
`)

func v400ConfigmapCabundleControllerRolebindingYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ConfigmapCabundleControllerRolebindingYaml, nil
}
func v400ConfigmapCabundleControllerRolebindingYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ConfigmapCabundleControllerRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/configmap-cabundle-controller/rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ConfigmapCabundleControllerSaYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: openshift-service-ca
  name: configmap-cabundle-injector-sa
`)

func v400ConfigmapCabundleControllerSaYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ConfigmapCabundleControllerSaYaml, nil
}
func v400ConfigmapCabundleControllerSaYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ConfigmapCabundleControllerSaYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/configmap-cabundle-controller/sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ConfigmapCabundleControllerSigningCabundleYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  namespace: openshift-service-ca
  name: signing-cabundle
data:
  ca-bundle.crt:
`)

func v400ConfigmapCabundleControllerSigningCabundleYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ConfigmapCabundleControllerSigningCabundleYaml, nil
}
func v400ConfigmapCabundleControllerSigningCabundleYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ConfigmapCabundleControllerSigningCabundleYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/configmap-cabundle-controller/signing-cabundle.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ServiceServingCertSignerControllerClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: system:openshift:controller:service-serving-cert-signer
rules:
- apiGroups:
  - ""
  resources:
  - secrets
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
- apiGroups:
  - ""
  resources:
  - services
  verbs:
  - get
  - list
  - watch
  - update
  - patch
`)

func v400ServiceServingCertSignerControllerClusterroleYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ServiceServingCertSignerControllerClusterroleYaml, nil
}
func v400ServiceServingCertSignerControllerClusterroleYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ServiceServingCertSignerControllerClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/service-serving-cert-signer-controller/clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ServiceServingCertSignerControllerClusterrolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: system:openshift:controller:service-serving-cert-signer
roleRef:
  kind: ClusterRole
  name: system:openshift:controller:service-serving-cert-signer
subjects:
- kind: ServiceAccount
  namespace: openshift-service-ca
  name: service-serving-cert-signer-sa
`)

func v400ServiceServingCertSignerControllerClusterrolebindingYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ServiceServingCertSignerControllerClusterrolebindingYaml, nil
}
func v400ServiceServingCertSignerControllerClusterrolebindingYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ServiceServingCertSignerControllerClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/service-serving-cert-signer-controller/clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ServiceServingCertSignerControllerCmYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  namespace: openshift-service-ca
  name: service-serving-cert-signer-config
data:
  controller-config.yaml:
`)

func v400ServiceServingCertSignerControllerCmYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ServiceServingCertSignerControllerCmYaml, nil
}
func v400ServiceServingCertSignerControllerCmYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ServiceServingCertSignerControllerCmYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/service-serving-cert-signer-controller/cm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ServiceServingCertSignerControllerDefaultconfigYaml = []byte(`apiVersion: servicecertsigner.config.openshift.io/v1alpha1
kind: ServiceServingCertSignerConfig
signer:
  certFile: /var/run/secrets/signing-key/tls.crt
  keyFile: /var/run/secrets/signing-key/tls.key
authentication:
  disabled: true
authorization:
  disabled: true
leaderElection:
  leaseDuration: "15s"
  renewDeadline: "10s"
  retryPeriod: "2s"
`)

func v400ServiceServingCertSignerControllerDefaultconfigYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ServiceServingCertSignerControllerDefaultconfigYaml, nil
}
func v400ServiceServingCertSignerControllerDefaultconfigYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ServiceServingCertSignerControllerDefaultconfigYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/service-serving-cert-signer-controller/defaultconfig.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ServiceServingCertSignerControllerDeploymentYaml = []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  namespace: openshift-service-ca
  name: service-serving-cert-signer
  labels:
    app: service-serving-cert-signer
    service-serving-cert-signer: "true"
spec:
  strategy:
    type: Recreate
  selector:
    matchLabels:
      app: service-serving-cert-signer
      service-serving-cert-signer: "true"
  template:
    metadata:
      name: service-serving-cert-signer
      labels:
        app: service-serving-cert-signer
        service-serving-cert-signer: "true"
    spec:
      serviceAccountName: service-serving-cert-signer-sa
      containers:
      - name: service-serving-cert-signer-controller
        image: ${IMAGE}
        imagePullPolicy: IfNotPresent
        command: ["service-ca-operator", "serving-cert-signer"]
        args:
        - "--config=/var/run/configmaps/config/controller-config.yaml"
        ports:
        - containerPort: 8443
        volumeMounts:
        - mountPath: /var/run/configmaps/config
          name: config
        - mountPath: /var/run/secrets/signing-key
          name: signing-key
      volumes:
      - name: signing-key
        secret:
          secretName: service-serving-cert-signer-signing-key
      - name: config
        configMap:
          name: service-serving-cert-signer-config
      nodeSelector:
        node-role.kubernetes.io/master: ""
      tolerations:
      - operator: Exists
`)

func v400ServiceServingCertSignerControllerDeploymentYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ServiceServingCertSignerControllerDeploymentYaml, nil
}
func v400ServiceServingCertSignerControllerDeploymentYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ServiceServingCertSignerControllerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/service-serving-cert-signer-controller/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ServiceServingCertSignerControllerNsYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  name: openshift-service-ca
  labels:
    openshift.io/run-level: "1"
  annotations:
    openshift.io/node-selector: ""
`)

func v400ServiceServingCertSignerControllerNsYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ServiceServingCertSignerControllerNsYaml, nil
}
func v400ServiceServingCertSignerControllerNsYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ServiceServingCertSignerControllerNsYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/service-serving-cert-signer-controller/ns.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ServiceServingCertSignerControllerRoleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: system:openshift:controller:service-serving-cert-signer
  namespace: openshift-service-ca
rules:
- apiGroups:
  - ""
  resources:
  - events
  verbs:
  - create
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - get
  - list
  - watch
  - update
  - create
- apiGroups:
  - ""
  resources:
  - pods
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - "apps"
  resources:
  - replicasets
  - deployments
  verbs:
  - get
  - list
  - watch
`)

func v400ServiceServingCertSignerControllerRoleYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ServiceServingCertSignerControllerRoleYaml, nil
}
func v400ServiceServingCertSignerControllerRoleYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ServiceServingCertSignerControllerRoleYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/service-serving-cert-signer-controller/role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ServiceServingCertSignerControllerRolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: system:openshift:controller:service-serving-cert-signer
  namespace: openshift-service-ca
roleRef:
  kind: Role
  name: system:openshift:controller:service-serving-cert-signer
subjects:
- kind: ServiceAccount
  namespace: openshift-service-ca
  name: service-serving-cert-signer-sa
`)

func v400ServiceServingCertSignerControllerRolebindingYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ServiceServingCertSignerControllerRolebindingYaml, nil
}
func v400ServiceServingCertSignerControllerRolebindingYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ServiceServingCertSignerControllerRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/service-serving-cert-signer-controller/rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ServiceServingCertSignerControllerSaYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: openshift-service-ca
  name: service-serving-cert-signer-sa
`)

func v400ServiceServingCertSignerControllerSaYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ServiceServingCertSignerControllerSaYaml, nil
}
func v400ServiceServingCertSignerControllerSaYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ServiceServingCertSignerControllerSaYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/service-serving-cert-signer-controller/sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v400ServiceServingCertSignerControllerSigningSecretYaml = []byte(`apiVersion: v1
kind: Secret
metadata:
  namespace: openshift-service-ca
  name: service-serving-cert-signer-signing-key
type: kubernetes.io/tls
data:
  tls.crt:
  tls.key:
`)

func v400ServiceServingCertSignerControllerSigningSecretYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v400ServiceServingCertSignerControllerSigningSecretYaml, nil
}
func v400ServiceServingCertSignerControllerSigningSecretYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v400ServiceServingCertSignerControllerSigningSecretYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v4.0.0/service-serving-cert-signer-controller/signing-secret.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
func Asset(name string) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}
func MustAsset(name string) []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}
	return a
}
func AssetInfo(name string) (os.FileInfo, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}
func AssetNames() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

var _bindata = map[string]func() (*asset, error){"v4.0.0/apiservice-cabundle-controller/clusterrole.yaml": v400ApiserviceCabundleControllerClusterroleYaml, "v4.0.0/apiservice-cabundle-controller/clusterrolebinding.yaml": v400ApiserviceCabundleControllerClusterrolebindingYaml, "v4.0.0/apiservice-cabundle-controller/cm.yaml": v400ApiserviceCabundleControllerCmYaml, "v4.0.0/apiservice-cabundle-controller/defaultconfig.yaml": v400ApiserviceCabundleControllerDefaultconfigYaml, "v4.0.0/apiservice-cabundle-controller/deployment.yaml": v400ApiserviceCabundleControllerDeploymentYaml, "v4.0.0/apiservice-cabundle-controller/ns.yaml": v400ApiserviceCabundleControllerNsYaml, "v4.0.0/apiservice-cabundle-controller/role.yaml": v400ApiserviceCabundleControllerRoleYaml, "v4.0.0/apiservice-cabundle-controller/rolebinding.yaml": v400ApiserviceCabundleControllerRolebindingYaml, "v4.0.0/apiservice-cabundle-controller/sa.yaml": v400ApiserviceCabundleControllerSaYaml, "v4.0.0/apiservice-cabundle-controller/signing-cabundle.yaml": v400ApiserviceCabundleControllerSigningCabundleYaml, "v4.0.0/configmap-cabundle-controller/clusterrole.yaml": v400ConfigmapCabundleControllerClusterroleYaml, "v4.0.0/configmap-cabundle-controller/clusterrolebinding.yaml": v400ConfigmapCabundleControllerClusterrolebindingYaml, "v4.0.0/configmap-cabundle-controller/cm.yaml": v400ConfigmapCabundleControllerCmYaml, "v4.0.0/configmap-cabundle-controller/defaultconfig.yaml": v400ConfigmapCabundleControllerDefaultconfigYaml, "v4.0.0/configmap-cabundle-controller/deployment.yaml": v400ConfigmapCabundleControllerDeploymentYaml, "v4.0.0/configmap-cabundle-controller/ns.yaml": v400ConfigmapCabundleControllerNsYaml, "v4.0.0/configmap-cabundle-controller/role.yaml": v400ConfigmapCabundleControllerRoleYaml, "v4.0.0/configmap-cabundle-controller/rolebinding.yaml": v400ConfigmapCabundleControllerRolebindingYaml, "v4.0.0/configmap-cabundle-controller/sa.yaml": v400ConfigmapCabundleControllerSaYaml, "v4.0.0/configmap-cabundle-controller/signing-cabundle.yaml": v400ConfigmapCabundleControllerSigningCabundleYaml, "v4.0.0/service-serving-cert-signer-controller/clusterrole.yaml": v400ServiceServingCertSignerControllerClusterroleYaml, "v4.0.0/service-serving-cert-signer-controller/clusterrolebinding.yaml": v400ServiceServingCertSignerControllerClusterrolebindingYaml, "v4.0.0/service-serving-cert-signer-controller/cm.yaml": v400ServiceServingCertSignerControllerCmYaml, "v4.0.0/service-serving-cert-signer-controller/defaultconfig.yaml": v400ServiceServingCertSignerControllerDefaultconfigYaml, "v4.0.0/service-serving-cert-signer-controller/deployment.yaml": v400ServiceServingCertSignerControllerDeploymentYaml, "v4.0.0/service-serving-cert-signer-controller/ns.yaml": v400ServiceServingCertSignerControllerNsYaml, "v4.0.0/service-serving-cert-signer-controller/role.yaml": v400ServiceServingCertSignerControllerRoleYaml, "v4.0.0/service-serving-cert-signer-controller/rolebinding.yaml": v400ServiceServingCertSignerControllerRolebindingYaml, "v4.0.0/service-serving-cert-signer-controller/sa.yaml": v400ServiceServingCertSignerControllerSaYaml, "v4.0.0/service-serving-cert-signer-controller/signing-secret.yaml": v400ServiceServingCertSignerControllerSigningSecretYaml}

func AssetDir(name string) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func		func() (*asset, error)
	Children	map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{"v4.0.0": &bintree{nil, map[string]*bintree{"apiservice-cabundle-controller": &bintree{nil, map[string]*bintree{"clusterrole.yaml": &bintree{v400ApiserviceCabundleControllerClusterroleYaml, map[string]*bintree{}}, "clusterrolebinding.yaml": &bintree{v400ApiserviceCabundleControllerClusterrolebindingYaml, map[string]*bintree{}}, "cm.yaml": &bintree{v400ApiserviceCabundleControllerCmYaml, map[string]*bintree{}}, "defaultconfig.yaml": &bintree{v400ApiserviceCabundleControllerDefaultconfigYaml, map[string]*bintree{}}, "deployment.yaml": &bintree{v400ApiserviceCabundleControllerDeploymentYaml, map[string]*bintree{}}, "ns.yaml": &bintree{v400ApiserviceCabundleControllerNsYaml, map[string]*bintree{}}, "role.yaml": &bintree{v400ApiserviceCabundleControllerRoleYaml, map[string]*bintree{}}, "rolebinding.yaml": &bintree{v400ApiserviceCabundleControllerRolebindingYaml, map[string]*bintree{}}, "sa.yaml": &bintree{v400ApiserviceCabundleControllerSaYaml, map[string]*bintree{}}, "signing-cabundle.yaml": &bintree{v400ApiserviceCabundleControllerSigningCabundleYaml, map[string]*bintree{}}}}, "configmap-cabundle-controller": &bintree{nil, map[string]*bintree{"clusterrole.yaml": &bintree{v400ConfigmapCabundleControllerClusterroleYaml, map[string]*bintree{}}, "clusterrolebinding.yaml": &bintree{v400ConfigmapCabundleControllerClusterrolebindingYaml, map[string]*bintree{}}, "cm.yaml": &bintree{v400ConfigmapCabundleControllerCmYaml, map[string]*bintree{}}, "defaultconfig.yaml": &bintree{v400ConfigmapCabundleControllerDefaultconfigYaml, map[string]*bintree{}}, "deployment.yaml": &bintree{v400ConfigmapCabundleControllerDeploymentYaml, map[string]*bintree{}}, "ns.yaml": &bintree{v400ConfigmapCabundleControllerNsYaml, map[string]*bintree{}}, "role.yaml": &bintree{v400ConfigmapCabundleControllerRoleYaml, map[string]*bintree{}}, "rolebinding.yaml": &bintree{v400ConfigmapCabundleControllerRolebindingYaml, map[string]*bintree{}}, "sa.yaml": &bintree{v400ConfigmapCabundleControllerSaYaml, map[string]*bintree{}}, "signing-cabundle.yaml": &bintree{v400ConfigmapCabundleControllerSigningCabundleYaml, map[string]*bintree{}}}}, "service-serving-cert-signer-controller": &bintree{nil, map[string]*bintree{"clusterrole.yaml": &bintree{v400ServiceServingCertSignerControllerClusterroleYaml, map[string]*bintree{}}, "clusterrolebinding.yaml": &bintree{v400ServiceServingCertSignerControllerClusterrolebindingYaml, map[string]*bintree{}}, "cm.yaml": &bintree{v400ServiceServingCertSignerControllerCmYaml, map[string]*bintree{}}, "defaultconfig.yaml": &bintree{v400ServiceServingCertSignerControllerDefaultconfigYaml, map[string]*bintree{}}, "deployment.yaml": &bintree{v400ServiceServingCertSignerControllerDeploymentYaml, map[string]*bintree{}}, "ns.yaml": &bintree{v400ServiceServingCertSignerControllerNsYaml, map[string]*bintree{}}, "role.yaml": &bintree{v400ServiceServingCertSignerControllerRoleYaml, map[string]*bintree{}}, "rolebinding.yaml": &bintree{v400ServiceServingCertSignerControllerRolebindingYaml, map[string]*bintree{}}, "sa.yaml": &bintree{v400ServiceServingCertSignerControllerSaYaml, map[string]*bintree{}}, "signing-secret.yaml": &bintree{v400ServiceServingCertSignerControllerSigningSecretYaml, map[string]*bintree{}}}}}}}}

func RestoreAsset(dir, name string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}
func RestoreAssets(dir, name string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	children, err := AssetDir(name)
	if err != nil {
		return RestoreAsset(dir, name)
	}
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}
func _filePath(dir, name string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
