import requests

class UpdateFetcher:
    GITHUB_API_URL = "https://api.github.com"

    def fetch_updates(self, repositories):
        updates = {}
        for repo in repositories:
            response = requests.get(f"{self.GITHUB_API_URL}/repos/{repo}/events")
            if response.status_code == 200:
                updates[repo] = response.json()
        return updates
