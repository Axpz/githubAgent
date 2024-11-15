import json
import os

class Config:
    def __init__(self):
        self.load_config()
    
    def load_config(self):
        with open('config.json', 'r') as f:
            config = json.load(f)
            self.github_token = config.get('github_token')
            if not self.github_token:
                self.github_token = os.getenv('OPENAI_API_KEY')
                if not self.github_token:
                    raise ValueError("GitHub token not found in config file or environment variables")
                
            self.notification_settings = config.get('notification_settings')
            self.subscriptions_file = config.get('subscriptions_file')
            self.update_interval = config.get('update_interval', 24 * 60 * 60)  # Default to 24 hours