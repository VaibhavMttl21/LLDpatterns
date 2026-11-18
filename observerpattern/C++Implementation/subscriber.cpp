#include "subscriber.hpp"
#include <iostream>

Subscriber::Subscriber(const std::string& name) {
    this->name = name;
}

void Subscriber::update(const std::string& videoTitle) {
    std::cout << name << " received notification: New video - "
              << videoTitle << '\n';
}