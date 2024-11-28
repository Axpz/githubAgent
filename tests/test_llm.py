import unittest
from unittest.mock import MagicMock, patch
from llm import LLM
from config import Config

# class TestNotifier(unittest.TestCase):
#     def test_notify(self):
#         # Add test cases for Notifier
#         pass


class TestLLM(unittest.TestCase):
    def setUp(self):
        self.config = Config()
        self.llm = LLM(config=self.config)
        for attr in dir(self.config):
            if not attr.startswith("__"):
                print(f"{attr}: {getattr(self.config, attr)}")
        print()

    def test_generate_daily_report_dry_run(self):
        pass

    def test_generate_daily_report_dry_run(self):
        self.config = Config()
        self.config.llm_dry_run = True
        self.llm = LLM(config=self.config)
        markdown_content = "Test content"
        result = self.llm.generate_daily_report(markdown_content)
        print(result)
        # self.assertEqual(result, "DRY RUN -> daily_progress/prompt.txt")
        # with open("daily_progress/prompt.txt", "r") as f:
        #     prompt = f.read()
        # self.assertIn("Test content", prompt)

if __name__ == '__main__':
    unittest.main()