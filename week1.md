Below is a **7-day Go project sprint**, designed for someone who:

* Knows Go syntax
* Can write small programs
* Needs **practice, structure, and confidence**
* Wants **progressive difficulty**, not repetition

No hand-holding. Each day builds real skill.

---

## 7-Day Go Project Sprint (Progressive)

### **Day 1 — In-Memory CRUD Manager**

**Goal:** Fluency with core Go data structures

**Project:** Student Manager
**Requirements:**

* `Student` struct (ID, Name, Score)
* Store students in a `map[int]Student`
* Functions:

  * Create student
  * Get student by ID
  * Update student
  * Delete student
* Use **methods** on a struct (not free functions)
* Use **pointers** for updates
* Print results to console

**What This Trains:**

* Structs, maps, methods, pointers
* Thinking in data models

---

### **Day 2 — Interface-Based Service**

**Goal:** Abstraction and clean design

**Project:** Storage Service
**Requirements:**

* Define an interface:

  ```go
  type StudentStore interface {
      Create(Student)
      Get(int) (Student, error)
      Update(Student) error
      Delete(int) error
  }
  ```
* Implement:

  * `InMemoryStudentStore`
* Refactor Day 1 code to use the interface
* Add proper error handling

**What This Trains:**

* Interfaces
* Decoupling logic
* Idiomatic Go design

---

### **Day 3 — HTTP API (No Database)**

**Goal:** Real Go usage

**Project:** Student REST API
**Requirements:**

* Use `net/http`
* Endpoints:

  * `POST /students`
  * `GET /students/{id}`
  * `PUT /students/{id}`
  * `DELETE /students/{id}`
* JSON input/output
* Use Day 2 store interface

**What This Trains:**

* HTTP servers
* JSON encoding/decoding
* Request handling

---

### **Day 4 — Concurrency & Safety**

**Goal:** Controlled concurrency

**Project:** Concurrent Student Access
**Requirements:**

* Protect your map with `sync.RWMutex`
* Add concurrent requests support
* Simulate:

  * Multiple goroutines reading/writing
* Ensure no data races

**What This Trains:**

* Goroutines
* Mutexes
* Thread-safe design

---

### **Day 5 — Background Worker**

**Goal:** System-style programs

**Project:** Grade Processor
**Requirements:**

* Channel-based job queue
* Worker goroutine that:

  * Receives student scores
  * Calculates grade
  * Stores result
* Periodic job using `time.Ticker`

**What This Trains:**

* Channels
* Goroutines
* Background processing

---

### **Day 6 — File I/O & Persistence**

**Goal:** Real-world data handling

**Project:** Persistent Student Store
**Requirements:**

* Save students to a JSON file
* Load on program startup
* Replace in-memory store with file-backed store
* Handle file errors properly

**What This Trains:**

* File I/O
* Encoding/decoding
* Program lifecycle thinking

---

### **Day 7 — Clean Architecture & Review**

**Goal:** Professional structure

**Project:** Refactor Everything
**Requirements:**

* Proper folder structure:

  ```
  cmd/
  internal/
    store/
    service/
    handler/
  ```
* Separate concerns cleanly
* Add comments explaining decisions
* Write a short README

**What This Trains:**

* Real-world Go architecture
* Maintainability
* Code clarity

---

## Rules for the 7 Days

1. **No frameworks**
2. **No copy-paste**
3. **Write from scratch**
4. **Finish each project**
5. **Commit daily**

---

## How This Transforms You

After these 7 days, you will:

* Think in **Go patterns**
* Be comfortable with services
* Understand why Go is used in cloud systems
* Be ready for databases, Docker, and Kubernetes

---


