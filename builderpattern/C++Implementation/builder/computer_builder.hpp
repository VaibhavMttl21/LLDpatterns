#pragma once
#include "computer.hpp"

class ComputerBuilder {
private:
    std::string cpu;
    int ram = 0;
    int storage = 0;
    bool gpu = false;

public:
    ComputerBuilder& setCPU(const std::string& cpu);
    ComputerBuilder& setRAM(int ram);
    ComputerBuilder& setStorage(int storage);
    ComputerBuilder& setGPU(bool gpu);

    Computer* build();
};