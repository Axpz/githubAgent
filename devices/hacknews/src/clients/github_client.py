import os
import logging
import requests
from datetime import datetime, date, timedelta

LOG = logging.getLogger(__name__)

REPORTS_DIR = "reports"

class GitHubClient:
    def __init__(self, token, timeout=20):
        self.token = token
        self.timeout = timeout
        self.headers = {'Authorization': f'token {self.token}'}

    def fetch_updates(self, repo, since=None, until=None):
        return {
            'commits': self.fetch_commits(repo, since, until),
            'issues': self.fetch_issues(repo, since, until),
            'pull_requests': self.fetch_pull_requests(repo, since, until)
        }

    def fetch_commits(self, repo, since=None, until=None):
        LOG.debug(f"Fetching commits for {repo}")
        url = f'https://api.github.com/repos/{repo}/commits'
        params = self._build_date_params(since, until)
        return self._make_request(url, params)

    def fetch_issues(self, repo, since=None, until=None):
        LOG.debug(f"Fetching issues for {repo}")
        url = f'https://api.github.com/repos/{repo}/issues'
        params = {'state': 'closed', **self._build_date_params(since, until)}
        return self._make_request(url, params)

    def fetch_pull_requests(self, repo, since=None, until=None):
        LOG.debug(f"Fetching pull requests for {repo}")
        url = f'https://api.github.com/repos/{repo}/pulls'
        params = {'state': 'closed', **self._build_date_params(since, until)}
        return self._make_request(url, params)

    def export_daily_progress(self, repo):
        LOG.debug(f"Exporting daily progress for {repo}")
        today = datetime.now().date().isoformat()
        updates = self.fetch_updates(repo, since=today)
        if not updates or not any(updates.values()):
            LOG.warning(f"No updates found for {repo} on {today}. Skipping export.")
            return None
        return self._export_to_file(repo, today, updates)

    def export_weekly_progress(self, repo):
        LOG.debug(f"Preparing to export weekly progress for {repo}")
        today = date.today()
        since = today - timedelta(days=7)
        updates = self.fetch_updates(repo, since)
        if not updates or not any(updates.values()):
            LOG.warning(f"No updates found for {repo} in the last 7 days. Skipping export.")
            return None
        return self._export_to_file(repo, f"{since}_to_{today}", updates)

    def export_progress_by_date_range(self, repo, days):
        today = date.today()
        since = today - timedelta(days=days)
        updates = self.fetch_updates(repo, since=since.isoformat(), until=today.isoformat())
        if not updates or not any(updates.values()):
            LOG.warning(f"No updates found for {repo} from {since} to {today}. Skipping export.")
            return None
        date_range = f"{since.isoformat()}_to_{today.isoformat()}"
        return self._export_to_file(repo, date_range, updates)

    def _build_date_params(self, since, until):
        params = {}
        if since:
            params['since'] = since
        if until:
            params['until'] = until
        return params

    def _make_request(self, url, params):
        LOG.debug(f"Making request to {url} with params {params}")
        try:
            response = requests.get(url, headers=self.headers, params=params, timeout=self.timeout)
            response.raise_for_status()
            return response.json()
        except requests.HTTPError as http_err:
            LOG.error(f"HTTP error for {url}: {http_err} | Response: {response.text if response else 'No response'}")
        except requests.RequestException as req_err:
            LOG.error(f"Request failed for {url}: {req_err}")
        except Exception as e:
            LOG.error(f"Unexpected error: {e}")
        return []

    def _export_to_file(self, repo, filename_suffix, updates):
        repo_dir = os.path.join(REPORTS_DIR, repo.replace("/", "_"))
        os.makedirs(repo_dir, exist_ok=True)
        file_path = os.path.join(repo_dir, f'{filename_suffix}.md')

        with open(file_path, 'w') as file:
            file.write(f"# Progress for {repo} ({filename_suffix})\n\n")
            
            # Commits
            file.write("\n## Commits\n")
            for commit in updates['commits']:
                file.write(f"- {commit['commit']['message']} ({commit['sha'][:7]})\n")
            
            # Issues
            file.write("\n## Issues Closed\n")
            for issue in updates['issues']:
                file.write(f"- {issue['title']} #{issue['number']}\n")
            
            # Pull Requests
            file.write("\n## Pull Requests Closed\n")
            for pr in updates['pull_requests']:
                file.write(f"- {pr['title']} #{pr['number']}\n")

        LOG.info(f"Progress file created: {file_path}")
        return file_path

if __name__ == "__main__":
    logging.basicConfig(
        level=logging.DEBUG,
        format='%(asctime)s - %(name)s - %(levelname)s - %(message)s',
        handlers=[logging.FileHandler(f"progress_exporter_{date.today().isoformat()}.log"), logging.StreamHandler()]
    )
    token = os.getenv('GITHUB_API_TOKEN')
    exporter = GitHubClient(token)
    repo = '0xPolygon/cdk'  # Replace with the target repository
    exporter.export_daily_progress(repo)
    exporter.export_weekly_progress(repo)

    repo = 'axpz/githubAgent'  # Replace with the target repository
    exporter.export_weekly_progress(repo)

    repo = 'dingodb/dingo'
    exporter.export_weekly_progress(repo)
