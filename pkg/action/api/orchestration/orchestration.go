package orchestration

import (
	"fmt"
	"github.com/yametech/yamecloud/pkg/action/api"
	"github.com/yametech/yamecloud/pkg/action/service/orchestration"
)

type orchestrationServer struct {
	name string
	*api.Server
	// action services
	*orchestration.Pod
	*orchestration.Deployment
	*orchestration.StatefulSet
	*orchestration.DaemonSet
	*orchestration.ReplicaSet
}

func (s *orchestrationServer) Name() string { return s.name }

func NewOrchestrationServer(serviceName string, server *api.Server) *orchestrationServer {
	orchestrationServer := &orchestrationServer{
		name:        serviceName,
		Server:      server,
		Pod:         orchestration.NewPod(server.Interface),
		Deployment:  orchestration.NewDeployment(server.Interface),
		StatefulSet: orchestration.NewStatefulSet(server.Interface),
		DaemonSet:   orchestration.NewDaemonSet(server.Interface),
		ReplicaSet:  orchestration.NewReplicaSet(server.Interface),
	}
	group := orchestrationServer.Group(fmt.Sprintf("/%s", serviceName))

	//Pod
	{
		group.GET("/api/v1/pods", orchestrationServer.ListPod)
		group.GET("/api/v1/namespaces/:namespace/pods", orchestrationServer.ListPod)
		group.GET("/api/v1/namespaces/:namespace/pods/:name", orchestrationServer.GetPod)
		group.GET("/api/v1/namespaces/:namespace/pods/:name/log", orchestrationServer.LogPod)
		group.POST("/api/v1/namespaces/:namespace/pods", orchestrationServer.ApplyPod)
		group.DELETE("/api/v1/namespaces/:namespace/pods/:name", orchestrationServer.DeletePod)
	}

	//DaemonSet
	{
		group.GET("/apis/apps/v1/daemonsets", orchestrationServer.ListDaemonSet)
		group.GET("/apis/apps/v1/namespaces/:namespace/daemonsets", orchestrationServer.ListDaemonSet)
		group.GET("/apis/apps/v1/namespaces/:namespace/daemonsets/:name", orchestrationServer.GetDaemonSet)
		group.POST("/apis/apps/v1/namespaces/:namespace/daemonsets", orchestrationServer.ApplyDaemonSet)
		group.DELETE("/apis/apps/v1/namespaces/:namespace/daemonsets/:name", orchestrationServer.DeleteDaemonSet)
	}

	//Deployment
	{
		group.GET("/apis/apps/v1/deployments", orchestrationServer.ListDeployment)
		group.GET("/apis/apps/v1/namespaces/:namespace/deployments", orchestrationServer.ListDeployment)
		group.GET("/apis/apps/v1/namespaces/:namespace/deployments/:name", orchestrationServer.GetDeployment)
		group.POST("/apis/apps/v1/namespaces/:namespace/deployments", orchestrationServer.ApplyDeployment)
		group.DELETE("/apis/apps/v1/namespaces/:namespace/deployments/:name", orchestrationServer.DeleteDeployment)
	}

	//ReplicaSet
	{
		group.GET("/apis/apps/v1/replicasets", orchestrationServer.ListReplicaSet)
		group.GET("/apis/apps/v1/namespaces/:namespace/replicasets", orchestrationServer.ListReplicaSet)
		group.GET("/apis/apps/v1/namespaces/:namespace/replicasets/:name", orchestrationServer.GetReplicaSet)
		group.POST("/apis/apps/v1/namespaces/:namespace/replicasets", orchestrationServer.ApplyReplicaSet)
		group.DELETE("/apis/apps/v1/namespaces/:namespace/replicasets/:name", orchestrationServer.DeleteReplicaSet)
	}

	//StatefulSet
	{
		group.GET("/apis/apps/v1/statefulsets", orchestrationServer.ListStatefulSet)
		group.GET("/apis/apps/v1/namespaces/:namespace/statefulsets", orchestrationServer.ListStatefulSet)
		group.GET("/apis/apps/v1/namespaces/:namespace/statefulsets/:name", orchestrationServer.GetStatefulSet)
		group.POST("/apis/apps/v1/namespaces/:namespace/statefulsets", orchestrationServer.ApplyStatefulSet)
		group.DELETE("/apis/apps/v1/namespaces/:namespace/statefulsets/:name", orchestrationServer.DeleteStatefulSet)
	}

	return orchestrationServer
}
