# metrics with prometheus go client

## HELP go_gc_duration_seconds A summary of the GC invocation durations.
## TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 0
go_gc_duration_seconds{quantile="0.25"} 0
go_gc_duration_seconds{quantile="0.5"} 0
go_gc_duration_seconds{quantile="0.75"} 0
go_gc_duration_seconds{quantile="1"} 0
go_gc_duration_seconds_sum 0
go_gc_duration_seconds_count 0
## HELP go_goroutines Number of goroutines that currently exist.
## TYPE go_goroutines gauge
go_goroutines 6
## HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
## TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 490360
## HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
## TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 490360
## HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
## TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 2748
## HELP go_memstats_frees_total Total number of frees.
## TYPE go_memstats_frees_total counter
go_memstats_frees_total 129
## HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
## TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 137216
## HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
## TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 490360
## HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
## TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 671744
## HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
## TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 1.130496e+06
## HELP go_memstats_heap_objects Number of allocated objects.
## TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 5337
## HELP go_memstats_heap_released_bytes_total Total number of heap bytes released to OS.
## TYPE go_memstats_heap_released_bytes_total counter
go_memstats_heap_released_bytes_total 0
## HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
## TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 1.80224e+06
## HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
## TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 0
## HELP go_memstats_lookups_total Total number of pointer lookups.
## TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 11
## HELP go_memstats_mallocs_total Total number of mallocs.
## TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 5466
## HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
## TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 3472
## HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
## TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 16384
## HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
## TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 18696
## HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
## TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 32768
## HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
## TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 4.473924e+06
## HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
## TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 798020
## HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
## TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 294912
## HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
## TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 294912
## HELP go_memstats_sys_bytes Number of bytes obtained by system. Sum of all system allocations.
## TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 3.084288e+06
## HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
## TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 0
## HELP process_max_fds Maximum number of open file descriptors.
## TYPE process_max_fds gauge
process_max_fds 1.048576e+06
## HELP process_open_fds Number of open file descriptors.
## TYPE process_open_fds gauge
process_open_fds 7
## HELP process_resident_memory_bytes Resident memory size in bytes.
## TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 5.955584e+06
## HELP process_start_time_seconds Start time of the process since unix epoch in seconds.
## TYPE process_start_time_seconds gauge
process_start_time_seconds 1.52435023503e+09
## HELP process_virtual_memory_bytes Virtual memory size in bytes.
## TYPE process_virtual_memory_bytes gauge
process_virtual_memory_bytes 4.8562176e+07
