// package utils provides utility functions for the Fiber application.

package utils

import (
	"github.com/gofiber/fiber/v2"
)

// ErrorResponse sends a standardized JSON error response to the client using the Fiber context.
// It sets the response status to Internal Server Error (500) and returns a JSON object containing
// a user-friendly message and a detailed error message.
//
// The function takes three parameters:
//   - c: The *fiber.Ctx object representing the Fiber context for the current HTTP request.
//   - msg: A user-friendly message describing the error, intended to be displayed to the end user.
//   - err: The error object. If err is not nil, its error message is included in the response;
//     otherwise, a default "No additional error information" message is used.
//
// The function returns an error object if there was an error sending the response; otherwise, it returns nil.
//
// Example usage:
//
//	if err != nil {
//	    return utils.ErrorResponse(c, "An error occurred", err)
//	}
func ErrorResponse(c *fiber.Ctx, msg string, err error) error {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	} else {
		errMsg = "No additional error information"
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": msg,
		"error":   errMsg,
	})
}
