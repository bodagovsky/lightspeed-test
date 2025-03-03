# IPv4 Unique Counter

A high-performance and memory-efficient tool for counting unique IPv4 addresses in a file using a bitmap approach.

## 🚀 Features

- **Ultra-Efficient Memory Usage**: Utilizes a bitmap to store IP addresses as encoded `uint32` integers, reducing memory requirements significantly.
- **Blazing Fast Performance**: Optimized for speed, outperforming traditional hash-based approaches.
- **Minimal Dependencies**: Pure Go implementation, requiring no external libraries.
- **Scalable**: Handles large datasets efficiently without excessive memory overhead.

## 🆚 Performance Comparison

| Implementation | Memory Usage |
| -------------- | ------------ |
| Hash Map       | \~16GB       |
| Bitmap         | \~0.5GB      |

The bitmap implementation drastically reduces memory consumption while maintaining excellent performance.

## 🛠 Installation & Usage

### Prerequisites

- Go installed (version 1.18 or later recommended)

### Running the Program

To execute the program, use the following command:

```sh
 go run main.go [filename]
```

Replace `[filename]` with the path to the file containing IPv4 addresses (one per line).

## 📜 License

MIT License. Feel free to modify and distribute!

