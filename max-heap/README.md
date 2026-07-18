# MaxHeap

## Pseudocode

### MaxHeapify

```text
algorithm MaxHeapify(B, s):
    // INPUT
    //    B = input array
    //    s = an index of the node
    // OUTPUT
    //    The heap tree that obeys max-heap property

    left <- 2 * s
    right <- 2 * s + 1

    if left <= B.length and B[left] > B[s]:
        largest <- left
    else:
        largest <- s

    if right <= B.length and B[right] > B[largest]:
        largest <- right

    if largest != s:
        swap(B[s], B[largest])
        MaxHeapify(B, largest)
```

### Bubble Up

```text
algorithm BubbleUp(B, i)
    // INPUT
    //    B = input array
    //    i = an index of the node
    // OUTPUT
    //    The heap tree that obeys max-heap property

    while i > 0:
        parent <- (i - 1) / 2

        if B[i] <= B[parent]
            break

        swap(B[i], B[parent])

        i = parent
```
