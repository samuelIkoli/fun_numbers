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

var Integrationson = map[string]interface{}{
    "data": map[string]interface{}{
      "date": map[string]interface{}{
        "created_at": "2025-02-20",
        "updated_at": "2025-02-20",
      },
      "descriptions": map[string]interface{}{
        "app_name": "ForexAPISam",
        "app_description": "This is a notification app to give recurring updates on the price of popular Forex symbols",
        "app_logo": "https://my-portfolio-343207.web.app/MyLogo4.png",
        "app_url": "https://fun-numbers.onrender.com/telex-webhook",
        "background_color": "#fff",
      },
      "is_active": false,
      "integration_type": "interval",
      "key_features": []string{
        "Forex",
        "Updates",
      },
      "permissions": map[string]interface{}{
                    "monitoring_user": map[string]interface{}{
                        "always_online": true,
                        "display_name": "Network Monitor",
                    },
      },
      "author": "Samuel Ikoli",
      "integration_category": "Monitoring & Logging",
      "settings": []map[string]interface{}{
        {
          "label": "interval",
          "type": "text",
          "required": true,
          "default": "* * * * *",
        },
        },
      },
      "target_url": "",
      "tick_url": "https://fun-numbers.onrender.com/tick",
    }

  type MonitorPayload struct {
    ChannelID string        `json:"channel_id,omitempty"`
    ReturnURL string        `json:"return_url,omitempty"`
    Settings  []interface{} `json:"settings,omitempty"`
  }
  