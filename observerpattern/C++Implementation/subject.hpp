#pragma once
#include "observer.hpp"

class Subject {
public:
    virtual void subscribe(Observer* observer) = 0;
    virtual void unsubscribe(Observer* observer) = 0;
    virtual void notify(const std::string& videoTitle) = 0;

    virtual ~Subject() = default;
};