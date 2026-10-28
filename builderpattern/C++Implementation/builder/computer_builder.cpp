#include "computer_builder.hpp"
#include <iostream>
#include <utility>

ComputerBuilder& ComputerBuilder::setCPU(const std::string& cpu) {
    this->cpu = cpu;
    return *this;
}

ComputerBuilder& ComputerBuilder::setRAM(int ram) {
    this->ram = ram;
    return *this;
}

ComputerBuilder& ComputerBuilder::setStorage(int storage) {
    this->storage = storage;
    return *this;
}

ComputerBuilder& ComputerBuilder::setGPU(bool gpu) {
    this->gpu = gpu;
    return *this;
}

Computer* ComputerBuilder::build() {
    return new Computer(cpu, ram, storage, gpu);
}