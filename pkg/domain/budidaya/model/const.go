package model

const (
	MANDIRI = "mandiri"
	TEAM    = "team"

	SUBMISION = "submission"
	REVIEWED  = "reviewed"
	ACTIVED   = "actived"
	DISABLED  = "disabled"
)

var MapStatus = map[string]map[string]bool{
	SUBMISION: {
		SUBMISION: false,
		REVIEWED:  true,
		ACTIVED:   false,
		DISABLED:  true,
	},
	REVIEWED: {
		SUBMISION: false,
		REVIEWED:  false,
		ACTIVED:   true,
		DISABLED:  true,
	},
	ACTIVED: {
		SUBMISION: false,
		REVIEWED:  false,
		ACTIVED:   false,
		DISABLED:  true,
	},
	DISABLED: {
		SUBMISION: false,
		REVIEWED:  false,
		ACTIVED:   false,
		DISABLED:  true,
	},
}
