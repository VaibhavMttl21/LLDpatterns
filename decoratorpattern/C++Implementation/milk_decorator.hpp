#pragma once
#include "coffee_decorator.hpp"

class MilkDecorator : public CoffeeDecorator {
public:
    MilkDecorator(Coffee* coffee) : CoffeeDecorator(coffee) {}

    std::string description() override {
        return coffee->description() + ", Milk";
    }

    double cost() override {
        return coffee->cost() + 30;
    }
};