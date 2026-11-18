#pragma once
#include "payment_strategy.hpp"

class CardPayment : public PaymentStrategy {
public:
    void pay(double amount) override;
};