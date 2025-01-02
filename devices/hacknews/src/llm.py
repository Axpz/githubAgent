import logging
from openai import OpenAI

LOG = logging.getLogger(__name__)


class LLM:
    def __init__(self, config=None):
        self.dry_run = config.llm_dry_run

        self.config = config

        self.model = config.llm_model
        self.api_key = config.llm_api_key
        self.base_url = config.llm_base_url
        self.client = OpenAI(api_key=f"{self.api_key}", base_url=self.base_url)
    
    def generate_report(self, system_prompt, user_content):
        messages = [
            {"role": "system", "content": system_prompt},
            {"role": "user", "content": user_content},
        ]
        return self._generate_report_openai(messages)
    
    def _generate_report_openai(self, messages):
        LOG.info(f"using OpenAI {self.model} template to generate reports")
        try:
            response = self.client.chat.completions.create(
                model=self.model,
                messages=messages,
                temperature=0.2
            )
            LOG.debug("chatGPT response: {}", response)
            return response.choices[0].message.content
        except Exception as e:
            LOG.error(f"generate report openai error: {e}")


if __name__ == '__main__':
    from config import Config  # 导入配置管理类
    config = Config()
    llm = LLM(config)

    markdown_content="""
        # Progress for langchain-ai/langchain (2024-08-20 to 2024-08-21)

        ## Issues Closed in the Last 1 Days
        - partners/chroma: release 0.1.3 #25599
        - docs: few-shot conceptual guide #25596
        - docs: update examples in api ref #25589
    """

    # 示例：生成 GitHub 报告
    system_prompt = "Your specific system prompt for GitHub report generation"
    github_report = llm.generate_report(system_prompt, markdown_content)
    print("----")
    LOG.info(f"github report created: {github_report}")
    print(github_report)
    print("----")
    LOG.debug(github_report)