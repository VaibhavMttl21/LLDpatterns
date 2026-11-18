#include "order_facade.hpp"
#include <iostream>

void OrderFacade::placeOrder(
    const std::string& product,
    double amount
) {
    std::cout << "\nStarting order process...\n\n";

    // Step 1: Check inventory
    if (!inventory.checkAvailability(product)) {
        std::cout << "Product is unavailable!\n";
        return;
    }

    // Step 2: Process payment
    if (!payment.processPayment(amount)) {
        std::cout << "Payment failed!\n";
        return;
    }

    // Step 3: Arrange shipping
    shipping.arrangeShipping(product);

    // Step 4: Send confirmation
    notification.sendConfirmation();

    std::cout << "\nOrder placed successfully!\n";
}