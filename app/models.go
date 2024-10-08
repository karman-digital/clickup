package clickupapp

import (
	"github.com/karman-digital/clickup/api/credentials"
	"github.com/karman-digital/clickup/api/timetracking"
	clickupintefaces "github.com/karman-digital/clickup/app/interfaces"
)

type ApiClient struct {
	TimeTracking timetracking.TimeTracking
	Tasks        clickupintefaces.Tasks
	Lists        clickupintefaces.Lists
}

type ClickUp struct {
	*credentials.Credentials
	ApiClient
}
