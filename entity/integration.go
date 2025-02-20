package entity

// Date structure
type Date struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Descriptions structure
type Descriptions struct {
	AppName        string `json:"app_name"`
	AppDescription string `json:"app_description"`
	AppLogo        string `json:"app_logo"`
	AppURL         string `json:"app_url"`
	BackgroundColor string `json:"background_color"`
}

// MonitoringUser structure
type MonitoringUser struct {
	AlwaysOnline bool   `json:"always_online"`
	DisplayName  string `json:"display_name"`
}

// Permissions structure
type Permissions struct {
	MonitoringUser MonitoringUser `json:"monitoring_user"`
}

// Setting structure
type Setting struct {
	Label    string `json:"label"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
	Default  string `json:"default"`
}

type Data struct {
	Date                Date         `json:"date"`
	Descriptions        Descriptions `json:"descriptions"`
	IsActive            bool         `json:"is_active"`
	IntegrationType     string       `json:"integration_type"`
	KeyFeatures         []string     `json:"key_features"`
	Permissions         Permissions  `json:"permissions"`
	Author              string       `json:"author"`
	IntegrationCategory string       `json:"integration_category"`
	Settings            []Setting    `json:"settings"`
	TargetURL           string       `json:"target_url"`
	TickURL             string       `json:"tick_url"`
}

// Integrationson main structure
type Integration struct {
	Data Data `json:"data"`
}

var Integrationjson = Integration{
  Data: Data{
    Date: Date{
      CreatedAt: "2025-02-20",
      UpdatedAt: "2025-02-20",
    },
    Descriptions: Descriptions{
      AppName:        "ForexPI",
      AppDescription: "This is a notification app to give recurring updates on the price of popular Forex symbols",
      AppLogo:        "https://my-portfolio-343207.web.app/MyLogo4.png",
      AppURL:         "https://fun-numbers.onrender.com/telex-webhook",
      BackgroundColor: "#fff",
    },
    IsActive:            false,
    IntegrationType:     "interval",
    KeyFeatures:         []string{"Forex", "Updates"},
    Permissions:         Permissions{MonitoringUser: MonitoringUser{AlwaysOnline: true, DisplayName: "Network Monitor"}},
    Author:              "Samuel Ikoli",
    IntegrationCategory: "Monitoring & Logging",
    Settings: []Setting{
      {
        Label:    "interval",
        Type:     "text",
        Required: true,
        Default:  "* * * * *",
      },
    },
    TargetURL: "https://ping.telex.im/v1/webhooks/01950b90-b1bf-75b7-b9e6-e831fdd18b5f",
    TickURL:   "https://fun-numbers.onrender.com/tick",
  },
}

  type MonitorPayload struct {
    ChannelID string        `json:"channel_id,omitempty"`
    ReturnURL string        `json:"return_url,omitempty"`
    Settings  []interface{} `json:"settings,omitempty"`
  }
  