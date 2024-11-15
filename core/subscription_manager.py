# core/subscription_manager.py

class SubscriptionManager:
    def __init__(self):
        self.subscriptions = []

    def get_subscription(self):
        print("Getting subscriptions")
        return self.subscriptions

    def add_subscription(self, repository):
        if repository not in self.subscriptions:
            self.subscriptions.append(repository)
            print(f"Added subscription: {repository}")
        else:
            print(f"Subscription already exists: {repository}")

    def remove_subscription(self, repository):
        if repository in self.subscriptions:
            self.subscriptions.remove(repository)
            print(f"Removed subscription: {repository}")
        else:
            print(f"Subscription not found: {repository}")