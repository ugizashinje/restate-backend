package enum

type EventType string

const (
	CreateWarrant       EventType = "createWarrant"
	UpdateWarrant       EventType = "updateWarrant"
	DeleteWarrant       EventType = "deleteWarrant"
	CreateTransportCost EventType = "createTransportCost"
	UpdateTransportCost EventType = "updateTransportCost"
	DeleteTransportCost EventType = "deleteTransportCost"
	CreateRoute         EventType = "createRoute"
	UpdateRoute         EventType = "updateRoute"
	DeleteRoute         EventType = "deleteRoute"
	CreateRepair        EventType = "createRepair"
	UpdateRepair        EventType = "updateRepair"
	DeleteRepair        EventType = "deleteRepair"
)

type LoginResult string

const (
	UserNotFound  LoginResult = "user_not_found"
	WrongPassword LoginResult = "wrong_password"
	InvalidToken  LoginResult = "invalid_token"
	Success       LoginResult = "success"
)

type UserStatus string

const (
	UserUnverified UserStatus = "unverified"
	UserVerified   UserStatus = "verified"
	UserSick       UserStatus = "sick"
	UserVacation   UserStatus = "vacation"
)

// create, assign driver, loading_time, unloading_time, start, cancel, finish.

type WarrantStatus string

const (
	Template    WarrantStatus = "template"
	Preparation WarrantStatus = "preparation"
	Assigned    WarrantStatus = "assigned"
	Transit     WarrantStatus = "transit"
	Transport   WarrantStatus = "transport"
	Interrupted WarrantStatus = "interrupted"
	Completed   WarrantStatus = "completed"
)

type ConfirmationStatus string

const (
	Unconfirmed ConfirmationStatus = "unconfirmed"
	Confirmed   ConfirmationStatus = "confirmed"
)

type RolesNames string
