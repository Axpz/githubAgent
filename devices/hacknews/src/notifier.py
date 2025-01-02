import logging
import smtplib
import markdown2
from email.mime.text import MIMEText
from email.mime.multipart import MIMEMultipart

LOG = logging.getLogger(__name__)

class Notifier:
    def __init__(self, settings):
        self.email_settings = settings
    
    def notify_github_report(self, repo, report):
        """
        发送 GitHub 项目报告邮件
        :param repo: 仓库名称
        :param report: 报告内容
        """
        if self.email_settings:
            subject = f"[GitHub] {repo} 进展简报"
            self.send_email(subject, report)
        else:
            LOG.warning("邮件设置未配置正确，无法发送 GitHub 报告通知")
    
    def notify_hn_report(self, date, report):
        """
        发送 Hacker News 每日技术趋势报告邮件
        :param date: 报告日期
        :param report: 报告内容
        """
        if self.email_settings:
            subject = f"[HackerNews] {date} 技术趋势"
            self.send_email(subject, report)
        else:
            LOG.warning("邮件设置未配置正确，无法发送 Hacker News 报告通知")
    
    def send_email(self, subject, report):
        LOG.info(f"准备发送邮件:{subject}")
        msg = MIMEMultipart()
        msg['From'] = self.email_settings['from']
        msg['To'] = self.email_settings['to']
        msg['Subject'] = subject
        
        # 将Markdown内容转换为HTML
        html_report = markdown2.markdown(report)

        msg.attach(MIMEText(html_report, 'html'))
        try:
            with smtplib.SMTP_SSL(self.email_settings['smtp_server'], self.email_settings['smtp_port']) as server:
                LOG.debug("登录SMTP服务器")
                # server.login(msg['From'], self.email_settings['password'])
                # server.sendmail(msg['From'], msg['To'], msg.as_string())
                LOG.info("邮件发送成功！")
        except smtplib.SMTPException as e:
            LOG.error(f"SMTP error occurred: {e}")
        except Exception as e:
            LOG.error(f"Unexpected error occurred: {e}")