package storages

import (
	"github.com/google/uuid"

	matchdata "snake_ai/internal/domain/match/data"
	"snake_ai/internal/domain/user"
)

var Storage StorageInterface

type StorageInterface interface {
	AddUser(user *user.User) (uuid.UUID, error)
	GetUserByEmail(email string) (*user.User, error)
	IsUserExisted(id uuid.UUID) (bool, error)
	GetPlayerById(id uuid.UUID) (*matchdata.Player, error)
	IncreasePlayerScore(id uuid.UUID) error
}
