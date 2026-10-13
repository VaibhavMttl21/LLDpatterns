#pragma once
using namespace std;
#include <string>

class Notification {
public:
    virtual void send(const string& message) = 0;

    virtual ~Notification() = default;
};