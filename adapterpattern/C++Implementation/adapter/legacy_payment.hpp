#pragma once
#include <iostream>

class LegacyPayment {
public:
    void makePaymentInRupees(double amount) {
        std::cout << "Processing legacy payment of Rs. "
                  << amount << '\n';
    }
};