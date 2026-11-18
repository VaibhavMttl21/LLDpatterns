#pragma once

#include "payment_strategy.hpp"

class ShoppingCart {
private:
    PaymentStrategy* strategy;

public:
    ShoppingCart(PaymentStrategy* strategy);

    void setPaymentStrategy(PaymentStrategy* strategy);

    void checkout(double amount);
};