#include <iostream>
#include <string>

#include "gui_factory.hpp"
#include "mac_factory.hpp"
#include "windows_factory.hpp"

int main() {
    std::string os;

    std::cout << "Choose operating system (mac/windows): ";
    std::cin >> os;

    GUIFactory* factory = nullptr;

    if (os == "mac") {
        factory = new MacFactory();
    } else if (os == "windows") {
        factory = new WindowsFactory();
    } else {
        std::cout << "Unsupported operating system\n";
        return 1;
    }

    Button* button = factory->createButton();
    TextBox* textbox = factory->createTextBox();

    button->render();
    textbox->render();

    delete button;
    delete textbox;
    delete factory;

    return 0;
}