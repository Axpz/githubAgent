import argparse
import threading
import time
from core.subscription_manager import SubscriptionManager
from core.update_fetcher import UpdateFetcher
from core.notifier import Notifier
from core.report_generator import ReportGenerator

def main():
    parser = argparse.ArgumentParser(description="Interactive Tool for Managing Subscriptions and Updates")
    parser.add_argument('-a', '--add-subscription', type=str, nargs='?', const='', help='Add a new subscription')
    parser.add_argument('-r', '--remove-subscription', type=str, help='Remove an existing subscription')
    parser.add_argument('-f', '--fetch-updates', action='store_true', help='Fetch updates immediately')
    args = parser.parse_args()

    subscription_manager = SubscriptionManager()
    fetcher = UpdateFetcher()
    notifier = Notifier()
    report_generator = ReportGenerator()

    if args.add_subscription:
        subscription_manager.add(args.add_subscription)
        print(f"Added subscription: {args.add_subscription}")

    if args.remove_subscription:
        subscription_manager.remove(args.remove_subscription)
        print(f"Removed subscription: {args.remove_subscription}")

    if args.fetch_updates:
        repositories = subscription_manager.get_subscription()
        updates = fetcher.fetch_updates(repositories)
        if updates:
            notifier.notify(updates)
        report_generator.generate_report(updates)
        print("Fetched and processed updates")

    # Start the scheduler in the background
    scheduler_thread = threading.Thread(target=scheduler, args=(subscription_manager, fetcher, notifier, report_generator))
    scheduler_thread.daemon = True
    scheduler_thread.start()

    # Keep the main thread alive to accept commands
    while True:
        time.sleep(1)

def scheduler(subscription_manager, fetcher, notifier, report_generator):
    while True:
        repositories = subscription_manager.get_subscription()
        updates = fetcher.fetch_updates(repositories)
        if updates:
            notifier.notify(updates)
        report_generator.generate_report(updates)
        time.sleep(3 * 3600)  # Run every three hour

if __name__ == "__main__":
    main()
