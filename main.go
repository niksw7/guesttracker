package main
import (
	"fmt"
	"time"

	"contrib.go.opencensus.io/exporter/ocagent"
	"go.opencensus.io/trace"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {


	ocagentHost := "oc-collector.tracing:55678"
	oce, _ := ocagent.NewExporter(
		ocagent.WithInsecure(),
		ocagent.WithReconnectionPeriod(1*time.Second),
		ocagent.WithAddress(ocagentHost),
		ocagent.WithServiceName("guesttracker"))

	trace.RegisterExporter(oce)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	r := gin.Default()
	r.POST("/track-guest", func(c *gin.Context) {
		fmt.Println(c.Request.Header)
		fmt.Println(c.Request.Host)
		fmt.Println("Tracking the user")
		var json LoginData
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"Added": json})
	})
	r.Run(":8081") // listen and serve on 0.0.0.0:8081
}



type LoginData struct {
	UserName     string `json:"username"`
	Email string `json:"email"`
}

