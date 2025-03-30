package pg

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"log"
	"stakeway/internal/model"
	"stakeway/pkg/errors"
)

type Repository interface {
	CreateRequest(ctx context.Context, request *model.ValidatorRequest) (string, error)
	GetRequest(ctx context.Context, requestID string) (*model.ValidatorResponse, error)
	UpdateRequest(ctx context.Context, requestID string, status string, keys []string, fee_recipient string) error
}

type DB struct {
	DB *sqlx.DB
}

func (r *DB) CreateRequest(ctx context.Context, request *model.ValidatorRequest) (string, error) {
	query := `INSERT INTO validator_requests (id, num_validators, fee_recipient, status) 
			  VALUES (:id, :num_validators, :fee_recipient, :status)`

	requestID := uuid.New().String()
	_, err := r.DB.NamedExec(query, map[string]interface{}{
		"id":             requestID,
		"num_validators": request.NumValidators,
		"fee_recipient":  request.FeeRecipient,
		"status":         "started",
	})
	if err != nil {
		return "", errors.Wrapf(err, "failed to create request")
	}
	return requestID, nil
}

func (r *DB) GetRequest(ctx context.Context, requestID string) (*model.ValidatorResponse, error) {
	var response model.ValidatorResponse

	query := `SELECT id AS request_id, status FROM validator_requests WHERE id = ?`
	err := r.DB.GetContext(ctx, &response, query, requestID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get request")
	}

	keysQuery := `SELECT key FROM validator_keys WHERE request_id = ?`

	var keys []string
	err = r.DB.SelectContext(ctx, &keys, keysQuery, requestID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get keys for request")
	}

	response.Keys = keys
	return &response, nil
}

func (r *DB) UpdateRequest(ctx context.Context, requestID string, status string, keys []string, feeRecipient string) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return errors.Wrapf(err, "failed to start transaction")
	}

	updateStatusQuery := `UPDATE validator_requests SET status = ? WHERE id = ?`
	if _, err = tx.ExecContext(ctx, updateStatusQuery, status, requestID); err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "failed to update status")
	}

	insertKeysQuery := `INSERT INTO validator_keys (request_id, key, fee_recipient) VALUES (?, ?, ?)`
	for _, key := range keys {
		if _, err := tx.ExecContext(ctx, insertKeysQuery, requestID, key, feeRecipient); err != nil {
			tx.Rollback()
			return errors.Wrapf(err, "failed to insert keys")
		}
	}

	if err = tx.Commit(); err != nil {
		return errors.Wrapf(err, "failed to commit transaction")
	}

	return nil
}

func NewRepository(dataSourceName string) (*DB, error) {
	db, err := sqlx.Connect("sqlite3", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to database")
	return &DB{db}, nil
}
