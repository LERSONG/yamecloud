package storage

import (
	"github.com/yametech/yamecloud/pkg/action/service"
	"github.com/yametech/yamecloud/pkg/k8s"
)

var _ service.IResourceService = &CephRBDStorage{}

type CephRBDStorage struct {
	service.Interface
}

func NewCephRBDStorage(svcInterface service.Interface) *CephRBDStorage {
	cephRBDStorage := &CephRBDStorage{Interface: svcInterface}
	svcInterface.Install(k8s.CephRBDStorage, cephRBDStorage)
	return cephRBDStorage
}

func (g *CephRBDStorage) Get(namespace, name string) (*service.UnstructuredExtend, error) {
	item, err := g.Interface.Get(namespace, k8s.CephRBDStorage, name)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (g *CephRBDStorage) List(namespace string, selector string) (*service.UnstructuredListExtend, error) {
	list, err := g.Interface.List(namespace, k8s.CephRBDStorage, selector)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (g *CephRBDStorage) Apply(namespace, name string, unstructuredExtend *service.UnstructuredExtend) (*service.UnstructuredExtend, bool, error) {
	item, isUpdate, err := g.Interface.Apply(namespace, k8s.CephRBDStorage, name, unstructuredExtend)
	if err != nil {
		return nil, isUpdate, err
	}
	return item, isUpdate, nil
}

func (g *CephRBDStorage) Delete(namespace, name string) error {
	err := g.Interface.Delete(namespace, k8s.CephRBDStorage, name)
	if err != nil {
		return err
	}
	return nil
}
