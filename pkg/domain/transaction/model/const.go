package model

const (
	ACTIVE  = "active"
	CANCEL  = "cancel"
	SUCCESS = "success"
)

var ValidateStatus = map[string]map[string]bool{
	ACTIVE: {
		CANCEL:  true,
		SUCCESS: true,
	},
	CANCEL: {
		ACTIVE:  false,
		SUCCESS: false,
	},
	SUCCESS: {
		ACTIVE:  false,
		SUCCESS: false,
	},
}
