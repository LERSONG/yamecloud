package storage

import (
	"github.com/yametech/yamecloud/pkg/action/service"
	"github.com/yametech/yamecloud/pkg/k8s"
)

var _ service.IResourceService = &PersistentVolume{}

type PersistentVolume struct {
	service.Interface
}

func NewPersistentVolume(svcInterface service.Interface) *PersistentVolume {
	persistentVolume := &PersistentVolume{Interface: svcInterface}
	svcInterface.Install(k8s.PersistentVolume, persistentVolume)
	return persistentVolume
}

func (g *PersistentVolume) Get(namespace, name string) (*service.UnstructuredExtend, error) {
	item, err := g.Interface.Get(namespace, k8s.PersistentVolume, name)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (g *PersistentVolume) List(namespace string, selector string) (*service.UnstructuredListExtend, error) {
	list, err := g.Interface.List(namespace, k8s.PersistentVolume, selector)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (g *PersistentVolume) Apply(namespace, name string, unstructuredExtend *service.UnstructuredExtend) (*service.UnstructuredExtend, bool, error) {
	item, isUpdate, err := g.Interface.Apply(namespace, k8s.PersistentVolume, name, unstructuredExtend)
	if err != nil {
		return nil, isUpdate, err
	}
	return item, isUpdate, nil
}

func (g *PersistentVolume) Delete(namespace, name string) error {
	err := g.Interface.Delete(namespace, k8s.PersistentVolume, name)
	if err != nil {
		return err
	}
	return nil
}
