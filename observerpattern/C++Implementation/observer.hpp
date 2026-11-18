#pragma once
#include <string>

class Observer {
public:
    virtual void update(const std::string& videoTitle) = 0;
    virtual ~Observer() = default;
};