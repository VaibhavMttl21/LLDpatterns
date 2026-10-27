#pragma once
#include <string>

class Logger {
private:
    // The single instance
    static Logger* instance;

    // Prevent creating objects directly
    Logger();

    // Prevent copying
    Logger(const Logger&) = delete;
    Logger& operator=(const Logger&) = delete;

public:
    static Logger* getInstance();

    void log(const std::string& message);
};