# Snap

Snap allows you to easily take pictures from the Raspberry Pi through the internet. Snap is also extensible with protocols such as Unix Domain Sockets.

Snap is experimental. Your [feedback](https://github.com/AudreyLim/snap/issues) is very much appreciated. 

---------------------------------------

* [Installation](#installation)
  * [Building the Compiler](#building-the-compiler)
  * [Cross-compiling Snap](#cross-compiling-snap)
* [Usage](#usage)
  * [HTTP](#http)
  * [Extend](#extend)
* [Tests](#tests)
* [License](#license)

---------------------------------------

## Installation

Note: For setting up the Pi, refer to this [guide](http://www.raspberrypi.org/help/quick-start-guide/).

### Building the Compiler

To run Snap on the Pi, ensure that you have enabled cross-compilation for Go. 

With Mac OS/Homebrew, install Go with this command

```
brew install go --cross-compile-common or --cross-compile-all
```

If you have installed Go from source, set the environment variables and call ./all.bash in the src folder for each GOARCH/GOOS combination you need. For a step-by-step guide, refer to [this](http://dave.cheney.net/2013/07/09/an-introduction-to-cross-compilation-with-go-1-1).

### Cross-compiling Snap

In your terminal, run the following command to compile `snap.go`:

```
GOARCH=arm GOOS=linux GOARM=5 go build snap.go
```

Next, copy the binary to the SD card for your Pi. You can do this by directly copying `snap` from your desktop/Mac to the SD card. If you are running the Pi [headless](https://www.andrewmunsell.com/blog/setting-up-raspberry-pi-as-headless-device), run this command:

```
scp snap pi@[IPAddressRaspberryPi]:/home/pi
```

## Usage

Once you have `snap` on your SD card, run `snap` to allow it to listen on sockets.

```
./snap
```

### HTTP

With `snap` running on the Pi, you can command the Pi to take a picture by calling the following URL from your browser:

```
http://[IPAddressRaspberryPi]:3000/snap?flip=[value]
```

The flip param is optional. The value can be `hf` for the picture to flip horizontally or `vf` to flip vertically.

Snap will stream the captured image so you can view it from your browser.

### Extend

You can also use `snap` in your own Pi projects by calling `snap` with Unix Domain Sockets. 

For an example implementation, see `unixclient.go`.

### Tests

To run tests on snap.go, compile raspistill-mock.go in the helpers folder:

```
go build raspistill-mock.go
```

Set the PATH_TO_RASPISTILL-MOCK environment variable to the path to `raspistill-mock`, eg. `./helpers/raspistill-mock`.

Run tests in the root folder:

```
go test
```

## License

MIT License - see the LICENSE.txt file in the source distribution.



