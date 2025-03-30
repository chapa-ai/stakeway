package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"regexp"
	"stakeway/internal/model"
	"stakeway/internal/service"
	"stakeway/pkg/response"
)

type Handler struct {
	service service.Service
}

var ethAddressRegex = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

func (h *Handler) CreateValidators(c *fiber.Ctx) error {
	var request model.ValidatorRequest
	if err := c.BodyParser(&request); err != nil {
		h.service.Logger.Errorf("failed to parse validator request: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(response.NewFailedResponse("error parsing request body"))
	}

	if err := h.ValidateInput(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.NewFailedResponse("error validating input data"))
	}
	
	requestID, err := h.service.CreateValidators(c.Context(), &request)
	if err != nil {
		h.service.Logger.Errorf("failed creating validator: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(response.NewFailedResponse("error creating validator"))
	}

	return c.Status(fiber.StatusAccepted).JSON(response.NewCreateValidatorResponse(requestID))
}

func (h *Handler) GetValidatorStatus(c *fiber.Ctx) error {
	requestID := c.Params("request_id")
	resp, err := h.service.GetValidatorStatus(c.Context(), requestID)
	if err != nil {
		h.service.Logger.Errorf("failed to parse validator request: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(response.NewFailedResponse("error processing request"))
	}

	return c.Status(fiber.StatusOK).JSON(response.NewGetValidatorStatusResponse(resp.Status, resp.Keys))
}

func (h *Handler) ValidateInput(request model.ValidatorRequest) error {
	if request.NumValidators <= 0 {
		err := fmt.Errorf("invalid num_validators: %d", request.NumValidators)
		h.service.Logger.Errorf("Validation error: %v", err)
		return err
	}

	if !ethAddressRegex.MatchString(request.FeeRecipient) {
		err := fmt.Errorf("invalid fee_recipient: %s", request.FeeRecipient)
		h.service.Logger.Errorf("Validation error: %v", err)
		return err
	}

	return nil
}
