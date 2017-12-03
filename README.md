# doni

By [ddspog](https://github.com/ddspog)

**doni** helps you setup experiments and collect metrics.

## License

You are free to copy, modify and distribute **cvinfo** with attribution under the terms of the MIT license. See the [LICENSE](https://github.com/ddspog/doni/blob/master/LICENSE) file for details.

## Installation

Install **doni** with:

```shell
go get -u github.com/ddspog/doni/cmd/doni
```

## How to use

Call `doni --help` and see more details.

## Contribution

This package has some objectives from now:

* Create start command, that will start the experiment configured on a experiment.yaml file. The files used on experiment need to be on same folder as configuration file:

```shell
doni start <path\to\Experiment.yaml>
```

* Create collect command, that will download the result of a experiment configured in some address.

```shell
doni collect <host:port>
```

Any interest in help is much appreciated.