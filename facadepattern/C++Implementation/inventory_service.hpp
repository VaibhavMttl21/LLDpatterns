#pragma once
#include <iostream>
#include <string>

class InventoryService {
public:
    bool checkAvailability(const std::string& product) {
        std::cout << "Checking inventory for: " << product << '\n';

        // Assume product is available
        return true;
    }
};