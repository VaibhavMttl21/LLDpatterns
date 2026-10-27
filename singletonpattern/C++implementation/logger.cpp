#include "logger.hpp"
#include <iostream>

// Initially, no instance exists
Logger* Logger::instance = nullptr;

Logger::Logger() {
    std::cout << "Logger created\n";
}

Logger* Logger::getInstance() {
    if (instance == nullptr) {
        instance = new Logger();
    }

    return instance;
}

void Logger::log(const std::string& message) {
    std::cout << "[LOG]: " << message << '\n';
}