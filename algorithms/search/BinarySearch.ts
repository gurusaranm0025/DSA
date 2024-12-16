export default function BinarySearch(haystack: number[], needle: number): number {
    let low = 0
    let high = haystack.length

    do {
        const mid = Math.floor(low + (high - low) / 2)

        if (haystack[mid] === needle) {
            return mid
        } else if (haystack[mid] < needle) {
            low = mid + 1
        } else {
            high = mid
        }

    } while (low < high)

    return -1
}

console.log(BinarySearch([1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40], 23))