#pragma once
#include <iostream>
#include <string>

class ShippingService {
public:
    void arrangeShipping(const std::string& product) {
        std::cout << "Shipping arranged for: "
                  << product << '\n';
    }
};