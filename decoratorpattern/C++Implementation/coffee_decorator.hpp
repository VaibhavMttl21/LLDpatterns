#pragma once
#include "coffee.hpp"

class CoffeeDecorator : public Coffee {
protected:
    Coffee* coffee;

public:
    CoffeeDecorator(Coffee* coffee) {
        this->coffee = coffee;
    }

    virtual ~CoffeeDecorator() {
        delete coffee;
    }
};