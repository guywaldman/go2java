# go2java

Minimal example of calling a Go function from Java using JNI (Java Native Interface).
An easier alternative is using JNA, however it is generally considered much slower compared to JNI.

## Prerequisites

- Go (tested with Go 1.23)
- Java (`brew install java`)
  - Make sure `JAVA_HOME` is defined (for Homebrew users, this is `$HOMEBREW_PREFIX/opt/openjdk`)
- `gcc` (`clang` may work too, have not tried)

## Overview

- [./internal/greeting/greeting.go]() contains the business logic in Go that we would like to expose to Java (or the JVM). It has a simple `Greet` function that takes a name and returns a greeting string (no dependency on `cgo`).
- [./internal/jni_export/main.go]() is the Go code that exposes the `Greet` function to Java via JNI. It uses the `C` package to create the necessary JNI bindings. Note the `cgo` comments, that include required headers as well as define helper functions. Other examples that you see online may require separate compilation of "helper libraries", but this is a much simpler method.

## Running the example

```shell
make run
```

