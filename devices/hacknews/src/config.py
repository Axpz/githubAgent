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
                self.github_token = os.getenv('GITHUB_API_TOKEN')
                if not self.github_token:
                    raise ValueError("GitHub token not found in config file or environment variables")
                
            self.notification_settings = config.get('notification_settings')
            self.email = self.notification_settings.get("email")
            self.email['password'] = os.getenv(self.email.get('password_env'))

            self.subscriptions_file = config.get('subscriptions_file')
            self.update_interval = config.get('update_interval', 24 * 60 * 60)  # Default to 24 hours

            self.llm_model = config.get('llm_model') or "gpt-3.5-turbo"
            self.llm_api_key = os.getenv(config.get('llm_token_env_variable'))
            self.llm_base_url = config.get('llm_base_url') or "https://api.openai.com"
            self.llm_dry_run = config.get('llm_dry_run', True)

            self.report_types = config.get('report_types', ["github", "hacker_news"])  # 默认报告类型