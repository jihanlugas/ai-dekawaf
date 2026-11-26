package constant

const (
	FAILED_TO_CREATE_UUID = "failed to generate UUID: %w"

	ErrUnexpectedSingingMethodErr = "Unexpected signing method: %v"
	ErrForbidden                  = "Sorry, you don't have the necessary permissions to perform this action. Please check your access rights or contact your administrator for assistance."
	ErrProjectNotAvailable        = "Unable to create security profile service due to the project being either suspended or terminated."
	ErrCheckProjectAvailability   = "Failed to check project availability"
)
