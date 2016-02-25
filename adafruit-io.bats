#!/usr/bin/env bats

@test "adafruit-io" {
	run ./adafruit-io
	[ "$status" -eq 0 ]
}

@test "adafruit-io send" {
	run ./adafruit-io send foo 56
	[ "$output" = "OK" ]
	[ "$status" -eq 0 ]
}

@test "adafruit-io get" {
	run ./adafruit-io get foo
	[ "$status" -eq 0 ]
	[ "$output" -eq "56" ]
}