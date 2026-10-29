#pragma once

#include "payment_processor.hpp"
#include "legacy_payment.hpp"

class PaymentAdapter : public PaymentProcessor {
private:
    LegacyPayment* legacyPayment;

public:
    PaymentAdapter(LegacyPayment* legacyPayment);

    void pay(double amount) override;
};