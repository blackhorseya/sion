package biz

import (
	"github.com/blackhorseya/irent/pkg/contextx"
	rm "github.com/blackhorseya/irent/pkg/entity/domain/rental/model"
)

// IBiz declare rental biz functions
//
//go:generate mockery --all --inpackage
type IBiz interface {
	// Readiness probes to know when a container is ready to start accepting traffic. A Pod is considered ready when all of its containers are ready. One use of this signal is to control which Pods are used as backends for Services. When a Pod is not ready, it is removed from Service load balancers.
	Readiness(ctx contextx.Contextx) error

	// Liveness probes to know when to restart a container. For example, liveness probes could catch a deadlock, where an application is running, but unable to make progress. Restarting a container in such a state can help to make the application more available despite bugs.
	Liveness(ctx contextx.Contextx) error

	// ListCars serve caller to given condition to list cars
	ListCars(ctx contextx.Contextx, condition QueryCarCondition) (info []*rm.Car, total int, err error)

	// UpdateInfoCars serve caller to update information of all cars
	UpdateInfoCars(ctx contextx.Contextx) (cars []*rm.Car, err error)
}

// QueryCarCondition is a condition
type QueryCarCondition struct {
	TopNum    int
	Latitude  float64
	Longitude float64
}
