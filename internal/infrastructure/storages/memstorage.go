package storages

import (
	"github.com/google/uuid"
	"sync"

	matchdata "snake_ai/internal/domain/match/data"
	"snake_ai/internal/domain/user"
)

type MemStorage struct {
	mu      sync.RWMutex
	users   []*user.User
	players []*matchdata.Player
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		users:   make([]*user.User, 0),
		players: make([]*matchdata.Player, 0),
	}
}
func (ms *MemStorage) AddUser(user *user.User) (uuid.UUID, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	user.Id = uuid.New()
	ms.users = append(ms.users, user)

	player := matchdata.NewPlayer()
	player.Id = user.Id
	player.Name = user.Email
	ms.players = append(ms.players, &player)

	return user.Id, nil
}
func (ms *MemStorage) GetUserByEmail(email string) (*user.User, error) {
	for _, u := range ms.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, ErrRecordNotFound
}
func (ms *MemStorage) IsUserExisted(id uuid.UUID) (bool, error) {
	for _, u := range ms.users {
		if u.Id == id {
			return true, nil
		}
	}
	return false, nil
}
func (ms *MemStorage) GetPlayerById(id uuid.UUID) (*matchdata.Player, error) {
	for _, p := range ms.players {
		if p.Id == id {
			return p, nil
		}
	}
	return nil, ErrRecordNotFound
}
func (ms *MemStorage) IncreasePlayerScore(id uuid.UUID) error {
	for _, p := range ms.players {
		if p.Id == id {
			p.Skill++
			return nil
		}
	}
	return ErrRecordNotFound
}
