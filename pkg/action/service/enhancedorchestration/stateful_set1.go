package enhancedorchestration

import (
	"github.com/yametech/yamecloud/pkg/action/service"
	"github.com/yametech/yamecloud/pkg/k8s"
)

var _ service.IResourceService = &StatefulSet1{}

type StatefulSet1 struct {
	service.Interface
}

func NewStatefulSet1(svcInterface service.Interface) *StatefulSet1 {
	statefulSet1 := &StatefulSet1{Interface: svcInterface}
	svcInterface.Install(k8s.StatefulSet1, statefulSet1)
	return statefulSet1
}

func (g *StatefulSet1) Get(namespace, name string) (*service.UnstructuredExtend, error) {
	item, err := g.Interface.Get(namespace, k8s.StatefulSet1, name)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (g *StatefulSet1) List(namespace string, selector string) (*service.UnstructuredListExtend, error) {
	list, err := g.Interface.List(namespace, k8s.StatefulSet1, selector)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (g *StatefulSet1) Apply(namespace, name string, unstructuredExtend *service.UnstructuredExtend) (*service.UnstructuredExtend, bool, error) {
	item, isUpdate, err := g.Interface.Apply(namespace, k8s.StatefulSet1, name, unstructuredExtend)
	if err != nil {
		return nil, isUpdate, err
	}
	return item, isUpdate, nil
}

func (g *StatefulSet1) Delete(namespace, name string) error {
	err := g.Interface.Delete(namespace, k8s.StatefulSet1, name)
	if err != nil {
		return err
	}
	return nil
}
