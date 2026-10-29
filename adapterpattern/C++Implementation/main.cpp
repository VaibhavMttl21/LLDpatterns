#include <iostream>
#include "adapter/payment_adapter.hpp"

int main() {
    LegacyPayment* legacy = new LegacyPayment();

    PaymentProcessor* processor = new PaymentAdapter(legacy);

    processor->pay(1500.0);

    delete processor;
    delete legacy;

    return 0;
}