package biz

import (
	"github.com/blackhorseya/irent/pkg/contextx"
	am "github.com/blackhorseya/irent/pkg/entity/domain/account/model"
)

// IBiz declare account biz functions
//
//go:generate mockery --all --inpackage
type IBiz interface {
	// Readiness probes to know when a container is ready to start accepting traffic. A Pod is considered ready when all of its containers are ready. One use of this signal is to control which Pods are used as backends for Services. When a Pod is not ready, it is removed from Service load balancers.
	Readiness(ctx contextx.Contextx) error

	// Liveness probes to know when to restart a container. For example, liveness probes could catch a deadlock, where an application is running, but unable to make progress. Restarting a container in such a state can help to make the application more available despite bugs.
	Liveness(ctx contextx.Contextx) error

	// Login serve caller to given id and password then login the system
	Login(ctx contextx.Contextx, id, password string) (info *am.Profile, err error)

	// GetByAccessToken serve caller to given access token to get user profile
	GetByAccessToken(ctx contextx.Contextx, token string) (info *am.Profile, err error)
}
