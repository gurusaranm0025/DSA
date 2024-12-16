export default function TwoCrystalBalls(breaks: boolean[]): number {
    const jmpAmt = Math.floor(Math.sqrt(breaks.length))
    let i = jmpAmt

    for (; i < breaks.length; i += jmpAmt) {
        if (breaks[i]) {
            break
        }
    }

    i -= jmpAmt

    for (let j = 0; j < jmpAmt && i < breaks.length; j++, i++) {
        if (breaks[i]) {
            return i
        }
    }

    return -1
}

console.log(TwoCrystalBalls([false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true]))