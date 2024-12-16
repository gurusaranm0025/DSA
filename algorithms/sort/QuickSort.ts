const array = [9, 3, 7, 4, 69, 420, 42]

function qs(arr: number[], lo: number, hi: number): void {
    if (lo >= hi) {
        return
    }

    const pivot = partition(arr, lo, hi)

    qs(arr, lo, pivot - 1)
    qs(arr, pivot + 1, hi)
}

function partition(arr: number[], lo: number, hi: number): number {
    const pivot = arr[hi];

    let idx = lo - 1;

    for (let i = lo; i < hi; i++) {
        if (arr[i] <= pivot) {
            idx++;
            const temp = arr[i]
            arr[i] = arr[idx]
            arr[idx] = temp
        }
    }

    idx++;
    arr[hi] = arr[idx]
    arr[idx] = pivot

    return idx
}

export default function QuickSort(array: number[]): number[] {
    qs(array, 0, array.length - 1)
    return array
}

console.log(QuickSort(array))