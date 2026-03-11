package dataforwarding

// DataForwarding represents organization data forwarding settings.
type DataForwarding struct {
	UUID    string                 `json:"uuid"`
	Forward DataForwardingSettings `json:"forward"`
}

// DataForwardingSettings holds all forwarding configurations.
type DataForwardingSettings struct {
	S3         ForwardS3         `json:"s3"`
	Sentinel   ForwardSentinel   `json:"sentinel"`
	SentinelV2 ForwardSentinelV2 `json:"sentinelV2"`
}

// ForwardS3 represents S3 forwarding settings.
type ForwardS3 struct {
	Bucket         string `json:"bucket"`
	Enabled        bool   `json:"enabled"`
	Encrypted      bool   `json:"encrypted"`
	Prefix         string `json:"prefix"`
	Role           string `json:"role"`
	CloudFormation string `json:"cloudformation"`
}

// ForwardSentinel represents Sentinel forwarding settings.
type ForwardSentinel struct {
	Enabled    bool   `json:"enabled"`
	CustomerID string `json:"customerId"`
	SharedKey  string `json:"sharedKey"`
	LogType    string `json:"logType"`
	Domain     string `json:"domain"`
}

// SentinelV2DataStream represents a Sentinel v2 data stream.
type SentinelV2DataStream struct {
	Enabled        bool    `json:"enabled"`
	DcrImmutableID *string `json:"dcrImmutableId"`
	StreamName     *string `json:"streamName"`
}

// ForwardSentinelV2 represents Sentinel v2 forwarding settings.
type ForwardSentinelV2 struct {
	Enabled       bool                 `json:"enabled"`
	SecretExists  bool                 `json:"secretExists"`
	AzureTenantID string               `json:"azureTenantId"`
	AzureClientID string               `json:"azureClientId"`
	Endpoint      string               `json:"endpoint"`
	Alerts        SentinelV2DataStream `json:"alerts"`
	ULogs         SentinelV2DataStream `json:"ulogs"`
	Telemetries   SentinelV2DataStream `json:"telemetries"`
	TelemetriesV2 SentinelV2DataStream `json:"telemetriesV2"`
}

// UpdateDataForwardingRequest is the request payload for updating data forwarding settings.
type UpdateDataForwardingRequest struct {
	S3         ForwardS3Input         `json:"s3"`
	Sentinel   ForwardSentinelInput   `json:"sentinel"`
	SentinelV2 ForwardSentinelV2Input `json:"sentinelV2"`
}

// ForwardS3Input captures S3 forwarding updates.
type ForwardS3Input struct {
	Bucket    string `json:"bucket"`
	Enabled   bool   `json:"enabled"`
	Encrypted bool   `json:"encrypted"`
	Prefix    string `json:"prefix"`
	Role      string `json:"role"`
}

// ForwardSentinelInput captures Sentinel forwarding updates.
type ForwardSentinelInput struct {
	Enabled    bool   `json:"enabled"`
	CustomerID string `json:"customerId"`
	SharedKey  string `json:"sharedKey"`
	LogType    string `json:"logType"`
	Domain     string `json:"domain"`
}

// DataStreamInput captures Sentinel v2 data stream updates.
type DataStreamInput struct {
	Enabled        bool    `json:"enabled"`
	DcrImmutableID *string `json:"dcrImmutableId,omitempty"`
	StreamName     *string `json:"streamName,omitempty"`
}

// ForwardSentinelV2Input captures Sentinel v2 forwarding updates.
type ForwardSentinelV2Input struct {
	Enabled           bool            `json:"enabled"`
	AzureTenantID     string          `json:"azureTenantId"`
	AzureClientID     string          `json:"azureClientId"`
	AzureClientSecret *string         `json:"azureClientSecret,omitempty"`
	Endpoint          string          `json:"endpoint"`
	Alerts            DataStreamInput `json:"alerts"`
	ULogs             DataStreamInput `json:"ulogs"`
	Telemetries       DataStreamInput `json:"telemetries"`
	TelemetriesV2     DataStreamInput `json:"telemetriesV2"`
}
