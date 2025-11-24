package Api

import (
	"github.com/gofiber/fiber/v3"
)

type DataStruct map[string]any

type Resp struct {
	Status  int         `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (r Resp) Respond(c fiber.Ctx) error {
	status := r.Status
	if status == 0 {
		status = fiber.StatusOK
	}

	return c.Status(status).JSON(fiber.Map{
		"message": r.Message,
		"data":    r.Data,
	})
}

func (r Resp) Error() string {
	return r.Message
}

func SettupApiResponse() fiber.Handler {
	return func(c fiber.Ctx) error {
		err := c.Next()

		if resp, ok := err.(Resp); ok {
			return resp.Respond(c)
		}

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(Resp{
				Status:  fiber.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		return nil
	}
}
