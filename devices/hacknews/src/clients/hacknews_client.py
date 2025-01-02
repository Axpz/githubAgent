import requests
import sys
import os
from bs4 import BeautifulSoup
from datetime import datetime

sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))
from logger import logger


class HackerNewsClient:
    """A client to interact with Hacker News and fetch top stories."""
    
    def __init__(self):
        # Base URL for Hacker News
        self.url = 'https://news.ycombinator.com/'

    def fetch_top_stories(self):
        """
        Fetch the top stories from Hacker News.

        Returns:
            list: A list of dictionaries containing titles and links of top stories.
        """
        logger.debug("Fetching top stories from Hacker News.")
        try:
            response = requests.get(self.url, timeout=30)

            # Check if the response status code is successful
            response.raise_for_status()  

            return self.parse_stories(response.text)
        except requests.RequestException as e:
            logger.error(f"Failed to fetch Hacker News top stories: {e}")
            return []

    def parse_stories(self, html_content):
        """
        Parse the HTML content to extract top stories.

        Args:
            html_content (str): The HTML content of the Hacker News page.

        Returns:
            list: A list of dictionaries containing story titles and links.
        """
        logger.debug("Parsing HTML content to extract top stories.")
        soup = BeautifulSoup(html_content, 'html.parser')
        story_elements = soup.find_all('tr', class_='athing')

        top_stories = []
        for story in story_elements:
            title_tag = story.find('span', class_='titleline').find('a')
            if title_tag:
                top_stories.append({
                    'title': title_tag.text,
                    'link': title_tag['href']
                })

        logger.info(f"Successfully parsed {len(top_stories)} stories.")
        return top_stories

    def export_top_stories(self, date=None, hour=None):
        """
        Export top stories to a markdown file.

        Args:
            date (str, optional): The date in 'YYYY-MM-DD' format. Defaults to today's date.
            hour (str, optional): The hour in 'HH' format. Defaults to the current hour.

        Returns:
            str: The path of the generated markdown file, or None if no stories were found.
        """
        logger.debug("Exporting top stories to a markdown file.")
        top_stories = self.fetch_top_stories()

        if not top_stories:
            logger.warning("No top stories found to export.")
            return None

        # Use the current date and time if not provided
        date = date or datetime.now().strftime('%Y-%m-%d')
        hour = hour or datetime.now().strftime('%H')

        # Prepare the directory and file paths
        dir_path = os.path.join('reports/hacker_news', date)
        os.makedirs(dir_path, exist_ok=True)
        file_path = os.path.join(dir_path, f'{hour}.md')

        # Write stories to the markdown file
        with open(file_path, 'w') as file:
            file.write(f"# Hacker News Top Stories ({date} {hour}:00)\n\n")
            for idx, story in enumerate(top_stories, start=1):
                file.write(f"{idx}. [{story['title']}]({story['link']})\n")

        logger.info(f"Top stories exported to: {file_path}")
        return file_path


if __name__ == "__main__":
    # Create an instance of the client and export top stories
    client = HackerNewsClient()
    client.export_top_stories()
