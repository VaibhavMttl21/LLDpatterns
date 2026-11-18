#include "shopping_cart.hpp"

ShoppingCart::ShoppingCart(PaymentStrategy* strategy) {
    this->strategy = strategy;
}

void ShoppingCart::setPaymentStrategy(PaymentStrategy* strategy) {
    this->strategy = strategy;
}

void ShoppingCart::checkout(double amount) {
    strategy->pay(amount);
}