package job

import (
	"github.com/yametech/yamecloud/pkg/action/service"
	"github.com/yametech/yamecloud/pkg/k8s"
)

var _ service.IResourceService = &CronJob{}

type CronJob struct {
	service.Interface
}

func NewCronJob(svcInterface service.Interface) *CronJob {
	cronJob := &CronJob{Interface: svcInterface}
	svcInterface.Install(k8s.CronJobs, cronJob)
	return cronJob
}

func (b *CronJob) Get(namespace, name string) (*service.UnstructuredExtend, error) {
	item, err := b.Interface.Get(namespace, k8s.CronJobs, name)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (b *CronJob) List(namespace string, selector string) (*service.UnstructuredListExtend, error) {
	list, err := b.Interface.List(namespace, k8s.CronJobs, selector)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (b *CronJob) Apply(namespace, name string, unstructuredExtend *service.UnstructuredExtend) (*service.UnstructuredExtend, bool, error) {
	item, isUpdate, err := b.Interface.Apply(namespace, k8s.CronJobs, name, unstructuredExtend)
	if err != nil {
		return nil, isUpdate, err
	}
	return item, isUpdate, nil
}

func (b *CronJob) Delete(namespace, name string) error {
	err := b.Interface.Delete(namespace, k8s.CronJobs, name)
	if err != nil {
		return err
	}
	return nil
}
