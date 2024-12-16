export default function LinearSearch(haystack: number[], needle: number): boolean {
    for (let i = 0; i < haystack.length; i++) {
        if (haystack[i] === needle) {
            return true;
        }
    }

    return false;
}

console.log(LinearSearch([5, 4, 1, 0, 6], 0))