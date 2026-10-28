#include "computer.hpp"
#include <iostream>

void Computer::display() const {
    std::cout << "CPU: " << cpu << '\n';
    std::cout << "RAM: " << ram << " GB\n";
    std::cout << "Storage: " << storage << " GB\n";
    std::cout << "GPU: " << (gpu ? "Yes" : "No") << '\n';
}