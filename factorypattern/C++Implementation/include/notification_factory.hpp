// The factory can return either an EmailNotification or an SMSNotification while the client only needs to know about Notification.
// std::unique_ptr also manages the object's lifetime automatically, so we don't need manual delete calls.


#pragma once
using namespace std;
#include <string>
#include "notification.hpp"

class NotificationFactory {
public:
    static Notification* createNotification(const std::string& type);
};