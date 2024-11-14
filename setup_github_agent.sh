#!/bin/bash

# 定义项目根目录
PROJECT_DIR="github-agent"

# 创建项目根目录
mkdir -p $PROJECT_DIR

# 创建文件和目录结构
mkdir -p $PROJECT_DIR/{core,utils,tests}

# 在项目根目录下创建主文件和配置文件
touch $PROJECT_DIR/{main.py,config.py,requirements.txt}

# 在 core 目录下创建各个模块文件
touch $PROJECT_DIR/core/{subscription.py,fetcher.py,notifier.py,report_generator.py}

# 在 utils 目录下创建工具文件
touch $PROJECT_DIR/utils/{logger.py,helpers.py}

# 在 tests 目录下创建测试文件
touch $PROJECT_DIR/tests/{test_subscription.py,test_fetcher.py,test_notifier.py,test_report.py}

# 创建每个文件的初始代码
echo 'from core.subscription import SubscriptionManager
from core.fetcher import UpdateFetcher
from core.notifier import Notifier
from core.report_generator import ReportGenerator

def main():
    subscription_manager = SubscriptionManager()
    fetcher = UpdateFetcher()
    notifier = Notifier()
    report_generator = ReportGenerator()
    
    repositories = subscription_manager.get_subscribed_repositories()
    updates = fetcher.fetch_updates(repositories)
    
    if updates:
        notifier.notify(updates)
    
    report_generator.generate_report(updates)

if __name__ == "__main__":
    main()' > $PROJECT_DIR/main.py

echo 'class SubscriptionManager:
    def __init__(self):
        self.subscriptions = []

    def add_subscription(self, repo_name):
        if repo_name not in self.subscriptions:
            self.subscriptions.append(repo_name)

    def remove_subscription(self, repo_name):
        if repo_name in self.subscriptions:
            self.subscriptions.remove(repo_name)

    def get_subscribed_repositories(self):
        return self.subscriptions' > $PROJECT_DIR/core/subscription.py

echo 'import requests

class UpdateFetcher:
    GITHUB_API_URL = "https://api.github.com"

    def fetch_updates(self, repositories):
        updates = {}
        for repo in repositories:
            response = requests.get(f"{self.GITHUB_API_URL}/repos/{repo}/events")
            if response.status_code == 200:
                updates[repo] = response.json()
        return updates' > $PROJECT_DIR/core/fetcher.py

echo 'class Notifier:
    def __init__(self):
        pass

    def notify(self, updates):
        for repo, events in updates.items():
            print(f"通知: 仓库 {repo} 有新的更新!")' > $PROJECT_DIR/core/notifier.py

echo 'import json
from datetime import datetime

class ReportGenerator:
    def generate_report(self, updates):
        report_data = {
            "date": datetime.now().strftime("%Y-%m-%d %H:%M:%S"),
            "updates": updates,
        }
        with open("update_report.json", "w") as file:
            json.dump(report_data, file)
        print("报告已生成：update_report.json")' > $PROJECT_DIR/core/report_generator.py

# 打印成功创建的文件结构
echo "Project structure created under '$PROJECT_DIR':"
tree $PROJECT_DIR

