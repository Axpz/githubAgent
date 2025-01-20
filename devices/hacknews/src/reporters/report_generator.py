import os
import logging

LOG = logging.getLogger(__name__)
logging.basicConfig(level=logging.INFO)  # Ensure logging is configured

class ReportGenerator:
    def __init__(self, llm, reporter):
        # Initialize with an LLM and preload prompts for report generation
        self.llm = llm
        self.report_types = reporter
        self.prompts = {}
        self._preload_prompts()

    def _preload_prompts(self):
        for report_type in self.report_types:
            prompt_file = f"prompts/{report_type}_prompt.md"
            LOG.info(f"Getting prompt file template {prompt_file}")
            self._load_prompt_file(report_type, prompt_file)

    def _load_prompt_file(self, report_type, prompt_file):
        if not os.path.exists(prompt_file):
            LOG.error(f"Prompt file does not exist: {prompt_file}")
            raise FileNotFoundError(f"Prompt file not found: {prompt_file}")
        with open(prompt_file, "r", encoding='utf-8') as file:
            self.prompts[report_type] = file.read()

    def _generate_report(self, markdown_file_path, report_type, output_suffix):
        """
        Generate reports from the markdown file.
        """
        markdown_content = self._read_markdown_file(markdown_file_path)
        system_prompt = self.prompts.get(report_type)

        if not system_prompt:
            LOG.error(f"No prompt found for report type: {report_type}")
            return None, None

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
            LOG.error("Report is None")
            return
        
        os.makedirs(os.path.dirname(report_file_path), exist_ok=True)
        with open(report_file_path, 'w+', encoding='utf-8') as report_file:
            report_file.write(report)


class GithubReporter(ReportGenerator):
    def __init__(self, llm):
        # Initialize with the llm and specify the report type for GitHub
        reporter = ["github"]
        super().__init__(llm, reporter)

    def generate_report(self, markdown_file_path="reports"):
        # Generate GitHub specific report
        return self._generate_report(markdown_file_path, "github", "_report.md")


class HacknewsReporter(ReportGenerator):
    def __init__(self, llm):
        # Initialize with the llm and specify the report type for Hacknews
        reporter = ["hackernews"]
        super().__init__(llm, reporter)

    def generate_report(self, markdown_file_path="reports"):
        # Generate Hacknews specific report
        return self._generate_report(markdown_file_path, "hackernews", "_report.md")
