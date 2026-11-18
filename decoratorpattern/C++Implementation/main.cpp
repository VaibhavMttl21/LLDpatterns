#include <iostream>
#include "plain_coffee.hpp"
#include "milk_decorator.hpp"
#include "chocolate_decorator.hpp"

int main() {
    Coffee* coffee = new PlainCoffee();

    coffee = new MilkDecorator(coffee);
    coffee = new ChocolateDecorator(coffee);

    std::cout << "Coffee: " << coffee->description() << '\n';
    std::cout << "Total cost: Rs. " << coffee->cost() << '\n';

    delete coffee;

    return 0;
}