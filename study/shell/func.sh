#!/bin/bash
hello() {
  echo "Hello $1"
}
today() {
  echo -n "Today's date is: "
  date +"%A, %B %-d, %Y"
}

fn () {
  foo=1
  echo "fn: foo = $foo"
}

hello
today
fn