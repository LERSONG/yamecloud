package magpie

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yametech/yamecloud/pkg/action/api"
	"k8s.io/client-go/tools/remotecommand"
	"net/http"
)

type magpieServer struct {
	name string
	*api.Server
}

func (s *magpieServer) Name() string { return s.name }

func NewMagpieServer(serviceName string, server *api.Server) *magpieServer {
	magpieServer := &magpieServer{
		name:   serviceName,
		Server: server,
	}
	group := magpieServer.Group(fmt.Sprintf("/%s", serviceName))
	serveHttp := WrapH(CreateAttachHandler("/magpie/shell/pod"))
	group.GET("/shell/pod/*path", serveHttp)
	group.GET("/attach/namespace/:namespace/pod/:name/container/:container/:shelltype", magpieServer.AttachPod)

	return magpieServer
}

func WrapH(h http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers,X-Access-Token,XKey,Authorization")

		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Attach Pod
func (s *magpieServer) AttachPod(g *gin.Context) {
	attachPodRequest := &AttachPodRequest{
		Namespace: g.Param("namespace"),
		Name:      g.Param("name"),
		Container: g.Param("container"),
		ShellType: g.Param("shelltype"),
		Shell:     g.Query("shell"),
	}

	sessionId, _ := generateTerminalSessionId()
	sharedSessionManager.set(sessionId,
		&sessionChannels{
			id:       sessionId,
			bound:    make(chan struct{}),
			sizeChan: make(chan remotecommand.TerminalSize),
		})

	//waitForTerminal(attachPodRequest, sessionId)
	go waitForTerminal(attachPodRequest, sessionId)
	g.JSON(http.StatusOK, gin.H{"op": BIND, "sessionId": sessionId})
}
