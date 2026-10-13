#include "email_notification.hpp"
using namespace std;
#include <iostream>

void EmailNotification::send(const string& message) {
    cout << "Sending Email: " << message << '\n';
}