export default function BubbleSort(array: number[]): number[] {
    const length = array.length

    for (let i = 0; i < length; i++) {
        for (let j = 0; j < length - 1 - i; j++) {
            if (array[j] > array[j + 1]) {
                let temp = array[j + 1]
                array[j + 1] = array[j]
                array[j] = temp
            }
        }
    }

    return array
}

console.log(BubbleSort([12, 1, 21, 5 + 6, 78, 89, 54, 56, 5, 3]))