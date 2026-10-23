#pragma once

#include "button.hpp"
#include "textbox.hpp"

class GUIFactory {
public:
    virtual Button* createButton() = 0;
    virtual TextBox* createTextBox() = 0;

    virtual ~GUIFactory() = default;
};