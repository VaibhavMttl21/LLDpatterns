#pragma once

#include "gui_factory.hpp"

class MacFactory : public GUIFactory {
public:
    Button* createButton() override;
    TextBox* createTextBox() override;
};