package main

import "time"

type RevealWebHook struct {
	Customer       string    `json:"customer"`
	TenantId       string    `json:"tenant_id"`
	TenantName     string    `json:"tenant_name"`
	TenantOrigin   string    `json:"tenant_origin"`
	AgentUuid      string    `json:"agent_uuid"`
	AgentHostname  string    `json:"agent_hostname"`
	Uuid           string    `json:"uuid"`
	EventType      string    `json:"event_type"`
	StartTimestamp time.Time `json:"start_timestamp"`
	Timestamp      time.Time `json:"timestamp"`
	AlarmType      string    `json:"alarm_type"`
	Description    string    `json:"description"`
	Juid           string    `json:"juid"`
	Username       string    `json:"username"`
	Score          int       `json:"score"`
	TotalScore     int       `json:"total_score"`
	Tags           []string  `json:"tags"`
	Sensors        []struct {
		Customer              string    `json:"customer"`
		TenantId              string    `json:"tenant_id"`
		TenantName            string    `json:"tenant_name"`
		TenantOrigin          string    `json:"tenant_origin,omitempty"`
		AgentUuid             string    `json:"agent_uuid"`
		AgentHostname         string    `json:"agent_hostname"`
		Uuid                  string    `json:"uuid"`
		EventType             string    `json:"event_type"`
		Timestamp             time.Time `json:"timestamp"`
		SensorType            string    `json:"sensor_type"`
		Description           string    `json:"description"`
		Juid                  string    `json:"juid"`
		Username              string    `json:"username"`
		Score                 int       `json:"score"`
		CreatedBy             string    `json:"created_by"`
		Tags                  []string  `json:"tags"`
		AnonymisedDescription string    `json:"anonymised_description"`
		Metadata              struct {
			SourceIp             []interface{} `json:"source_ip"`
			SourcePort           []interface{} `json:"source_port"`
			DestinationIp        []interface{} `json:"destination_ip"`
			DestinationPort      []interface{} `json:"destination_port"`
			Url                  []interface{} `json:"url,omitempty"`
			Host                 []interface{} `json:"host"`
			ApplicationName      []interface{} `json:"application_name"`
			FileName             []interface{} `json:"file_name"`
			FilePath             []interface{} `json:"file_path"`
			TargetFileName       []interface{} `json:"target_file_name"`
			TargetFilePath       []interface{} `json:"target_file_path"`
			RecipientMailAddress []interface{} `json:"recipient_mail_address"`
			SenderMailAddress    []interface{} `json:"sender_mail_address"`
			WifiSsid             []interface{} `json:"wifi_ssid"`
			WifiBssid            []interface{} `json:"wifi_bssid"`
			UsbVid               []interface{} `json:"usb_vid"`
			UsbPid               []interface{} `json:"usb_pid"`
			UsbSerial            []interface{} `json:"usb_serial"`
			ContentPatternName   []interface{} `json:"content_pattern_name"`
			AccountName          []interface{} `json:"account_name"`
			CertificateName      []interface{} `json:"certificate_name"`
			UrL                  []interface{} `json:"ur l,omitempty"`
		} `json:"metadata"`
		ProcessInfo []struct {
			Uuid          string `json:"uuid"`
			BinaryName    string `json:"binary_name"`
			BinaryPath    string `json:"binary_path"`
			Username      string `json:"username"`
			AppIdentifiEr string `json:"app_identifi er,omitempty"`
			Signed        bool   `json:"signed"`
			AppIdentifier string `json:"app_identifier,omitempty"`
		} `json:"process_info"`
		TenantORigin string `json:"tenant_o rigin,omitempty"`
	} `json:"sensors"`
}
