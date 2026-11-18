#include <iostream>

#include "upi_payment.hpp"
#include "card_payment.hpp"
#include "paypal_payment.hpp"

void UPIPayment::pay(double amount) {
    std::cout << "Paid Rs. " << amount << " using UPI\n";
}

void CardPayment::pay(double amount) {
    std::cout << "Paid Rs. " << amount << " using Credit/Debit Card\n";
}

void PayPalPayment::pay(double amount) {
    std::cout << "Paid Rs. " << amount << " using PayPal\n";
}