package repository

import "server/internal/repository/mongodb"

type Repository interface {
}

// assertion
var _ Repository = (*mongodb.MongoClient)(nil)
