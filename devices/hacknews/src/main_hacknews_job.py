import sys
import signal

from logger import LOG

from config import Config
from clients.hacknews_client import HackerNewsClient
from notifiers.hacknewsnotifier import HacknewsNotifier
from reporters.report_generator import HacknewsReporter
from llm import LLM
from subscription_manager import SubscriptionManager
from datetime import datetime

def Exit(signum, frame):
    LOG.info("exit")
    sys.exit(0)

def main():
    LOG.info("start, lets go")
    signal.signal(signal.SIGTERM, Exit)

    config = Config()
    hacknews_client = HackerNewsClient()
    notifier = HacknewsNotifier(config.email)
    llm = LLM(config)
    reporter = HacknewsReporter(llm)

    markdown_file_path = hacknews_client.export_top_stories()
    report, _ = reporter.generate_report(markdown_file_path)
    if 0 <= datetime.now().hour < 24:
        notifier.notify(report, datetime.now().strftime('%Y-%m-%d'))
        LOG.info(f"[send email]")

    LOG.info(f"[定时任务执行完毕]")


if __name__ == '__main__':
    main()
