# Dumdum

A mindless binary that can be controlled and examined through Unix Signals. It
is intended for use during integration testing on \*nix systems.

## Usage

* Manual
* Upstart
* Initd
* Systemd
* Launchd
* Others?

## Control Signals

* `HUP` - Reload configuration
* `WINCH` - Restart all workers, clearing memory bloat and simulated CPU
* `TTIN` - Increase the number of spawned workers by one
* `TTOU` - Descrease the number of spawned workers by one
* `USR1` - Bloat the process RSS by a fixed amount, specified in the configuration
* `USR2` - Simulate a fixed amount of CPU load
