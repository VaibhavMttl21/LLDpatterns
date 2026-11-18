#pragma once

#include "observer.hpp"
#include <string>

class Subscriber : public Observer {
private:
    std::string name;

public:
    Subscriber(const std::string& name);

    void update(const std::string& videoTitle) override;
};