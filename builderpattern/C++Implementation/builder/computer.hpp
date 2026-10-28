#pragma once
#include <string>

class Computer {
private:
    std::string cpu;
    int ram;
    int storage;
    bool gpu;

public:
    Computer(std::string cpu, int ram, int storage, bool gpu)
        : cpu(cpu), ram(ram), storage(storage), gpu(gpu) {}

    void display() const;
};