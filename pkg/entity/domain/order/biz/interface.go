//go:generate mockgen -destination=./mock_${GOFILE} -package=biz -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/irent/pkg/contextx"
	am "github.com/blackhorseya/irent/pkg/entity/domain/account/model"
	om "github.com/blackhorseya/irent/pkg/entity/domain/order/model"
	rm "github.com/blackhorseya/irent/pkg/entity/domain/rental/model"
)

// IBiz declare order biz functions
type IBiz interface {
	// Readiness probes to know when a container is ready to start accepting traffic. A Pod is considered ready when all of its containers are ready. One use of this signal is to control which Pods are used as backends for Services. When a Pod is not ready, it is removed from Service load balancers.
	Readiness(ctx contextx.Contextx) error

	// Liveness probes to know when to restart a container. For example, liveness probes could catch a deadlock, where an application is running, but unable to make progress. Restarting a container in such a state can help to make the application more available despite bugs.
	Liveness(ctx contextx.Contextx) error

	// GetArrears serve caller to given user to get user's arrears information
	GetArrears(ctx contextx.Contextx, from *am.Profile, target *am.Profile) (info *om.Arrears, err error)

	// ListLease serve caller to given user to get user's lease
	ListLease(ctx contextx.Contextx, from *am.Profile) (orders []*om.Lease, err error)

	// BookRental serve caller to book a rental
	BookRental(ctx contextx.Contextx, from *am.Profile, target *rm.Car) (info *om.Lease, err error)

	// CancelLease serve caller to given user and order's id then to cancel the booking
	CancelLease(ctx contextx.Contextx, from *am.Profile, target *om.Lease) error
}
