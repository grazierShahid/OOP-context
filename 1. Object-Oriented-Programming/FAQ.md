# OOP — FAQ

Short intro:  
We finished the main OOP context. This FAQ collects common interview questions, short answers, and quick comparisons you will be asked. Keep answers simple, practical, and business-oriented — like you explain to a teammate.

---

## Definition
FAQ: short answers to clarify common confusions quickly.

---

## Abstraction vs Interface vs Encapsulation — quick
- Abstraction: hide implementation, expose only essential behavior (what). Use abstract classes or interfaces when you want a simpler view.
- Interface: a contract (method signatures). Use for capability/role and loose coupling. Since Java 8 interfaces can have default/static helpers.
- Encapsulation: bundle data + methods and restrict direct access (how). Use access modifiers, getters/setters, validation to protect state.

When to use:
- Interface: unrelated types must provide same behavior.
- Abstract class: related types share code/state.
- Encapsulation: always.

---

## Compile-time vs Run-time polymorphism
- Compile-time (static): decided by compiler. Example: method overloading, static binding.
- Run-time (dynamic): decided at runtime based on actual object type. Example: method overriding, virtual dispatch.

Java note: instance methods use dynamic dispatch by default. static/private methods are resolved at compile time.

---

## Overloading vs Overriding — rules + tiny examples
- Overloading: same name, different parameter lists in same class (compiler chooses). Return type alone cannot distinguish.
- Overriding: subclass supplies same-signature method; runtime chooses based on object. Use @Override.

```java
// Overloading (compile-time)
class Calc { int add(int a,int b){return a+b;} double add(double a,double b){return a+b;} }

// Overriding (runtime)
class P{ void show(){System.out.println("P");} }
class C extends P{ @Override void show(){System.out.println("C");} }
```

---

## Abstract class vs Interface — when & trade-offs
- Interface: use for pure contracts, multiple capability inheritance, mocking in tests.
- Abstract class: use when sharing fields or partial behavior among close types.
- Trade-off: class can implement many interfaces but extend only one abstract class.

Short rule: prefer interface for API/DI; abstract class when code/state sharing simplifies design.

---

## Diamond problem
- What: ambiguity from multiple inheritance when same ancestor appears twice.
- Java: prevents multiple class inheritance; interfaces avoid class-diamond. Default methods can create conflicts — implementer must resolve.
- C++: supports multiple inheritance; use virtual inheritance or explicit qualification to resolve duplicates.

---

## Composition vs Inheritance
- Rule of thumb: prefer composition (has-a) over inheritance (is-a).
- Use inheritance for true "is-a" relationships and polymorphism; composition for reuse without tight coupling.

Example idea: Car has Engine (composition), not Car extends Engine.

---

## static vs instance members
- static: class-level, one copy, accessed via ClassName.member, hidden (not overridden).
- instance: per-object, supports overriding (dynamic dispatch for methods).

---

## Resource cleanup — finalize(), destructors, best practice
- Java: finalize() is deprecated and unreliable. Use try-with-resources, AutoCloseable, explicit close(), or Cleaner.
- Go: no destructors; use defer for deterministic cleanup and rely on GC for memory.

---

## Go quick notes
- Interfaces are implicit — a type implements an interface by providing methods.
- Receiver matters: use pointer receiver to mutate or avoid copies; value receiver otherwise.
- Exported names start with capital letters; unexported start lowercase.
- Embedding gives composition and method promotion.

---

## Extra interview questions to practice
- Explain OOP pillars in one sentence each.
- Overloading vs overriding: give examples and binding times.
- Abstract class vs interface: when to choose which.
- What is method hiding and how does it differ from overriding?
- Show equals/hashCode contract and a common bug.
- Explain LSP with a code example that breaks it.
- How to manage resources in Java? (try-with-resources)
- How do default methods in interfaces change polymorphism?
- What is dependency inversion and why use it?
- When to use composition over inheritance? Give an example.
- How does JVM choose which overridden method to call?
- What are pitfalls of multithreading with mutable objects? (brief)

---

## Short model answers (copy-paste ready)
- "OOP groups data and behavior into objects; main benefits: modularity, reuse, maintainability."
- "Abstraction hides details; encapsulation protects state; interface defines contract."
- "Compile-time polymorphism = overloading; runtime = overriding (virtual dispatch)."
- "Prefer composition over inheritance to reduce coupling and increase flexibility."
- "Use interfaces for DI and testing; use abstract classes to share common code/state."

---

## Closing note (your style)
Study these short Q&A and practice small code snippets. In interviews give the concise definition first, then a tiny example, then a short trade-off or when-to-use line. That shows both theory and practical sense — exactly what business and engineering teams want.