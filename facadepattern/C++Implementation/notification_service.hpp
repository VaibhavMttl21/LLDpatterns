#pragma once
#include <iostream>

class NotificationService {
public:
    void sendConfirmation() {
        std::cout << "Order confirmation sent!\n";
    }
};