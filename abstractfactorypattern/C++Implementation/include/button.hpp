#pragma once

class Button {
public:
    virtual void render() = 0;
    virtual ~Button() = default;
};