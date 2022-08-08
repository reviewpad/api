# Reviewpad Golang API

## Build

We use Taskfile. To build, you need to install the dependencies. 

For osx, you can simply run:

```
task install-deps
```

For other operating systems, you just need to adapt the `PROTOC_ZIP` variable in the script [installdep.sh](./installdep.sh).

To build, run:

```
task build
```

This will generate and compile the sources in [entities](./entities) and [services](./services) from the protobuf specifications in the [pb](../pb/) directory.

## License

Reviewpad is available under the GNU Lesser General Public License v3.0 license.

See LICENSE for the full license text.