package repo

import (
	"time"

	rm "github.com/blackhorseya/irent/pkg/entity/domain/rental/model"
)

type car struct {
	ID          string       `db:"id"`
	Area        string       `db:"area"`
	ProjectID   string       `db:"project_id"`
	ProjectName string       `db:"project_name"`
	Seat        int64        `db:"seat"`
	TypeName    string       `db:"type_name"`
	latitude    float64      `db:"latitude"`
	longitude   float64      `db:"longitude"`
	Status      rm.CarStatus `db:"status"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   time.Time    `db:"updated_at"`
}
