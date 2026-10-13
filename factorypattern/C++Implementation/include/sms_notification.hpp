#pragma once
using namespace std;
#include "notification.hpp"

class SMSNotification : public Notification {
public:
    void send(const string& message) override;
};