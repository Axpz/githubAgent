import logging
import smtplib
import ssl
import markdown2
from email.mime.text import MIMEText
from email.mime.multipart import MIMEMultipart
from email.header import Header

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
        msg['Subject'] = Header(subject, 'utf-8')

        receivers = self.email_settings['to']
        
        # 将Markdown内容转换为HTML
        html_report = markdown2.markdown(report)
        msg.attach(MIMEText(html_report, 'html', 'utf-8'))

        context = ssl.create_default_context()
        try:
            with smtplib.SMTP_SSL(self.email_settings['smtp_server'], self.email_settings['smtp_port'], context=context) as server:
                LOG.debug("登录SMTP服务器")
                # 登录SMTP服务器
                # server.set_debuglevel(2)
                # server.ehlo()
                # server.starttls()

                server.login(msg['From'], self.email_settings['password'])
                # 发送邮件
                
                # server.sendmail(msg['From'], receivers, msg.as_string())
                LOG.info("send email to %s", receivers)
                for receiver in receivers:
                    try:
                        msg['To'] = receiver
                        server.sendmail(msg['From'], receiver, msg.as_string())
                        LOG.info("send email to %s", receiver)
                    except Exception as e:
                        LOG.error(f"SMTP error occurred: {e}")
        except Exception as e:
            LOG.error(f"Unexpected error occurred: {e}")




