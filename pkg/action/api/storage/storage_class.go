package storage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yametech/yamecloud/pkg/action/api/common"
	"github.com/yametech/yamecloud/pkg/action/service"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"net/http"
)

// Get StorageClass
func (s *storageServer) GetStorageClass(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	if namespace == "" || name == "" {
		common.RequestParametersError(g, fmt.Errorf("params not obtain namespace=%s name=%s", namespace, name))
		return
	}
	item, err := s.StorageClass.Get(namespace, name)
	if err != nil {
		common.InternalServerError(g, err, err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// Subscribe StorageClass
func (s *storageServer) ListStorageClass(g *gin.Context) {
	list, err := s.StorageClass.List(g.Param("namespace"), "")
	if err != nil {
		common.InternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, list)
}

// Update or Create StorageClass
func (s *storageServer) ApplyStorageClass(g *gin.Context) {
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
	newUnstructuredExtend, isUpdate, err := s.StorageClass.Apply(namespace, name, &service.UnstructuredExtend{Unstructured: _unstructured})
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

// Delete StorageClass
func (s *storageServer) DeleteStorageClass(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	if namespace == "" || name == "" {
		common.RequestParametersError(g, fmt.Errorf("params not obtain namespace=%s name=%s", namespace, name))
		return
	}
	err := s.StorageClass.Delete(namespace, name)
	if err != nil {
		common.InternalServerError(g, err, err)
		return
	}
	g.JSON(http.StatusOK, nil)
}
