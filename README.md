# Numetric Flow

Hello GoSolve! This is my solution for your technical interview, hope you will enjoy it! :)

I created the application using technologies:
- Go (Gin)
- Next.js
- Tailwind CSS
- Docker

## Numbers, Index and Searching

I created `generate_numbers.py` script for generating numbers in `numbers.txt` - the script can be configured.

Numbers are loaded into slice in golang during backend startup.

Searching mechanism is implemented using BinarySearch with tolerance (%) parameter - it will return found number or closest number with tolerance.

# Manual test

I crafted `Makefile` to automate test in golang, to run testing unvoke

```bash
$ make test
```

![make](https://i.postimg.cc/wvBLwSbt/make.png)

## Frontend Context

Frontend is written using next.js and tailwind css, it has numbers context to share numbers state and interface between component.

## Gen AI and logo

I generated a logo for the application, because I believe AI is perfect for accelerating our work as developers! ;]

<img src="https://i.postimg.cc/SQgm9j97/logo.webp" width="254" height="254" /> 

## Preview

This is sample image how the app looks like on the frontside side after response from backend

![img](https://i.postimg.cc/YCyXNzSz/Screenshot-2024-02-19-at-10-54-26-Numetric-Flow.png)

## How to run?

Instruction how to run the app using `Manual` or `Docker`

Before, invoke python script for numbers generating.

```bash
$ python generate_numbers.py
```

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
