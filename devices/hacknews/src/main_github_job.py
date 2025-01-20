import sys
import signal

from logger import LOG

from config import Config
from clients.github_client import GitHubClient
from notifiers.githubnotifier import GithubNotifier
from reporters.report_generator import GithubReporter
from llm import LLM
from subscription_manager import SubscriptionManager

def Exit(signum, frame):
    LOG.info("exit")
    sys.exit(0)

def main():
    LOG.info("start, lets go")
    signal.signal(signal.SIGTERM, Exit)

    config = Config()
    github_client = GitHubClient(config.github_token)
    notifier = GithubNotifier(config.email)
    llm = LLM(config)
    reporter = GithubReporter(llm)

    subscription_manager = SubscriptionManager(config.subscriptions_file)
    subscriptions = subscription_manager.list_subscriptions()
    LOG.info(f"subscriptions: {subscriptions}")
    for repo in subscriptions:
        markdown_file_path = github_client.export_progress_by_date_range(repo, 3)
        LOG.info(f"github repo files generated into {markdown_file_path}")
        
        # generate from markdown file
        report, _ = reporter.generate_report(markdown_file_path)
        notifier.notify(report, repo)
    LOG.info(f"[定时任务执行完毕]")


if __name__ == '__main__':
    main()