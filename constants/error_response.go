package constants

import (
	"fmt"

	"github.com/kalleriakronos24/imbee-backend/dto"
)

func GetErrorResponse(kind string, err error, message string) dto.Response {
	var response dto.Response

	switch kind {
	// ==== url params or body ====
	case "payload-error":
		response = dto.Response{
			Data:    false,
			Kind:    "payload-error",
			Message: "Unknown Payload. Please try again later",
			Error:   err.Error(),
		}
		return response
	case "param-query-error":
		response = dto.Response{
			Data:    false,
			Kind:    "param-query-error",
			Message: "Unknown Param. Please try again later",
			Error:   err.Error(),
		}
		return response
	case "uuid-error":
		response = dto.Response{
			Data:    false,
			Kind:    "uuid-error",
			Message: "Unknown UUID. Please try again later",
			Error:   err.Error(),
		}
		return response

		// ==== data checking ====
	case "data-not-found":
		response = dto.Response{
			Data:    false,
			Kind:    "data-not-found",
			Error:   err.Error(),
			Message: message,
		}
		return response
	case "data-existing":
		response = dto.Response{
			Data:    false,
			Kind:    "data-existing",
			Message: "Data is already exist",
		}
		return response

	case "data-existing-email":
		response = dto.Response{
			Data:    false,
			Kind:    "data-existing",
			Message: fmt.Sprintf("Data is already exist with email %s", message),
		}
		return response
	case "data-existing-username":
		response = dto.Response{
			Data:    false,
			Kind:    "data-existing",
			Message: fmt.Sprintf("Data is already exist with username %s", message),
		}
		return response

		// ==== common crud operations fails ====
	case "insert-failed":
		response = dto.Response{
			Data:    false,
			Message: fmt.Sprintf("Failed inserting new %s", message),
			Kind:    "insert-failed",
			Error:   err.Error(),
		}
		return response
	case "retrieve-failed":
		response = dto.Response{
			Data:    false,
			Kind:    "retrieve-failed",
			Message: fmt.Sprintf("Failed when finding %s", message),
			Error:   err.Error(),
		}
		return response
	case "update-failed":
		response = dto.Response{
			Data:    false,
			Kind:    "update-failed",
			Message: fmt.Sprintf("Failed when updating %s", message),
			Error:   err.Error(),
		}
		return response
	case "delete-failed":
		response = dto.Response{
			Data:    false,
			Kind:    "delete-failed",
			Message: fmt.Sprintf("Failed when deleting %s", message),
			Error:   err.Error(),
		}
		return response

	// ==== general / logical ====
	case "logical":
		response = dto.Response{
			Data:    false,
			Kind:    "logical",
			Message: message,
			Error:   err.Error(),
		}
		return response
	default:
		return response
	}
}
