package servicemesh

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yametech/yamecloud/pkg/action/api/common"
	"github.com/yametech/yamecloud/pkg/action/service"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"net/http"
)

// Get VirtualService
func (s *serviceMeshServer) GetVirtualService(g *gin.Context) {
	name := g.Param("name")
	if name == "" {
		common.RequestParametersError(g, fmt.Errorf("params not obtain name=%s", name))
		return
	}
	item, err := s.VirtualService.Get("", name)
	if err != nil {
		common.InternalServerError(g, err, err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// Subscribe VirtualService
func (s *serviceMeshServer) ListVirtualService(g *gin.Context) {
	list, err := s.VirtualService.List(g.Param("namespace"), "")
	if err != nil {
		common.InternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, list)
}

// Update or Create VirtualService
func (s *serviceMeshServer) ApplyVirtualService(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	if namespace == "" || name == "" {
		common.RequestParametersError(g, fmt.Errorf("params not obtain namespace=%s name=%s", namespace, name))
		return
	}
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
	newUnstructuredExtend, err := s.VirtualService.Apply(namespace, name, &service.UnstructuredExtend{Unstructured: _unstructured})
	if err != nil {
		common.InternalServerError(g, newUnstructuredExtend, fmt.Errorf("apply object error (%s)", err))
		return
	}

	g.JSON(http.StatusOK, newUnstructuredExtend)
}
