import threading
import shlex
import logging
import sys
import signal
import os
import time
import schedule
from datetime import datetime


from logger import LOG

from argparse import ArgumentError

from config import Config
from scheduler import Scheduler
from clients.github_client import GitHubClient
from clients.hacknews_client import HackerNewsClient
from devices.hacknews.src.notifiers.basenotifier import Notifier
from devices.hacknews.src.report_generator.report_generator import ReportGenerator
from llm import LLM
from subscription_manager import SubscriptionManager
from command_handler import CommandHandler

def Exit(signum, frame):
    LOG.info("exit")
    sys.exit(0)

def github_job(subscription_manager, github_client, report_generator, notifier, days):
    LOG.info("project progress report job")
    subscriptions = subscription_manager.list_subscriptions()  # 获取当前所有订阅
    LOG.info(f"subscriptions: {subscriptions}")
    for repo in subscriptions:
        markdown_file_path = github_client.export_progress_by_date_range(repo, days)
        LOG.info(f"github repo files generated into {markdown_file_path}")
        
        # generate from markdown file
        report, _ = report_generator.generate_github_report(markdown_file_path)
        notifier.notify_github_report(repo, report)
    LOG.info(f"[定时任务执行完毕]")

def hacknews_topic_job(hacker_news_client, report_generator):
    LOG.info("[开始执行定时任务]Hacker News 热点话题跟踪")
    markdown_file_path = hacker_news_client.export_top_stories()
    _, _ = report_generator.generate_hacknews_topic_report(markdown_file_path)
    LOG.info(f"[定时任务执行完毕]")

def hacknews_daily_job(hacker_news_client, report_generator, notifier):
    LOG.info("[开始执行定时任务]Hacker News 今日前沿技术趋势")
    # 获取当前日期，并格式化为 'YYYY-MM-DD' 格式
    date = datetime.now().strftime('%Y-%m-%d')
    # 生成每日汇总报告的目录路径
    directory_path = os.path.join('reports/hacker_news', date)
    # 生成每日汇总报告并保存
    report, _ = report_generator.generate_hacknews_daily_report(directory_path)
    notifier.notify_hn_report(date, report)
    LOG.info(f"[定时任务执行完毕]")
def main():
    LOG.info("start, lets go")
    signal.signal(signal.SIGTERM, Exit)

    config = Config()
    github_client = GitHubClient(config.github_token)
    hacknews_client = HackerNewsClient()
    notifier = Notifier(config.email)
    llm = LLM(config)
    report_generator = ReportGenerator(llm, config.report_types)
    subscription_manager = SubscriptionManager(config.subscriptions_file) # 订阅通知
    
    # github_job(subscription_manager, github_client, report_generator, notifier, 3)
    
    hacknews_topic_job(hacknews_client, report_generator)

    schedule.every(3).days.at("15:00").do(github_job, subscription_manager, github_client, report_generator, notifier, 3)

    schedule.every(4).hours.at(":00").do(hacknews_topic_job, hacknews_client, report_generator)

    try:
        while True:
            schedule.run_pending()
            time.sleep(2)
    except Exception as e:
        LOG.error(f"main proc error: {str(e)}")
        sys.exit(1)

if __name__ == '__main__':
    main()