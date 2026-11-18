#pragma once

#include <string>

#include "inventory_service.hpp"
#include "payment_service.hpp"
#include "shipping_service.hpp"
#include "notification_service.hpp"

class OrderFacade {
private:
    InventoryService inventory;
    PaymentService payment;
    ShippingService shipping;
    NotificationService notification;

public:
    void placeOrder(const std::string& product, double amount);
};