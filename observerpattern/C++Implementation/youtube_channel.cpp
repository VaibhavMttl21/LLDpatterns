#include "youtube_channel.hpp"
#include <iostream>
#include <algorithm>

void YouTubeChannel::subscribe(Observer* observer) {
    observers.push_back(observer);
    std::cout << "New subscriber added!\n";
}

void YouTubeChannel::unsubscribe(Observer* observer) {
    auto it = std::find(observers.begin(), observers.end(), observer);

    if (it != observers.end()) {
        observers.erase(it);
        std::cout << "Subscriber removed!\n";
    }
}

void YouTubeChannel::notify(const std::string& videoTitle) {
    for (Observer* observer : observers) {
        observer->update(videoTitle);
    }
}

void YouTubeChannel::uploadVideo(const std::string& videoTitle) {
    std::cout << "\nNew video uploaded: " << videoTitle << '\n';

    notify(videoTitle);
}