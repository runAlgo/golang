# Fan-In Concurrency Pattern in Go

This is a **fan-in concurrency** example — one of the **most important and practical Go concurrency patterns**.

This explanation walks through the code **slowly, simply, and visually**, so you can understand *how it actually works at runtime*.

---

## 🔹 Big Idea

> **Fan-in = many channels → one channel**

```
ch1 ─┐
ch2 ─┼──► merged channel
ch3 ─┘
```

Instead of listening to many channels separately, we **merge them into one channel** and listen to just that.

---

## 🧩 Step-by-Step Explanation

---

## 1️⃣ `fanInInt` Function

```go
func fanInInt(channels ...chan int) chan int {
```

### What this function does

* Accepts **any number of `chan int`**
* Returns **one single output channel**

Think of it as a **mixer** that combines multiple inputs into one output.

---

### Create the output channel

```go
out := make(chan int)
```

This channel:

* Receives values from **all input channels**
* Is returned to `main()`

---

### Loop through all input channels

```go
for _, ch := range channels {
```

If you pass:

```go
ch1, ch2, ch3
```

The loop runs **3 times**, once for each channel.

---

### Start one goroutine per channel

```go
go func(c chan int) {
```

For **each channel**:

* A **separate goroutine** is started
* That goroutine listens **only to its own channel**

So internally:

* Goroutine #1 → listens to `ch1`
* Goroutine #2 → listens to `ch2`
* Goroutine #3 → listens to `ch3`

---

### Keep receiving values

```go
for v := range c {
	out <- v
}
```

This means:

1. Wait for a value on the input channel
2. When it arrives, send it to the `out` channel

Example flow:

```
ch2 sends 20
↓
goroutine receives 20
↓
20 is sent into merged channel
```

---

### Important Detail (Avoids a Common Bug)

```go
}(ch)
```

This ensures the **correct channel is passed** to the goroutine.
Without this, Go’s loop variable behavior could cause bugs.

---

### Return the merged channel

```go
return out
```

Now `main()` only needs to listen to **one channel**.

---

## 2️⃣ `main()` Function

---

### Create input channels

```go
ch1 := make(chan int)
ch2 := make(chan int)
ch3 := make(chan int)
```

These act as **independent producers**.

---

### Merge the channels

```go
merged := fanInInt(ch1, ch2, ch3)
```

Now:

```
merged receives from ch1 OR ch2 OR ch3
```

---

### Send values

```go
go func() {
	ch1 <- 10
	ch2 <- 20
	ch3 <- 30
}()
```

Sending happens inside a goroutine so `main()` does not block.

---

## 3️⃣ Receive from the merged channel

```go
for i := 0; i < 3; i++ {
	fmt.Println("Received:", <-merged)
}
```

* We expect **3 values**
* The order is **not guaranteed**

Possible output:

```
Received: 20
Received: 10
Received: 30
```

---

## 🧠 Why There Is No Deadlock

✔ Every send has a receiver
✔ Each input channel has its own goroutine
✔ The merged channel is actively read

Balanced sending and receiving = **safe execution**

---

## 🔑 Key Takeaways

### ✅ Why Fan-In Is Powerful

* Handles **N channels**
* Avoids large `select` blocks
* Clean and scalable
* Widely used in:

  * Worker pools
  * Pipelines
  * Servers
  * Event processing systems

---

### ⚠️ Important Note

This example does **not close channels**.

In real-world applications, you should:

* Close channels properly **or**
* Use `context.Context` for graceful shutdown

---

## 🧠 Mental Model

```
ch1 ─► goroutine ┐
ch2 ─► goroutine ├─► merged ─► main
ch3 ─► goroutine ┘
```

---