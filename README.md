haggar [![Build Status](https://travis-ci.org/gorsuch/haggar.svg?branch=master)](https://travis-ci.org/gorsuch/haggar)
======

An experimental [carbon](https://github.com/graphite-project/carbon) load generation tool named after [Haggar the Witch](http://www.cheezey.org/voltron/haggar.htm), a real menace to [Voltron](http://www.voltron.com/).

![](http://f.cl.ly/items/050Y473L1x0j1y1s0744/Image%202014-08-07%20at%2015.08.35.png)

It's also named after pants.

![](http://slimages.macys.com/is/image/MCY/products/8/optimized/1096328_fpx.tif?01AD=3T7DyZyp_siLqj1q-neozCxIommQ92M1GsNc5fe_xTNqBcjyGG2gMxA&01RI=C624DC2009B77F9&01NA=&$filterlrg$&wid=370)


It behaves like a swarm of [collectd](https://collectd.org/) agents firing a fixed number of metrics at a fixed interval to your [carbon-compatible](https://github.com/graphite-project/carbon) endpoint.  The number of agents increase over time until a maximum number is reached.  At any given time, you can pause the spawning of new agents by sending `SIGUSR1`.  Spawning can be resumed by doing the same.

## Installation

```sh
$ go get github.com/shiimaxx/haggar
```

## Command-line flags

```sh
$ haggar -h
Usage of haggar:
  -cache_conections: if set, keep connections open between flushes
  -carbon="localhost:2003": address of carbon host
  -datapoints=1: number of datapoints each metrics
  -metrics=10000: number of metrics for each agent to hold
  -prefix="haggar": prefix for metrics
  -tasks=100: number of tasks that will pass to woker
  -workers=100: max number of workers to run concurrently
```

## Credits

This tool was designed and developed by [@gorsuch](https://github.com/gorsuch) and [@obfuscurity](https://github.com/obfuscurity).
