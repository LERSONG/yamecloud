package storage

import (
	"fmt"
	"github.com/yametech/yamecloud/pkg/action/api"
	"github.com/yametech/yamecloud/pkg/action/service/storage"
)

type storageServer struct {
	name string
	*api.Server
	// action services
	*storage.CephRBDStorage
	*storage.StorageClass
	*storage.PersistentVolumeClaim
	*storage.PersistentVolume
}

func (s *storageServer) Name() string { return s.name }

func NewStorageServer(serviceName string, server *api.Server) *storageServer {
	storageServer := &storageServer{
		name:                  serviceName,
		Server:                server,
		CephRBDStorage:        storage.NewCephRBDStorage(server.Interface),
		StorageClass:          storage.NewStorageClass(server.Interface),
		PersistentVolume:      storage.NewPersistentVolume(server.Interface),
		PersistentVolumeClaim: storage.NewPersistentVolumeClaim(server.Interface),
	}
	group := storageServer.Group(fmt.Sprintf("/%s", serviceName))

	// CephRBDStorage
	{
		group.GET("/apis/yamecloud.io/v1/cephrbdstorages", storageServer.ListCephRBDStorage)
		group.GET("/apis/yamecloud.io/v1/namespaces/:namespace/cephrbdstorages", storageServer.ListCephRBDStorage)
		group.GET("/apis/yamecloud.io/v1/namespaces/:namespace/cephrbdstorages/:name", storageServer.GetCephRBDStorage)
		group.POST("/apis/yamecloud.io/v1/namespaces/:namespace/cephrbdstorages", storageServer.ApplyCephRBDStorage)
		group.DELETE("/apis/yamecloud.io/v1/namespaces/:namespace/cephrbdstorages/:name", storageServer.DeleteCephRBDStorage)
	}

	// StorageClass
	{
		group.GET("/apis/storage.k8s.io/v1/storageclasses", storageServer.ListStorageClass)
		group.GET("/apis/storage.k8s.io/v1/namespaces/:namespace/storageclasses", storageServer.ListStorageClass)
		group.GET("/apis/storage.k8s.io/v1/namespaces/:namespace/storageclasses/:name", storageServer.GetStorageClass)
		group.POST("/apis/storage.k8s.io/v1/namespaces/:namespace/storageclasses", storageServer.ApplyStorageClass)
		group.DELETE("/apis/storage.k8s.io/v1/namespaces/:namespace/storageclasses/:name", storageServer.DeleteStorageClass)
	}

	// PersistentVolume
	{
		group.GET("/api/v1/persistentvolumes", storageServer.ListPersistentVolume)
		group.GET("/api/v1/namespaces/:namespace/persistentvolumes", storageServer.ListPersistentVolume)
		group.GET("/api/v1/namespaces/:namespace/persistentvolumes/:name", storageServer.GetPersistentVolume)
		group.POST("/api/v1/namespaces/:namespace/persistentvolumes", storageServer.ApplyPersistentVolume)
		group.DELETE("/api/v1/namespaces/:namespace/persistentvolumes/:name", storageServer.DeletePersistentVolume)
	}

	// PersistentVolumeClaim
	{
		group.GET("/api/v1/persistentvolumeclaims", storageServer.ListPersistentVolumeClaim)
		group.GET("/api/v1/namespaces/:namespace/persistentvolumeclaims", storageServer.ListPersistentVolumeClaim)
		group.GET("/api/v1/namespaces/:namespace/persistentvolumeclaims/:name", storageServer.GetPersistentVolumeClaim)
		group.POST("/api/v1/namespaces/:namespace/persistentvolumeclaims", storageServer.ApplyPersistentVolumeClaim)
		group.DELETE("/api/v1/namespaces/:namespace/persistentvolumeclaims/:name", storageServer.DeletePersistentVolumeClaim)
	}

	return storageServer
}
