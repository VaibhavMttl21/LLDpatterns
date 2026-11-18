# Design Pattern Flowcharts

This document contains interaction flowcharts for the nine design
patterns currently implemented in this repository.

Each diagram shows the key participants and how a request or interaction
moves through them. These are conceptual diagrams based on the
repository's examples, not line-by-line source-code maps.

## 1. Factory Pattern

**Example:** Notification creation.

The client asks the factory for a notification. The factory chooses the
concrete implementation and returns it through the common `Notification`
interface.

``` mermaid
flowchart TD
    A["Client"] -->|"createNotification(type)"| B["Notification Factory"]
    B --> C{"Notification type?"}
    C -->|"Email"| D["Email Notification"]
    C -->|"SMS"| E["SMS Notification"]
    D --> F["Notification interface"]
    E --> F
    F --> G["Client calls notify()"]
```

## 2. Abstract Factory Pattern

**Example:** Cross-platform GUI components.

The client chooses a concrete factory. That factory creates a matching
family of products, such as a button and a text box for the same
operating system.

``` mermaid
flowchart TD
    A["Client"] --> B{"Choose UI Factory"}
    B --> C["Mac Factory"]
    B --> D["Windows Factory"]
    C --> E["Mac Button"]
    C --> F["Mac TextBox"]
    D --> G["Windows Button"]
    D --> H["Windows TextBox"]
    E --> I["Button interface"]
    G --> I
    F --> J["TextBox interface"]
    H --> J
    I --> K["Client uses UI components"]
    J --> K
```

## 3. Singleton Pattern

**Example:** Logger.

Clients call `getInstance()`. The class creates the instance if it does
not exist; otherwise, it returns the existing instance.

``` mermaid
flowchart TD
    A["Client 1"] --> C["Logger.getInstance()"]
    B["Client 2"] --> C
    C --> D{"Does instance exist?"}
    D -->|"No"| E["Create Logger instance"]
    D -->|"Yes"| F["Return existing instance"]
    E --> G["Shared Logger instance"]
    F --> G
    G --> H["Client uses Logger"]
```

## 4. Builder Pattern

**Example:** Computer construction.

The client configures the builder through a sequence of setter calls.
Calling `build()` produces the completed `Computer` object.

``` mermaid
flowchart TD
    A["Client"] --> B["Computer Builder"]
    B --> C["setCPU()"]
    C --> D["setRAM()"]
    D --> E["setStorage()"]
    E --> F["setGPU()"]
    F --> G["build()"]
    G --> H["Computer object"]
```

## 5. Adapter Pattern

**Example:** Payment integration with a legacy API.

The client calls the expected `PaymentProcessor` interface. The adapter
translates that call into the method understood by the legacy payment
system.

``` mermaid
flowchart TD
    A["Client"] --> B["PaymentProcessor interface"]
    B --> C["Payment Adapter"]
    C -->|"Translates pay(amount)"| D["Legacy Payment"]
    D --> E["makePaymentInRupees(amount)"]
    E --> F["Payment processed"]
```

## 6. Decorator Pattern

**Example:** Coffee with optional add-ons.

Each decorator implements the same `Coffee` interface and wraps another
coffee object. Each layer adds its own description and cost.

``` mermaid
flowchart TD
    A["Client"] --> B["Chocolate Decorator"]
    B --> C["Milk Decorator"]
    C --> D["Plain Coffee"]
    D --> E["Base description and cost"]
    E --> C
    C --> F["Add milk description and cost"]
    F --> B
    B --> G["Add chocolate description and cost"]
    G --> H["Final decorated coffee"]
```

## 7. Facade Pattern

**Example:** Online order placement.

The client calls one method on `OrderFacade`. The facade coordinates
inventory, payment, shipping, and notification services, and stops the
workflow if an earlier step fails.

``` mermaid
flowchart TD
    A["Client"] --> B["Order Facade"]
    B --> C["Inventory Service"]
    C --> D{"Item available?"}
    D -->|"No"| E["Order fails"]
    D -->|"Yes"| F["Payment Service"]
    F --> G{"Payment successful?"}
    G -->|"No"| H["Order fails"]
    G -->|"Yes"| I["Shipping Service"]
    I --> J["Notification Service"]
    J --> K["Order completed"]
```

## 8. Observer Pattern

**Example:** YouTube channel and subscribers.

Subscribers register with the channel. When a new video is uploaded, the
channel notifies its registered observers through `update()`.

``` mermaid
flowchart TD
    A["Subscriber 1"] -->|"subscribe"| B["YouTube Channel (Subject)"]
    C["Subscriber 2"] -->|"subscribe"| B
    D["Subscriber 3"] -->|"subscribe"| B
    E["New video uploaded"] --> B
    B --> F["notify()"]
    F --> G["Subscriber 1: update()"]
    F --> H["Subscriber 2: update()"]
    F --> I["Subscriber 3: update()"]
```

## 9. Strategy Pattern

**Example:** Payment methods in a shopping cart.

The shopping cart (context) uses the `PaymentStrategy` interface. The
client selects a concrete strategy, and the cart delegates payment
execution to it.

``` mermaid
flowchart TD
    A["Client"] --> B["Shopping Cart (Context)"]
    B --> C["PaymentStrategy interface"]
    C --> D["UPI Payment"]
    C --> E["Card Payment"]
    C --> F["PayPal Payment"]
    A --> G["Select or change strategy"]
    G --> B
    B --> H["checkout()"]
    H --> I["Selected strategy.pay(amount)"]
    I --> J["Payment result"]
```

------------------------------------------------------------------------

**Note:** The diagrams are simplified views of the example
implementations. They focus on the roles and interactions that make each
pattern work.
