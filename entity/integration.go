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
      "settings": []map[string]interface{}{
        {
          "label": "interval",
          "type": "dropdown",
          "required": true,
          "default": "daily",
          "options": []string{
            "* * * * *",
            "*/5 * * * *",
            "*/60 * * * *",
            "*/550 * * * *",
            "0 * * * *",
          },
        },
      },
      "target_url": "https://ping.telex.im/v1/webhooks/01950b90-b1bf-75b7-b9e6-e831fdd18b5f",
      "tick_url": "https://fun-numbers.onrender.com/tick",
    },
  }