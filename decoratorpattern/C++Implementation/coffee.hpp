#pragma once
#include <string>

class Coffee {
public:
    virtual std::string description() = 0;
    virtual double cost() = 0;

    virtual ~Coffee() = default;
};