# Performance analysis

prepare
-------
```import "https://github.com/pkg/profile"```

```
func main() {
  defer profile.Start().Stop()
  ...
}
```

runing
------
output with cpu profiling
```
> go build && day15
2017/12/16 13:57:08 profile: cpu profiling enabled, /var/folders/q8/6zrc0l454w7c14fw_jrg7b100000gn/T/profile003339230/cpu.pprof
The judge's final count is '592'
The judge's final count is '320'
2017/12/16 13:57:11 profile: cpu profiling disabled, /var/folders/q8/6zrc0l454w7c14fw_jrg7b100000gn/T/profile003339230/cpu.pprof
```

start go profiler

```go tool pprof day15 /var/folders/q8/6zrc0l454w7c14fw_jrg7b100000gn/T/profile003339230/cpu.pprof ```

```
(pprof) top
Showing nodes accounting for 1940ms, 100% of 1940ms total
      flat  flat%   sum%        cum   cum%
    1800ms 92.78% 92.78%     1800ms 92.78%  main.generate.func1 /.../advent-of-code-2017/day15/main.go
     100ms  5.15% 97.94%     1900ms 97.94%  main.judge /Users/oleg/go/src/github.com/olbura/advent-of-code-2017/day15/main.go
      40ms  2.06%   100%       40ms  2.06%  runtime.usleep /usr/local/Cellar/go/1.9.2/libexec/src/runtime/sys_darwin_amd64.s
         0     0%   100%     1900ms 97.94%  main.main /.../advent-of-code-2017/day15/main.go
         0     0%   100%     1900ms 97.94%  runtime.main /usr/local/Cellar/go/1.9.2/libexec/src/runtime/proc.go
         0     0%   100%       40ms  2.06%  runtime.mstart /usr/local/Cellar/go/1.9.2/libexec/src/runtime/proc.go
         0     0%   100%       40ms  2.06%  runtime.mstart1 /usr/local/Cellar/go/1.9.2/libexec/src/runtime/proc.go
         0     0%   100%       40ms  2.06%  runtime.sysmon /usr/local/Cellar/go/1.9.2/libexec/src/runtime/proc.go
```

list closure of generate func

```
(pprof) list func1
Total: 1.94s
ROUTINE ======================== main.generate.func1 in /.../advent-of-code-2017/day15/main.go
     1.80s      1.80s (flat, cum) 92.78% of Total
         .          .     21:	faktor int
         .          .     22:	mult   int
         .          .     23:}
         .          .     24:
         .          .     25:func generate(g generator) func() int {
      30ms       30ms     26:	return func() int {
         .          .     27:		for {
     240ms      240ms     28:			g.value = (g.value * g.faktor) % 2147483647
     1.44s      1.44s     29:			if g.value%g.mult == 0 {
      90ms       90ms     30:				return g.value & 0xffff
         .          .     31:			}
         .          .     32:		}
         .          .     33:	}
         .          .     34:}
         .          .     35:
```

conclusion
----------
the [```remainder```](https://golang.org/ref/spec#Arithmetic_operators) operation takes the most time for calculation
