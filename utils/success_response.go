// Package utils provides utility functions to support various operations in a Fiber application.

package utils

import "github.com/gofiber/fiber/v2"

// SuccessResponse sends a standardized JSON success response to the client using the Fiber context.
// It sets the response status to OK (200) and returns a JSON object containing a user-friendly
// message and additional data.
//
// Parameters:
//   - c: The *fiber.Ctx object representing the Fiber context for the current HTTP request.
//   - msg: A user-friendly message indicating the success of the operation.
//   - data: The data payload to be included in the response. This can be any type that is valid for JSON encoding.
//
// Returns:
//   - An error if the JSON response cannot be sent; otherwise, it returns nil.
//
// Example usage:
//
//	func someFiberHandler(c *fiber.Ctx) error {
//	    // Some processing...
//	    result := someOperation()
//	    return utils.SuccessResponse(c, "Operation successful", result)
//	}
func SuccessResponse(c *fiber.Ctx, msg string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": msg,
		"data":    data,
	})
}
