#pragma once
#include "payment_strategy.hpp"

class UPIPayment : public PaymentStrategy {
public:
    void pay(double amount) override;
};