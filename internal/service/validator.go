package service

import (
	"context"
	"github.com/google/uuid"
	"log"
	"stakeway/internal/model"
	"stakeway/pkg/errors"
	"time"
)

func (s *Service) CreateValidators(ctx context.Context, request *model.ValidatorRequest) (string, error) {
	requestID, err := s.DB.CreateRequest(ctx, request)
	if err != nil {
		s.Logger.Errorf("failed to create validator request: %v", err)
		return "", errors.Wrapf(err, "failed to create validator request")
	}

	go s.processValidators(ctx, requestID, request.NumValidators, request.FeeRecipient)

	return requestID, nil
}

func (s *Service) GetValidatorStatus(ctx context.Context, requestID string) (*model.ValidatorResponse, error) {
	return s.DB.GetRequest(ctx, requestID)
}

func (s *Service) processValidators(ctx context.Context, requestID string, numValidators int, feeRecipient string) {
	keys := make([]string, numValidators)
	for i := 0; i < numValidators; i++ {
		keys[i] = uuid.New().String()
		time.Sleep(20 * time.Millisecond) // delay imitation
	}

	if err := s.DB.UpdateRequest(ctx, requestID, "successful", keys, feeRecipient); err != nil {
		log.Printf("Failed to update request: %v", err)
	}
}
