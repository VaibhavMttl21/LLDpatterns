#pragma once

#include "subject.hpp"
#include <vector>
#include <string>

class YouTubeChannel : public Subject {
private:
    std::vector<Observer*> observers;

public:
    void subscribe(Observer* observer) override;
    void unsubscribe(Observer* observer) override;
    void notify(const std::string& videoTitle) override;

    void uploadVideo(const std::string& videoTitle);
};