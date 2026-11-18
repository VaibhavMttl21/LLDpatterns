#include "shopping_cart.hpp"
#include "upi_payment.hpp"
#include "card_payment.hpp"
#include "paypal_payment.hpp"

int main() {
    PaymentStrategy* upi = new UPIPayment();
    PaymentStrategy* card = new CardPayment();
    PaymentStrategy* paypal = new PayPalPayment();

    ShoppingCart cart(upi);

    cart.checkout(1500);

    // Change strategy at runtime
    cart.setPaymentStrategy(card);
    cart.checkout(2500);

    cart.setPaymentStrategy(paypal);
    cart.checkout(3500);

    delete upi;
    delete card;
    delete paypal;

    return 0;
}