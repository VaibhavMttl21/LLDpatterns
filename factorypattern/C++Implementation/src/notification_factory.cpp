#include "notification_factory.hpp"
#include "email_notification.hpp"
#include "sms_notification.hpp"

#include <stdexcept>
using namespace std;
Notification* NotificationFactory::createNotification(
    const string& type
) {
    if (type == "email") {
        return new EmailNotification();
    }

    if (type == "sms") {
        return new SMSNotification();
    }

    throw invalid_argument("Unknown notification type: " + type);
}