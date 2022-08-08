# Reviewpad API

## Build

To build, you need to install the dependencies. 

For osx, you can simply run:

```
./installdep-osx.sh
```

For other operating systems, you just need to adapt the `PROTOC_ZIP` variable in the script.

Then, you should be able to generate the Go code in `go` from the proto files in `pb` by running:

```
./genproto.sh
```

## License

Reviewpad is available under the GNU Lesser General Public License v3.0 license.

See LICENSE for the full license text.