from .basenotifier import Notifier

class HacknewsNotifier(Notifier):
    def __init__(self, settings):
        super().__init__(settings)

    def notify(self, report, date):
        """
        发送 Hacker News 每日技术趋势报告邮件
        :param report: 报告内容
        """
        subject = f"[HackerNews] {date} 技术趋势"
        self.send_email(subject, report)