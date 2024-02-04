package repositories

import (
	"cleanarchitecturego/cockroach/entities"
)

type CockroachRepository interface {
	InsertCockroachData(in *entities.InsertCockroachDto) error
}
