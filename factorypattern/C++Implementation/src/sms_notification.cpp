#include "sms_notification.hpp"
using namespace std;
#include <iostream>

void SMSNotification::send(const string& message) {
    cout << "Sending SMS: " << message << '\n';
}