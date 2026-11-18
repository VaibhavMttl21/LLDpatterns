#pragma once
#include "coffee_decorator.hpp"

class ChocolateDecorator : public CoffeeDecorator {
public:
    ChocolateDecorator(Coffee* coffee) : CoffeeDecorator(coffee) {}

    std::string description() override {
        return coffee->description() + ", Chocolate";
    }

    double cost() override {
        return coffee->cost() + 40;
    }
};