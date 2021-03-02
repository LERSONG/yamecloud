package orchestration

import (
	"context"
	"github.com/yametech/yamecloud/pkg/action/service"
	"github.com/yametech/yamecloud/pkg/k8s"
	"io"
	"strconv"
	"time"
)

var _ service.IResourceService = &Pod{}

type Pod struct {
	service.Interface
}

func NewPod(svcInterface service.Interface) *Pod {
	pod := &Pod{
		Interface: svcInterface,
	}
	svcInterface.Install(k8s.Pod, pod)
	return pod
}

func (g *Pod) Get(namespace, name string) (*service.UnstructuredExtend, error) {
	item, err := g.Interface.Get(namespace, k8s.Pod, name)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (g *Pod) List(namespace string, selector string) (*service.UnstructuredListExtend, error) {
	list, err := g.Interface.List(namespace, k8s.Pod, selector)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (g *Pod) Apply(namespace, name string, unstructuredExtend *service.UnstructuredExtend) (*service.UnstructuredExtend, bool, error) {
	item, isUpdate, err := g.Interface.Apply(namespace, k8s.Pod, name, unstructuredExtend)
	if err != nil {
		return nil, isUpdate, err
	}
	return item, isUpdate, nil
}

func (g *Pod) Delete(namespace, name string) error {
	err := g.Interface.Delete(namespace, k8s.Pod, name)
	if err != nil {
		return err
	}
	return nil
}

func (p *Pod) Logs(
	namespace, name, container string,
	follow bool, previous bool, timestamps bool,
	sinceSeconds int64,
	sinceTime *time.Time,
	limitBytes int64,
	tailLines int64,
	out io.Writer,
) error {
	req :=
		p.Interface.RESTClientSet().
			CoreV1().
			RESTClient().
			Get().
			Namespace(namespace).
			Name(name).
			Resource("pods").
			SubResource("log").
			Param("container", container).
			Param("follow", strconv.FormatBool(follow)).
			Param("previous", strconv.FormatBool(previous)).
			Param("timestamps", strconv.FormatBool(timestamps))

	if sinceSeconds != 0 {
		req.Param("sinceSeconds", strconv.FormatInt(sinceSeconds, 10))
	}
	if sinceTime != nil {
		req.Param("sinceTime", sinceTime.Format(time.RFC3339))
	}
	if limitBytes != 0 {
		req.Param("limitBytes", strconv.FormatInt(limitBytes, 10))
	}
	if tailLines != 0 {
		req.Param("tailLines", strconv.FormatInt(tailLines, 10))
	}
	readCloser, err := req.Stream(context.Background())
	if err != nil {
		return err
	}
	defer readCloser.Close()
	_, err = io.Copy(out, readCloser)

	return err
}
