package entity

var Integrationson = map[string]interface{}{
    "data": map[string]interface{}{
      "date": map[string]interface{}{
        "created_at": "2025-02-20",
        "updated_at": "2025-02-20",
      },
      "descriptions": map[string]interface{}{
        "app_name": "ForexPI",
        "app_description": "This is a notification app to give recurring updates on the price of popular Forex symbols",
        "app_logo": "https://my-portfolio-343207.web.app/MyLogo4.png",
        "app_url": "https://fun-numbers.onrender.com/telex-webhook",
        "background_color": "#fff",
      },
      "is_active": false,
      "integration_type": "interval",
      "integration_category": "Monitoring & Logging",
      "key_features": []string{
        "Forex",
        "Updates",
      },
      "author": "Samuel Ikoli",
      "permissions": map[string]interface{}{
                    "monitoring_user": map[string]interface{}{
                        "always_online": true,
                        "display_name": "Network Monitor",
                    },
      },
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
  