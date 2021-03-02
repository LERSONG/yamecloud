package enhancedorchestration

import (
	"github.com/yametech/yamecloud/pkg/action/service"
	"github.com/yametech/yamecloud/pkg/k8s"
)

var _ service.IResourceService = &Water{}

type Water struct {
	service.Interface
}

func NewWater(svcInterface service.Interface) *Water {
	water := &Water{Interface: svcInterface}
	svcInterface.Install(k8s.Water, water)
	return water
}

func (g *Water) Get(namespace, name string) (*service.UnstructuredExtend, error) {
	item, err := g.Interface.Get(namespace, k8s.Water, name)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (g *Water) List(namespace string, selector string) (*service.UnstructuredListExtend, error) {
	list, err := g.Interface.List(namespace, k8s.Water, selector)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (g *Water) Apply(namespace, name string, unstructuredExtend *service.UnstructuredExtend) (*service.UnstructuredExtend, bool, error) {
	item, isUpdate, err := g.Interface.Apply(namespace, k8s.Water, name, unstructuredExtend)
	if err != nil {
		return nil, isUpdate, err
	}
	return item, isUpdate, nil
}

func (g *Water) Delete(namespace, name string) error {
	err := g.Interface.Delete(namespace, k8s.Water, name)
	if err != nil {
		return err
	}
	return nil
}
