package respository

import "auth/internal/respository/postgres"

type Repository interface {
}

var (
	_ Repository = (*postgres.PostgreSQL)(nil)
)
