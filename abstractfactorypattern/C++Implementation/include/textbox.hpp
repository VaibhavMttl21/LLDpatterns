#pragma once

class TextBox {
public:
    virtual void render() = 0;
    virtual ~TextBox() = default;
};