#include "youtube_channel.hpp"
#include "subscriber.hpp"

int main() {
    // Create the channel
    YouTubeChannel* channel = new YouTubeChannel();

    // Create subscribers
    Subscriber* s1 = new Subscriber("Vaibhav");
    Subscriber* s2 = new Subscriber("Aman");
    Subscriber* s3 = new Subscriber("Karan");

    // Register subscribers
    channel->subscribe(s1);
    channel->subscribe(s2);
    channel->subscribe(s3);

    // Upload a video
    channel->uploadVideo("Observer Design Pattern in C++");

    // Unsubscribe one subscriber
    channel->unsubscribe(s2);

    // Upload another video
    channel->uploadVideo("Strategy Design Pattern in C++");

    // Clean up
    delete channel;
    delete s1;
    delete s2;
    delete s3;

    return 0;
}