#include "payment_adapter.hpp"

PaymentAdapter::PaymentAdapter(LegacyPayment* legacyPayment) {
    this->legacyPayment = legacyPayment;
}

void PaymentAdapter::pay(double amount) {
    legacyPayment->makePaymentInRupees(amount);
}