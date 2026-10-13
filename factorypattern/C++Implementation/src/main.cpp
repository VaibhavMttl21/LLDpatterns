#include "notification_factory.hpp"

#include <iostream>
#include <string>
#include <stdexcept>

using namespace std;

int main()
{
    Notification *notification = nullptr;

    try
    {
        string type;

        cout << "Choose notification type (email/sms): ";
        cin >> type;

        notification = NotificationFactory::createNotification(type);

        notification->send("Your order has been confirmed!");

        delete notification;
        notification = nullptr;
    }
    catch (const invalid_argument &error)
    {
        delete notification;
        cerr << "Error: " << error.what() << '\n';
        return 1;
    }

    return 0;
}

// Abstract class	        Cannot be instantiated directly if it has a pure virtual function.
// Pure virtual function	Declared with = 0; requires concrete derived classes to implement it.
// Inheritance	            A derived class acquires or extends behavior from a base class.
// virtual	                Enables dynamic dispatch for virtual functions.
// override	                Makes the compiler verify that you're overriding a virtual function.
// Base-class pointer	    Can point to a derived object when the inheritance relationship allows it.
// Runtime polymorphism	    A virtual call selects the implementation according to the actual object.
// Virtual destructor	    Enables safe destruction through a base-class pointer.
// new	                    Dynamically creates an object and returns a pointer.
// delete	                estroys the object and releases its allocated memory.