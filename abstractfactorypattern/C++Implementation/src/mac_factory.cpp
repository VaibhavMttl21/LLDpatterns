#include "mac_factory.hpp"
#include "mac_button.hpp"
#include "mac_textbox.hpp"

Button* MacFactory::createButton() {
    return new MacButton();
}

TextBox* MacFactory::createTextBox() {
    return new MacTextBox();
}