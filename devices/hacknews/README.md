# Hacker news

This is an open-source AI-based tool designed for developers and project managers to efficiently track and the updates in subscribed hacker news and github repositories. By providing regular (daily/weekly) summaries, GitHub Agent helps teams stay informed on project progress, enabling faster responses and more effective collaboration.

## Features

- **Subscription Management**: Easily manage GitHub repository subscriptions.
- **Update Fetching**: Automatically fetch the latest updates from subscribed repositories.
- **Notification System**: Notify users of new updates via email, Slack, etc.
- **Report Generation**: Generate regular update reports (daily/weekly) to keep track of project changes.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/axpz/githubAgent.git
   cd githubAgent/devices/hacknews/
   ```

2. Activate your Python environment:

   ```bash
   source ~/python/env/jupyter/bin/activate
   ```

3. Install dependencies:

   ```bash
   pip install -r requirements.txt
   ```

## Usage

1. Set up your repository subscriptions: Use the subscription management module (subscription.py) to add or remove repositories you want to track.

2. Run GitHub Agent: To start fetching updates and generating reports, run:

   ```bash
   cd src && python main.py
   ```

3. Configure Notifications: Adjust the notifier.py module to set up your preferred notification channels (e.g., email, Slack).

4. Generate Reports: The agent will automatically generate reports at the specified intervals. You can find the generated report in `reports`.
