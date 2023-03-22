package app

import (
	"github.com/andrewarrow/feedback/router"
)

func HandleDashboard(c *router.Context, second, third string) {
	if c.User == nil {
		c.UserRequired = true
		return
	}

	servers := c.SelectAll("server", "", []any{})
	c.SendContentInLayout("dashboard.html", servers, 200)
}
