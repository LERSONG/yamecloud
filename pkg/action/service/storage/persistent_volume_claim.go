package storage

import (
	"github.com/yametech/yamecloud/pkg/action/service"
	"github.com/yametech/yamecloud/pkg/k8s"
)

var _ service.IResourceService = &PersistentVolumeClaim{}

type PersistentVolumeClaim struct {
	service.Interface
}

func NewPersistentVolumeClaim(svcInterface service.Interface) *PersistentVolumeClaim {
	PersistentVolumeClaim := &PersistentVolumeClaim{Interface: svcInterface}
	svcInterface.Install(k8s.PersistentVolumeClaim, PersistentVolumeClaim)
	return PersistentVolumeClaim
}

func (g *PersistentVolumeClaim) Get(namespace, name string) (*service.UnstructuredExtend, error) {
	item, err := g.Interface.Get(namespace, k8s.PersistentVolumeClaim, name)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (g *PersistentVolumeClaim) List(namespace string, selector string) (*service.UnstructuredListExtend, error) {
	list, err := g.Interface.List(namespace, k8s.PersistentVolumeClaim, selector)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (g *PersistentVolumeClaim) Apply(namespace, name string, unstructuredExtend *service.UnstructuredExtend) (*service.UnstructuredExtend, bool, error) {
	item, isUpdate, err := g.Interface.Apply(namespace, k8s.PersistentVolumeClaim, name, unstructuredExtend)
	if err != nil {
		return nil, isUpdate, err
	}
	return item, isUpdate, nil
}

func (g *PersistentVolumeClaim) Delete(namespace, name string) error {
	err := g.Interface.Delete(namespace, k8s.PersistentVolumeClaim, name)
	if err != nil {
		return err
	}
	return nil
}
