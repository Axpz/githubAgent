import logging
import smtplib
import markdown2
from email.mime.text import MIMEText
from email.mime.multipart import MIMEMultipart

LOG = logging.getLogger(__name__)

class Notifier:
    def __init__(self, settings):
        self.email_settings = settings
        
        # 检查邮件设置中是否包含必需的字段
        required_keys = ['from', 'to', 'smtp_server', 'smtp_port', 'password']
        for key in required_keys:
            if key not in self.email_settings:
                LOG.error(f"邮件设置缺少必要的字段: {key}")
                raise ValueError(f"邮件设置缺少必要的字段: {key}")
    
    def send_email(self, subject, report):
        """
        发送电子邮件
        :param subject: 邮件主题
        :param report: 邮件内容(Markdown格式)
        """
        if not report.strip():
            LOG.warning("报告内容为空，邮件未发送")
            return

        if not self.email_settings:
            LOG.warning("邮件设置未配置正确，无法发送报告通知")
            return

        LOG.info(f"准备发送邮件:{subject}")
        msg = MIMEMultipart()
        msg['From'] = self.email_settings['from']
        # msg['To'] = self.email_settings['to']
        msg['Subject'] = subject

        bcc_emails = self.email_settings['to'].split(';')
        
        # 将Markdown内容转换为HTML
        html_report = markdown2.markdown(report)
        msg.attach(MIMEText(html_report, 'html'))

        try:
            with smtplib.SMTP_SSL(self.email_settings['smtp_server'], self.email_settings['smtp_port']) as server:
                LOG.debug("登录SMTP服务器")
                # 登录SMTP服务器
                server.login(msg['From'], self.email_settings['password'])
                # 发送邮件
                server.sendmail(msg['From'], bcc_emails, msg.as_string())
                LOG.info("邮件发送成功！")
        except smtplib.SMTPException as e:
            LOG.error(f"SMTP error occurred: {e}")
        except Exception as e:
            LOG.error(f"Unexpected error occurred: {e}")




