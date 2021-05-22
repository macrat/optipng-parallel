optipng-parallel
================

Run [OptiPNG](http://optipng.sourceforge.net/) parallelly.

## Install

1. Install [OptiPNG](http://optipng.sourceforge.net/).

2. Build optipng-parallel

``` shell
$ go install github.com/macrat/optipng-parallel
```

## Usage

``` shell
$ ls **/*.png | optipng-parallel
```

``` shell
$ ls **/*.png | optipng-parallel -o7
```
