# machina
高度可定制状态机

# Run example
    cd example
    go run main.go
```
output:

1. UnitFsm.handle(unit, 'tick', 100, 5);
transition: idle undefined => idle
enter idle1
enter idle2
idle tick1: idle 100 5
exit idle
transition: cd idle => cd
enter cd
idle tick2: cd 100 5
2. UnitFsm.handle(unit, 'confirm', 100, 5);
cd confirm1: cd 100
exit cd
transition: ready cd => ready
ready enter1: ready
exit ready
transition: idle ready => idle
enter idle1
enter idle2
idle confirm: idle 100
ready enter2: idle
cd confirm2:, idle 100
```