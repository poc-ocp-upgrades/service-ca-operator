package controller

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
)

type ParentFilter interface {
	Parent(obj metav1.Object) (namespace, name string)
	Filter
}
type Filter interface {
	Add(obj metav1.Object) bool
	Update(oldObj, newObj metav1.Object) bool
	Delete(obj metav1.Object) bool
}
type ParentFunc func(obj metav1.Object) (namespace, name string)
type FilterFuncs struct {
	ParentFunc	ParentFunc
	AddFunc		func(obj metav1.Object) bool
	UpdateFunc	func(oldObj, newObj metav1.Object) bool
	DeleteFunc	func(obj metav1.Object) bool
}

func (f FilterFuncs) Parent(obj metav1.Object) (namespace, name string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if f.ParentFunc == nil {
		return obj.GetNamespace(), obj.GetName()
	}
	return f.ParentFunc(obj)
}
func (f FilterFuncs) Add(obj metav1.Object) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if f.AddFunc == nil {
		return false
	}
	return f.AddFunc(obj)
}
func (f FilterFuncs) Update(oldObj, newObj metav1.Object) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if f.UpdateFunc == nil {
		return false
	}
	return f.UpdateFunc(oldObj, newObj)
}
func (f FilterFuncs) Delete(obj metav1.Object) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if f.DeleteFunc == nil {
		return false
	}
	return f.DeleteFunc(obj)
}
func FilterByNames(parentFunc ParentFunc, names ...string) ParentFilter {
	_logClusterCodePath()
	defer _logClusterCodePath()
	set := sets.NewString(names...)
	has := func(obj metav1.Object) bool {
		return set.Has(obj.GetName())
	}
	return FilterFuncs{ParentFunc: parentFunc, AddFunc: has, UpdateFunc: func(oldObj, newObj metav1.Object) bool {
		return has(newObj)
	}, DeleteFunc: has}
}
