## for-select pattern

`for-select`在Go中到处可见。

这种 pattern 有下面两种形式是一样的，用哪种取决于个人喜好:

```Go
for {
    select {
    case <-done:
        return
    default:
    }
    // Do non-preemptable work    
}
```

```Go
for {
    select {
        case <- done:
            return
        default:
        // Do non-preemptable work    
    }
}

```