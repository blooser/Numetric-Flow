# Numetric Flow

Hello GoSolve! This is my solution for your technical interview, hope you will enjoy it! :)

I created the application using technologies:
- Go (Gin)
- Next.js
- Tailwind CSS
- Docker

## Numbers, Index and Searching

I created `generate_numbers.py` script for generating numbers in `numbers.txt` - the script can be configured.

Numbers are loaded indo slice in golang during backend startup.

Searching mechanism is implemented using BinarySearch with tolerance (%) parameter - it will return found number or closest number with tolerance.

## Preview

![img](https://i.postimg.cc/YCyXNzSz/Screenshot-2024-02-19-at-10-54-26-Numetric-Flow.png)

## How to run?

Instruction how to run the app using `Manual` or `Docker`

## Manual

To run backend, invoke

```bash
$ go run .
```

To run frontend, invoke

```bash
$ npm run dev
```

## Docker

Docker is easy, just execute

```bash
$ docker-compose up -d --build
```

## Author

Mateusz Solnica (blooser@protonmail.com)
