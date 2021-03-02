package orchestration

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yametech/yamecloud/pkg/action/api/common"
	"github.com/yametech/yamecloud/pkg/action/service"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"net/http"
	"time"
)

type logRequest struct {
	Container  string    `form:"container" json:"container"`
	Timestamps bool      `form:"timestamps" json:"timestamps"`
	SinceTime  time.Time `form:"sinceTime" json:"sinceTime"`
	TailLines  int64     `form:"tailLines" json:"tailLines"`
}

// Get Pod
func (s *orchestrationServer) GetPod(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	if namespace == "" || name == "" {
		common.RequestParametersError(g, fmt.Errorf("params not obtain namespace=%s name=%s", namespace, name))
		return
	}
	item, err := s.Pod.Get(namespace, name)
	if err != nil {
		common.InternalServerError(g, err, err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// Subscribe Pod
func (s *orchestrationServer) ListPod(g *gin.Context) {
	list, err := s.Pod.List(g.Param("namespace"), "")
	if err != nil {
		common.InternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, list)
}

// Update or Create Pod
func (s *orchestrationServer) ApplyPod(g *gin.Context) {
	namespace := g.Param("namespace")
	raw, err := g.GetRawData()
	if err != nil {
		common.RequestParametersError(g, fmt.Errorf("get raw data error (%s)", err))
		return
	}

	_unstructured := &unstructured.Unstructured{}
	if err := _unstructured.UnmarshalJSON(raw); err != nil {
		common.RequestParametersError(g, fmt.Errorf("unmarshal from json data error (%s)", err))
		return
	}
	name := _unstructured.GetName()
	newUnstructuredExtend, isUpdate, err := s.Pod.Apply(namespace, name, &service.UnstructuredExtend{Unstructured: _unstructured})
	if err != nil {
		common.InternalServerError(g, newUnstructuredExtend, fmt.Errorf("apply object error (%s)", err))
		return
	}

	if isUpdate {
		g.JSON(
			http.StatusOK,
			[]service.UnstructuredExtend{
				*newUnstructuredExtend,
			})
	} else {
		g.JSON(http.StatusOK, newUnstructuredExtend)
	}
}

// Delete Pod
func (s *orchestrationServer) DeletePod(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	if namespace == "" || name == "" {
		common.RequestParametersError(g, fmt.Errorf("params not obtain namespace=%s name=%s", namespace, name))
		return
	}
	err := s.Pod.Delete(namespace, name)
	if err != nil {
		common.InternalServerError(g, err, err)
		return
	}
	g.JSON(http.StatusOK, nil)
}

// Log Pod
func (s *orchestrationServer) LogPod(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	lq := &logRequest{}
	if err := g.Bind(lq); err != nil || namespace == "" || name == "" {
		common.RequestParametersError(g, err)
		return
	}

	buf := bytes.NewBufferString("")
	err := s.Pod.Logs(
		namespace,
		name,
		lq.Container,
		false,
		false,
		lq.Timestamps,
		0,
		&lq.SinceTime,
		0,
		lq.TailLines,
		buf,
	)
	if err != nil {
		common.RequestParametersError(g, err)
		return
	}

	g.JSON(http.StatusOK, buf.String())
}
