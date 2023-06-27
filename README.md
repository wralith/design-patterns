# Design Patterns

Design Patterns with reasonable code examples -_i guess_-. 🤓

![cover image](docs/Cover.png)

You can check [docs](docs) folder for markdown files with in depth explanations
and some conditions and opportunities to use certain design pattern. 📕

Code examples for each language can be found in its separate folder _that
structured with what that languages conventions are_ with implementation of
pattern and the tests.

- To understand how to implement the pattern, check the example cases, don't
  worry there are comments for you to understand concepts and examples better.
- To understand how the code will be used in client, check test codes.

If you have any suggestion or spotted a mistake, feel free to reach out anytime.

## Creational Patterns 🏗

### Builder

Builder pattern provides a flexible and organized approach to constructing
complex objects.

[**In Depth Explanation 👈**](docs/Builder.md)

#### Code Example

|  Language  |                Implementation Code                |                      Test Code                      |
|:----------:|:-------------------------------------------------:|:---------------------------------------------------:|
|     Go     |       [Example Code](go/builder/builder.go)       |       [Test Code](go/builder/builder_test.go)       |
| TypeScript | [Example Code](typescript/src/builder/builder.ts) | [Test Code](typescript/src/builder/builder.test.ts) |

---

### Prototype

The Prototype is used to create objects by cloning or copying existing objects.

[**In Depth Explanation 👈**](docs/Prototype.md)

|  Language  |                  Implementation Code                  |                        Test Code                        |
|:----------:|:-----------------------------------------------------:|:-------------------------------------------------------:|
|     Go     |       [Example Code](go/prototype/prototype.go)       |       [Test Code](go/prototype/prototype_test.go)       |
| TypeScript | [Example Code](typescript/src/prototype/prototype.ts) | [Test Code](typescript/src/prototype/prototype.test.ts) |

---

### Singleton

The Singleton is used to ensure only one instance of a class is shared across
whole application.

[**In Depth Explanation 👈**](docs/Singleton.md)

|  Language  |                  Implementation Code                  |                        Test Code                        |
|:----------:|:-----------------------------------------------------:|:-------------------------------------------------------:|
|     Go     |       [Example Code](go/singleton/singleton.go)       |       [Test Code](go/singleton/singleton_test.go)       |
| TypeScript | [Example Code](typescript/src/singleton/singleton.ts) | [Test Code](typescript/src/singleton/singleton.test.ts) |
| C#         | [Example Code](csharp/Singleton/Logger/ConsoleLogger.cs) | [Test Code](csharp/Singleton/Logger.UnitTest/LoggerTest.cs) |

---

## Structural Patterns

### Adapter

The Adapter design pattern allows objects with incompatible interfaces to work together.

[**In Depth Explanation 👈**](docs/Adapter.md)

|  Language  |                Implementation Code                |                      Test Code                      |
|:----------:|:-------------------------------------------------:|:---------------------------------------------------:|
|     Go     |       [Example Code](go/adapter/adapter.go)       |       [Test Code](go/adapter/adapter_test.go)       |
| TypeScript | [Example Code](typescript/src/adapter/adapter.ts) | [Test Code](typescript/src/adapter/adapter.test.ts) |

---

## Behavioral Patterns

### Observer

The Observer pattern allows notifying multiple objects about an event that happened on the object they're observing.

[**In Depth Explanation 👈**](docs/Observer.md)

|  Language  |                 Implementation Code                 |                       Test Code                       |
|:----------:|:---------------------------------------------------:|:-----------------------------------------------------:|
|     Go     |       [Example Code](go/observer/observer.go)       |       [Test Code](go/observer/observer_test.go)       |
| TypeScript | [Example Code](typescript/src/observer/observer.ts) | [Test Code](typescript/src/observer/observer.test.ts) |
