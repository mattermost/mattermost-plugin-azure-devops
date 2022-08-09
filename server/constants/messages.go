package constants

const (
	// TODO: all these messages are to be verified from Mike at the end
	GenericErrorMessage             = "Something went wrong, please try again later"
	Error                           = "error"
	NotAuthorized                   = "not authorized"
	OrganizationRequired            = "organization is required"
	ProjectRequired                 = "project is required"
	TaskTypeRequired                = "task type is required"
	TaskTitleRequired               = "task title is required"
	ConnectAccount                  = "[Click here to link your Azure DevOps account](%s%s)"
	ConnectAccountFirst             = "You do not have any Azure Devops account connected. Kindly link the account first"
	UserConnected                   = "Your Azure Devops account is succesfully connected!"
	UserAlreadyConnected            = "Your Azure Devops account is already connected"
	UserDisconnected                = "Your Azure Devops account is now disconnected"
	UnableToDisconnectUser          = "Unable to disconnect user"
	UnableToCheckIfAlreadyConnected = "Unable to check if user account is already connected"
	UnableToStoreOauthState         = "Unable to store oAuth state for the userID %s"
	AuthAttemptExpired              = "Authentication attempt expired, please try again"
	InvalidAuthState                = "Invalid oauth state, please try again"
	CreatedTask                     = "Link for newly created task: %s"
)
