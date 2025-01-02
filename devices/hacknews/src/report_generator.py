import os
import logging

from logger import LOG
from logger import log

class ReportGenerator:
    def __init__(self, llm, report_types):
        # Initialize with an LLM and preload prompts for report generation
        self.llm = llm  
        self.report_types = report_types
        self.prompts = {}
        self._preload_prompts()

    def _preload_prompts(self):
        for report_type in self.report_types:
            prompt_file = f"prompts/{report_type}_prompt.txt"
            log.info(f"get prompt file template {prompt_file}")

            self._load_prompt_file(report_type, prompt_file)

    def _load_prompt_file(self, report_type, prompt_file):
        if not os.path.exists(prompt_file):
            LOG.error(f"Prompt file does not exist: {prompt_file}")
            raise FileNotFoundError(f"Prompt file not found: {prompt_file}")
        with open(prompt_file, "r", encoding='utf-8') as file:
            self.prompts[report_type] = file.read()

    def _generate_report(self, markdown_file_path, report_type, output_suffix):
        """
        generate reports from the markdown file
        """
        markdown_content = self._read_markdown_file(markdown_file_path)

        system_prompt = self.prompts.get(report_type)

        report = self.llm.generate_report(system_prompt, markdown_content)

        # Generate the report file path by appending the suffix
        report_file_path = os.path.splitext(markdown_file_path)[0] + output_suffix
        self._save_report(report_file_path, report)

        LOG.info(f"{report_type} report saved to {report_file_path}")
        return report, report_file_path

    def _read_markdown_file(self, markdown_file_path):
        with open(markdown_file_path, 'r', encoding='utf-8') as file:
            return file.read()

    def _save_report(self, report_file_path, report):
        if not report:
            LOG.error("report is none")
            return
        
        os.makedirs(os.path.dirname(report_file_path), exist_ok=True)
        with open(report_file_path, 'w+', encoding='utf-8') as report_file:
            report_file.write(report)

    def generate_github_report(self, markdown_file_path):
        return self._generate_report(markdown_file_path, "github", "_report.md")

    def generate_hacknews_topic_report(self, markdown_file_path):
        return self._generate_report(markdown_file_path, "hacker_news_hours_topic", "_topic.md")

    def generate_hacknews_daily_report(self, directory_path):
        """
        Generate a Hacker News daily summary report and save it to hacker_news/tech_trends/ directory.
        The input is a directory path that contains all the *_topic.md files generated.
        """
        markdown_content = self._aggregate_topic_reports(directory_path)
        system_prompt = self.prompts.get("hacker_news")

        base_name = os.path.basename(directory_path.rstrip('/'))
        report_file_path = os.path.join("hacker_news/tech_trends/", f"{base_name}_trends.md")

        report = self.llm.generate_report(system_prompt, markdown_content)

        self._save_report(report_file_path, report)

        LOG.info(f"Hacker News daily summary report saved to {report_file_path}")
        return report, report_file_path

    def _aggregate_topic_reports(self, directory_path):
        """
        Aggregate all markdown files ending with '_topic.md' in the directory to create the input for the daily summary report.
        """
        markdown_content = ""
        for filename in os.listdir(directory_path):
            if filename.endswith("_topic.md"):
                with open(os.path.join(directory_path, filename), 'r', encoding='utf-8') as file:
                    markdown_content += file.read() + "\n"
        return markdown_content


if __name__ == '__main__':
    from config import Config
    from llm import LLM

    config = Config()
    llm = LLM(config)
    report_generator = ReportGenerator(llm, config.report_types)

    # hn_hours_file = "./hacker_news/2024-09-01/14.md"
    hn_daily_dir = "/tmp/reports/2099-09-09/"

    # Generate the Hacker News hourly topic report
    # report, report_file_path = report_generator.generate_hn_topic_report(hn_hours_file)
    report, report_file_path = report_generator.generate_hacknews_daily_report(hn_daily_dir)
    LOG.debug(report)
