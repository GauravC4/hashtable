# Custom HashMap in Golang

This is a toy project to understand internal implementation details of hashmap. I tried to build a custom implementation of a hashmap in Golang. The implementation uses integer key-value pairs for simplicity and employs the **chaining** and **open addressing** to resolve collisions.

> **Warning**
> README file is LLM generated (code is not).

## Features

- **Custom HashMap Implementation**: Built from scratch, demonstrating core hashmap functionality.
- **Collision Handling**: Uses the chaining and open addressing with linear probing to manage hash collisions.
- **Benchmarks**: Performance comparison against Golang's inbuilt `map`.

## Benchmarks

Below are the benchmark results comparing operations (`set`, `get`, and `remove`) of the custom implementation (`mymap`) and Golang's inbuilt `map` for 10k iterations:


| Operation | Inbuilt Benchmark | MyMap (Chaining) | MyMap (Open Addressing Linear Probing) |
|-----------|-------------------|------------------|----------------------------------------|
| Set       | 105.2 ns/op       | 116.7 ns/op      | 118.2 ns/op                            |
| Get       | 13.68 ns/op       | 15.11 ns/op      | 15.12 ns/op                            |
| Remove    | 76.90 ns/op       | 91.04 ns/op      | 72.28 ns/op                            |

### System Configuration
- **OS**: macOS (Darwin)
- **Architecture**: ARM64
- **CPU**: Apple M3 (8 cores used)

### Observations
Micro-benchmarks like these are miss leading but this was just for me to have a performance metric to compare against. It will be interesting to see how a quadratic probing function performs reducing collission clustering due to linear probing.

Checked golang runtime code for map and it uses something called swiss tables which does some bit manipulation magic to check for 8 or 16 chunks at once. https://abseil.io/about/design/swisstables

## How to Run

Clone the repo and run benchmark using 
```bash
go test -bench=.
```