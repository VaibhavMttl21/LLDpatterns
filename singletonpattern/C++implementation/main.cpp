#include "logger.hpp"

int main() {
    Logger* logger1 = Logger::getInstance();
    Logger* logger2 = Logger::getInstance();

    logger1->log("Application started");
    logger2->log("User logged in");

    if (logger1 == logger2) {
        // Both pointers refer to the same object
        logger1->log("Both references use the same logger");
    }

    return 0;
}