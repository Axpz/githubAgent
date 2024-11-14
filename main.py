from core.subscription import SubscriptionManager
from core.update_fetcher import UpdateFetcher
from core.notifier import Notifier
from core.report_generator import ReportGenerator

def main():
    subscription_manager = SubscriptionManager()
    fetcher = UpdateFetcher()
    notifier = Notifier()
    report_generator = ReportGenerator()
    
    repositories = subscription_manager.get_subscribed_repositories()
    updates = fetcher.fetch_updates(repositories)
    
    if updates:
        notifier.notify(updates)
    
    report_generator.generate_report(updates)

if __name__ == "__main__":
    main()
