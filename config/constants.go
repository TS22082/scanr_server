package config

const (
	Production      = "production"
	Development     = "development"
	DoNotReplyEmail = "scanrdonotreply@gmail.com"
)

const (
	DbName      = "todo_db_golang"
	Users       = "users"
	Tokens      = "register_tokens"
	EmailTokens = "email_tokens"
)

const (
	JWTNotSetMessage        = "jwt secret is not set in .env file"
	ErrorParsingJWT         = "Error parsing JWT - likely expried or invalid"
	InvalidJWT              = "Invalid JWT"
	RunningDevMode          = "Running in development mode."
	EnvError                = "Error loading .env file, continuing without it"
	SuccessfulVerification  = "User successfully registered"
	ObjectAssertionError    = "Error asserting ObjectID"
	TemplateParseError      = "Error parsing template"
	TemplateExecError       = "Error executing template"
	VerifyEmailSubject      = "Verify your account"
	EmailSendError          = "Error sending email"
	TokenDeleteFailed       = "Error deleting token"
	TokenCreatedFailed      = "Error creating token"
	UserNotFound            = "User does not exist, or token is invalid"
	UserUpdateFailed        = "Error updating user"
	UserCreateSuccess       = "User successfully created"
	UserCreateFailed        = "Error creating user"
	UserExistsError         = "User already exists"
	ErrorCheckingUserExists = "There was an error checking to see if the user exists"
	ErrorGettingUser        = "Error getting user"
	ErrorParsingBody        = "Error parsing request body"
	PasswordHashingError    = "Error hashing password"
	UserDoesNotExist        = "User does not exist"
	UserNotVerified         = "User is not verified"
	InvalidCredentials      = "Invalid credentials"
	ErrorGeneratingJWT      = "Error generating JWT"
	Success                 = "Success"
)
