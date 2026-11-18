# Design Patterns in Go & C++

A hands-on collection of **Object-Oriented Design Patterns** implemented
in **Go and C++**.

This repository demonstrates how common design patterns solve recurring
software-design problems through practical examples in both languages.
It is intended to strengthen OOP fundamentals and build a foundation for
Low-Level Design (LLD).

## Design Patterns Covered

  -----------------------------------------------------------------------------------
  \#                      Pattern                             What it solves
  ----------------------- ----------------------------------- -----------------------
  1                       [Factory](factorypattern/)          Creates objects without
                                                              exposing the creation
                                                              logic to the client.

  2                       [Abstract                           Creates families of
                          Factory](abstractfactorypattern/)   related objects.

  3                       [Singleton](singletonpattern/)      Ensures a class has a
                                                              single shared instance.

  4                       [Builder](builderpattern/)          Constructs complex
                                                              objects step by step.

  5                       [Adapter](adapterpattern/)          Makes incompatible
                                                              interfaces work
                                                              together.

  6                       [Decorator](decoratorpattern/)      Adds responsibilities
                                                              to an object
                                                              dynamically.

  7                       [Facade](facadepattern/)            Provides a simple
                                                              interface to a complex
                                                              subsystem.

  8                       [Observer](observerpattern/)        Notifies registered
                                                              observers when a
                                                              subject changes.

  9                       [Strategy](strategypattern/)        Makes algorithms or
                                                              behaviours
                                                              interchangeable at
                                                              runtime.
  -----------------------------------------------------------------------------------

## How the Patterns Work

The interaction diagrams for all nine patterns are maintained separately
to keep this README easy to navigate:

➡️ **[View the Design Pattern Flowcharts](DESIGN_PATTERN_FLOWS.md)**

The diagrams illustrate the main participants and the flow of
interactions in each example.

## Repository Structure

Each pattern has its own directory, with separate Go and C++
implementations.

``` text
design-patterns/
├── factorypattern/
│   ├── GoImplementation/
│   └── C++Implementation/
├── abstractfactorypattern/
│   ├── GoImplementation/
│   └── C++Implementation/
├── singletonpattern/
│   ├── GoImplementation/
│   └── C++Implementation/
├── builderpattern/
│   ├── GoImplementation/
│   └── C++Implementation/
├── adapterpattern/
│   ├── GoImplementation/
│   └── C++Implementation/
├── decoratorpattern/
│   ├── GoImplementation/
│   └── C++Implementation/
├── facadepattern/
│   ├── GoImplementation/
│   └── C++Implementation/
├── observerpattern/
│   ├── GoImplementation/
│   └── C++Implementation/
├── strategypattern/
│   ├── GoImplementation/
│   └── C++Implementation/
├── DESIGN_PATTERN_FLOWS.md
└── README.md
```

## Technologies and Concepts

-   **Go:** interfaces, structs, methods, and composition.
-   **C++:** classes, inheritance, polymorphism, and object composition.
-   **OOP:** encapsulation, abstraction, inheritance, and polymorphism.

## Getting Started

### Prerequisites

-   [Go](https://go.dev/dl/)
-   A C++ compiler such as [G++](https://gcc.gnu.org/)
-   Git

### Clone the repository

``` bash
git clone <your-repository-url>
cd design-patterns
```

### Run a Go example

Navigate to a pattern's Go implementation directory:

``` bash
cd factorypattern/GoImplementation
go run .
```

### Compile and run a C++ example

Navigate to a pattern's C++ implementation directory:

``` bash
cd factorypattern/C++Implementation
g++ -std=c++17 -I . *.cpp -o pattern_demo
./pattern_demo
```

> **Note:** The exact compilation command can vary depending on the
> source files in a pattern. Use the relevant source files if the
> directory contains multiple independent `main()` functions.

## Learning Path

For each pattern:

1.  Understand the problem it addresses.
2.  Review the interaction flow in
    [DESIGN_PATTERN_FLOWS.md](DESIGN_PATTERN_FLOWS.md).
3.  Explore the Go implementation.
4.  Study the equivalent C++ implementation.
5.  Compile, run, and experiment with the example.

## Future Scope

Additional patterns and practical LLD examples may be added over time.
Potential next patterns include State, Command, Template Method, Proxy,
Chain of Responsibility, Composite, Bridge, Iterator, Mediator, and
Visitor.

## Contributing

This is a personal learning repository. Suggestions, improvements, and
discussions are welcome. Feel free to open an issue or submit a pull
request.

------------------------------------------------------------------------

**Happy Learning & Building!**
