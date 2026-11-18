#pragma once
#include "coffee.hpp"

class PlainCoffee : public Coffee {
public:
    std::string description() override;
    double cost() override;
};