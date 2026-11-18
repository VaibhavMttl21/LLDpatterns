#pragma once
#include <iostream>

class PaymentService {
public:
    bool processPayment(double amount) {
        std::cout << "Processing payment of Rs. "
                  << amount << '\n';

        return true;
    }
};