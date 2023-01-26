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

func newCar(val *rm.Car) *car {
	return &car{
		ID:          val.Id,
		Area:        val.CarOfArea,
		ProjectID:   val.ProjectId,
		ProjectName: val.ProjectName,
		Seat:        val.Seat,
		TypeName:    val.CarTypeName,
		latitude:    val.Latitude,
		longitude:   val.Longitude,
		Status:      val.Status,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}

func (c *car) ToEntity() *rm.Car {
	return &rm.Car{
		Id:          c.ID,
		CarType:     "",
		CarTypeName: c.TypeName,
		CarOfArea:   c.Area,
		ProjectName: c.ProjectName,
		ProjectId:   c.ProjectID,
		Latitude:    c.latitude,
		Longitude:   c.longitude,
		Seat:        c.Seat,
		Distance:    0,
		Status:      c.Status,
	}
}
