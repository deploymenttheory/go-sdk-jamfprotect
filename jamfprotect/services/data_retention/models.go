package dataretention

// DataRetention represents the full response including the org UUID and retention settings.
type DataRetention struct {
	Retention DataRetentionSettings `json:"retention"`
}

// DataRetentionSettings represents organization retention settings.
type DataRetentionSettings struct {
	Database DataRetentionDatabase `json:"database"`
	Cold     DataRetentionCold     `json:"cold"`
	Updated  string                `json:"updated"`
}

// DataRetentionDatabase represents database retention settings.
type DataRetentionDatabase struct {
	Log   RetentionDays `json:"log"`
	Alert RetentionDays `json:"alert"`
}

// DataRetentionCold represents cold storage retention settings.
type DataRetentionCold struct {
	Alert RetentionDays `json:"alert"`
}

// RetentionDays represents a retention period in days.
type RetentionDays struct {
	NumberOfDays int64 `json:"numberOfDays"`
}

// UpdateDataRetentionRequest is the request payload for updating data retention settings.
type UpdateDataRetentionRequest struct {
	DatabaseLogDays   int64
	DatabaseAlertDays int64
	ColdAlertDays     int64
}
