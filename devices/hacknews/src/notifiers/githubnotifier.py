from .basenotifier import Notifier

class GithubNotifier(Notifier):
    def __init__(self, settings):
        super().__init__(settings)

    def notify(self, report, repo):
        """
        发送 GitHub 项目报告邮件
        :param report: 报告内容
        """
        subject = f"[GitHub] {repo} 进展简报"
        self.send_email(subject, report)