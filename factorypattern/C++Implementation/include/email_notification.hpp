#pragma once
using namespace std;
#include "notification.hpp"

class EmailNotification : public Notification {
public:
    void send(const string& message) override;
};