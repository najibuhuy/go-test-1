package constant

const (
	ErrLoadENV = "error loading .env file: %w"

	ErrConvertStringToInt = "error when convert string to int: %w"
	ErrToMarshalJSON      = "failed json marshal: %w"
)

const (
	ErrConnectToDB             = "failed connect to db: %w"
	ErrInvalidJWTSigningMethod = "invalid jwt signing method"
	ErrInvalidJWTToken         = "invalid jwt token"
)

const (
	ErrInvalidRequest = "invalid request"
	ErrGeneral        = "general error"
)

const (
	ErrHttpServer = "failed create http server"
)
