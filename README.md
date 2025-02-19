# Custom HashMap in Golang

This is a toy project to understand internal implementation details of hashmap. I tried to build a custom implementation of a hashmap in Golang. The implementation uses integer key-value pairs for simplicity and employs the **chaining method** to resolve collisions.

> **Warning**
> README file is LLM generated (code is not), sorry for any errors.

## Features

- **Custom HashMap Implementation**: Built from scratch, demonstrating core hashmap functionality.
- **Collision Handling**: Uses the chaining method to manage hash collisions.
- **Benchmarks**: Performance comparison against Golang's inbuilt `map`.

## Benchmarks

Below are the benchmark results comparing operations (`set`, `get`, and `remove`) of the custom implementation (`mymap`) and Golang's inbuilt `map`:


| Operation  | Inbuilt Map (ns/op) | MyMap (Chaining) (ns/op) | MyMap (Open Addressing) (ns/op) |
|------------|---------------------|--------------------------|----------------------------------|
| **Set**    | 94.16              | 137.6                   | 139.2                           |
| **Get**    | 8.61               | 44.87                   | 1894.87                         |
| **Remove** | 54.66              | 42.23                   | 136.7                           |


### System Configuration
- **OS**: macOS (Darwin)
- **Architecture**: ARM64
- **CPU**: Apple M3 (8 cores used)

### Observations
- **Set and Remove**: Custom implementation performs comparably to the inbuilt map.
- **Get**: The custom implementation is ~5x slower than the inbuilt map, likely due to the lack of spatial locality. Open addressing might improve performance for this operation. 

Update: It did not, open addressing made it worse maybe due to clustering of collissions. I am curious about quadratic probing function.

Check golang runtime code for map and it uses something called swiss tables which does some bit manipulation magic to check for 8 or 16 chunks at once. https://abseil.io/about/design/swisstables

## How to Run

Clone the repo and run benchmark using 
```bash
go test -bench=.
```