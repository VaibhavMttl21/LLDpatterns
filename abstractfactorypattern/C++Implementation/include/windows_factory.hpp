#pragma once

#include "gui_factory.hpp"

class WindowsFactory : public GUIFactory {
public:
    Button* createButton() override;
    TextBox* createTextBox() override;
};