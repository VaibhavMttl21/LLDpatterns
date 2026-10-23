#include "windows_factory.hpp"
#include "windows_button.hpp"
#include "windows_textbox.hpp"

Button* WindowsFactory::createButton() {
    return new WindowsButton();
}

TextBox* WindowsFactory::createTextBox() {
    return new WindowsTextBox();
}