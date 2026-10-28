#include <iostream>
#include "builder/computer_builder.hpp"

int main() {
    ComputerBuilder builder;

    Computer* pc = builder
        .setCPU("Intel i7")
        .setRAM(32)
        .setStorage(1024)
        .setGPU(true)
        .build();

    pc->display();

    delete pc;
    return 0;
}