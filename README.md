# chuck
An Internet Chuck Norris DB client in Go

## Install
Just download the binary for your platform from the releases page, rename it to chuck, set executable permission if necessary, and you're all set!

## Usage
```bash
chuck [-n]
  -n Only nerdy Chuck Norris jokes.
```
On Windows, add the .exe extension.

## Disclaimer
This client only retrieves a single random joke, optionally filtering to only nerdy
jokes. The full icndb.com API allows other interesting options, but I kept it simple
for the moment.
