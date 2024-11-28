import os
from openai import OpenAI

class LLM:
    def __init__(self, config=None):
        self.dry_run = config.llm_dry_run

        self.model = config.llm_model
        self.api_key = config.llm_api_key
        self.base_url = config.llm_base_url
        self.client = OpenAI(api_key=f"{self.api_key}", base_url=self.base_url)

    def generate_daily_report(self, markdown_content, dry_run=True):
        prompt = f"以下是项目的最新进展，根据功能合并同类项，形成一份简报，至少包含：1）新增功能；2）主要改进；3）修复问题；:\n\n{markdown_content}"
        if self.dry_run:
            with open("daily_progress/prompt.txt", "w+") as f:
                f.write(prompt)
            return "DRY RUN -> daily_progress/prompt.txt"

        print("Before call GPT")
        response = self.client.chat.completions.create(
            model=self.model,
            messages=[
                {"role": "system", "content": "Answer question < 200 words as a concise, software, full stack expert."},
                {"role": "user", "content": prompt}
            ],
            temperature=0.2
        )
        print("After call GPT")
        print(response)
        return response.choices[0].message.content