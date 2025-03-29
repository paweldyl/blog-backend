package db

import (
	"github.com/jackc/pgx/v5"
)

const UniqueViolation = "23505"

var ErrRecordNotFound = pgx.ErrNoRows
