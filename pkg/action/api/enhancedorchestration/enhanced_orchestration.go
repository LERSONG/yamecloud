package enhancedorchestration

import (
	"fmt"
	"github.com/yametech/yamecloud/pkg/action/api"
	"github.com/yametech/yamecloud/pkg/action/service/enhancedorchestration"
)

type enhancedOrchestrationServer struct {
	name string
	*api.Server
	// action services
	*enhancedorchestration.Stone
	*enhancedorchestration.StatefulSet1
	*enhancedorchestration.Water
	*enhancedorchestration.Injector
}

func (s *enhancedOrchestrationServer) Name() string { return s.name }

func NewEnhancedOrchestrationServer(serviceName string, server *api.Server) *enhancedOrchestrationServer {
	enhancedOrchestrationServer := &enhancedOrchestrationServer{
		name:         serviceName,
		Server:       server,
		Stone:        enhancedorchestration.NewStone(server.Interface),
		StatefulSet1: enhancedorchestration.NewStatefulSet1(server.Interface),
		Water:        enhancedorchestration.NewWater(server.Interface),
		Injector:     enhancedorchestration.NewInjector(server.Interface),
	}
	group := enhancedOrchestrationServer.Group(fmt.Sprintf("/%s", serviceName))

	//Stone
	{
		group.GET("/apis/nuwa.nip.io/v1/stones", enhancedOrchestrationServer.ListStone)
		group.GET("/apis/nuwa.nip.io/v1/namespaces/:namespace/stones", enhancedOrchestrationServer.ListStone)
		group.GET("/apis/nuwa.nip.io/v1/namespaces/:namespace/stones/:name", enhancedOrchestrationServer.GetStone)
		group.POST("/apis/nuwa.nip.io/v1/namespaces/:namespace/stones", enhancedOrchestrationServer.ApplyStone)
		group.DELETE("/apis/nuwa.nip.io/v1/namespaces/:namespace/stones/:name", enhancedOrchestrationServer.DeleteStone)
	}

	// Water
	{
		group.GET("/apis/nuwa.nip.io/v1/waters", enhancedOrchestrationServer.ListWater)
		group.GET("/apis/nuwa.nip.io/v1/namespaces/:namespace/waters", enhancedOrchestrationServer.ListWater)
		group.GET("/apis/nuwa.nip.io/v1/namespaces/:namespace/waters/:name", enhancedOrchestrationServer.GetWater)
		group.POST("/apis/nuwa.nip.io/v1/namespaces/:namespace/waters", enhancedOrchestrationServer.ApplyWater)
		group.DELETE("/apis/nuwa.nip.io/v1/namespaces/:namespace/waters/:name", enhancedOrchestrationServer.DeleteWater)
	}

	// Injector
	{
		group.GET("/apis/nuwa.nip.io/v1/injectors", enhancedOrchestrationServer.ListInjector)
		group.GET("/apis/nuwa.nip.io/v1/namespaces/:namespace/injectors", enhancedOrchestrationServer.ListInjector)
		group.GET("/apis/nuwa.nip.io/v1/namespaces/:namespace/injectors/:name", enhancedOrchestrationServer.GetInjector)
		group.POST("/apis/nuwa.nip.io/v1/namespaces/:namespace/injectors", enhancedOrchestrationServer.ApplyInjector)
		group.DELETE("/apis/nuwa.nip.io/v1/namespaces/:namespace/injectors/:name", enhancedOrchestrationServer.DeleteInjector)
	}

	// StatefulSet1
	{
		group.GET("/apis/nuwa.nip.io/v1/statefulsets", enhancedOrchestrationServer.ListStatefulSet1)
		group.GET("/apis/nuwa.nip.io/v1/namespaces/:namespace/statefulsets", enhancedOrchestrationServer.ListStatefulSet1)
		group.GET("/apis/nuwa.nip.io/v1/namespaces/:namespace/statefulsets/:name", enhancedOrchestrationServer.GetStatefulSet1)
		group.POST("/apis/nuwa.nip.io/v1/namespaces/:namespace/statefulsets", enhancedOrchestrationServer.ApplyStatefulSet1)
		group.DELETE("/apis/nuwa.nip.io/v1/namespaces/:namespace/statefulsets/:name", enhancedOrchestrationServer.DeleteStatefulSet1)
	}

	return enhancedOrchestrationServer
}
