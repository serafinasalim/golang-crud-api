package utils

const (
	ErrInvalidRequest     = "Invalid request payload"
	ErrTaskNotFound       = "Task not found"
	ErrTaskNoRecord       = "No tasks found"
	ErrInternalServer     = "Internal server error"
	ErrInvalidTaskId      = "Invalid task id"
	ErrInvalidCredentials = "Invalid username or password"
	ErrUsernameDontExist  = "Username don't exist "
	ErrTokenGeneration    = "Error generating token"
	ErrInvalidToken       = "Invalid or expired token"
	ErrTokenMissing       = "Token is required"
	ErrUnauthorized       = "Unauthorized access"
	ErrUserAlreadyExists  = "User already exists"
	ErrUserNotFound       = "User not found"
	ErrPasswordHashing    = "Error hashing password"
	ErrInvalidEmailFormat = "Invalid email format"
	ErrPasswordMismatch   = "Passwords do not match"
)
