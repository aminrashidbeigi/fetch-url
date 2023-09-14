# Fetch Web

Package fetchurl provides a command-line program for fetching web pages and saving them to disk for later retrieval and browsing.

## Introduction

This command-line utility allows users to download web pages and store them as HTML files on their local disk. Additionally, it records metadata about the fetched web pages, including the date and time of the last fetch, the number of links on the page, and the number of images on the page.

## Usage

To use this utility, run the following command:

    ./fetchurl fetch [options] url [url2 ...]

## Options

- `-metadata` or `-m`: Print metadata information for each fetched web page.

## Examples

1. Fetch a single web page and save it to a file:

    `./fetchurl fetch https://www.example.com`

2. Fetch multiple web pages and save them to separate files:

    `./fetchurl fetch https://www.example.com https://www.anotherexample.com`

3. Fetch a web page and display metadata information:

    `./fetchurl fetch --metadata https://www.example.com`

4. Fetch multiple web pages and display metadata for each:

    `./fetchurl fetch --metadata https://www.example.com https://www.anotherexample.com`

## Run with Docker

```
docker build -t image .
docker run image ./fetchurl fetch --metadata https://www.example.com
```
