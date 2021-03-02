package job

import (
	"fmt"
	"github.com/yametech/yamecloud/pkg/action/api"
	"github.com/yametech/yamecloud/pkg/action/service/job"
)

type jobServer struct {
	name string
	*api.Server
	// action services
	*job.Job
	*job.CronJob
}

func (s *jobServer) Name() string { return s.name }

func NewJobServer(serviceName string, server *api.Server) *jobServer {
	jobServer := &jobServer{
		name:    serviceName,
		Server:  server,
		Job:     job.NewJob(server.Interface),
		CronJob: job.NewCronJob(server.Interface),
	}
	group := jobServer.Group(fmt.Sprintf("/%s", serviceName))

	// Job
	{
		group.GET("/apis/batch/v1/jobs", jobServer.ListJob)
		group.GET("/apis/batch/v1/namespaces/:namespace/jobs", jobServer.ListJob)
		group.GET("/apis/batch/v1/namespaces/:namespace/jobs/:name", jobServer.GetJob)
		group.POST("/apis/batch/v1/namespaces/:namespace/jobs", jobServer.ApplyJob)
		group.DELETE("/apis/batch/v1/namespaces/:namespace/jobs/:name", jobServer.DeleteJob)
	}

	// CronJob
	{
		group.GET("/apis/batch/v1beta1/cronjobs", jobServer.ListCronJob)
		group.GET("/apis/batch/v1beta1/namespaces/:namespace/cronjobs", jobServer.ListCronJob)
		group.GET("/apis/batch/v1beta1/namespaces/:namespace/cronjobs/:name", jobServer.GetCronJob)
		group.POST("/apis/batch/v1beta1/namespaces/:namespace/cronjobs", jobServer.ApplyCronJob)
		group.DELETE("/apis/batch/v1beta1/namespaces/:namespace/cronjobs/:name", jobServer.DeleteCronJob)
	}

	return jobServer
}
