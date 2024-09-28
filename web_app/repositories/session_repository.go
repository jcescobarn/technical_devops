package repositories

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type SessionRepository struct {
	Client *redis.Client
}

func NewSessionRepository(client *redis.Client) *SessionRepository {
	return &SessionRepository{
		Client: client,
	}
}

func (sr *SessionRepository) Create(SessionID, data string, expiration time.Duration) {

	ctx := context.Background()

	err := sr.Client.Set(ctx, sessionID, data, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (sr *SessionRepository) Get(sessionID string) (string, error) {
	ctx := context.Background()

	val, err := sr.Client.Get(ctx, sessionID).Result()
	if err == redis.Nil {
		return "", err
	} else if err != nil {
		return "", err
	}

	return val, nil
}

func (sr *SessionRepository) Delete() error {
	ctx := context.Background()

	err := sr.Client.Del(ctx, sessionID).Err()
	if err != nil {
		return err
	}
	return nil
}

func (sr *SessionRepository) Update(sessionID, data string, expiration time.Duration) error {
	ctx := context.Background()

	err := sr.Client.Set(ctx, sessionID, data, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}
