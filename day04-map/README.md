# 📘 Go Learning Journey – Day 4  
## Maps, Multiple Greetings & Backward Compatibility  

---

## 📅 Overview  

On **Day 4** of my Go learning journey, I explored working with multiple people dynamically instead of handling just one name at a time.

In this session, I learned about:

- Using **maps** as a key-value data structure in Go  
- Looping through slices using `range`  
- Returning both a **map and an error** from a function  
- Properly propagating errors  
- Implementing **backward compatibility** by adding new functionality without breaking existing code  

📖 Reference:  
https://go.dev/doc/tutorial/greetings-multiple-people  

---

## 🗂 Maps in Go (Key-Value Data Structure)

A **map** in Go is a built-in reference type used to associate a **key** with a **value**.

It allows:
- Fast lookups  
- Structured data storage  
- Easy association between related data  

### 🔹 Syntax

```go
map[KeyType]ValueType
