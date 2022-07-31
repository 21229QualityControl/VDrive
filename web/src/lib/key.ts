const binds: Record<string, number> = {
  "UP": 38,
  "DOWN": 40,
  "LEFT": 37,
  "RIGHT": 39,
  "Q": 81,
  "W": 87,
  "E": 69,
  "R": 82,
  "T": 84,
  "Y": 89,
  "U": 85,
  "I": 73,
  "O": 79,
  "P": 80,
  "A": 65,
  "S": 83,
  "D": 68,
  "F": 70,
  "G": 71,
  "H": 72,
  "J": 74,
  "K": 75,
  "L": 76,
  "Z": 90,
  "X": 88,
  "C": 67,
  "V": 86,
  "B": 66,
  "N": 78,
  "M": 77,
  "COMMA": 188,
  "PERIOD": 190,
  "SLASH": 191,
  "SEMICOLON": 186,
  "APOSTROPHE": 222,
  "LBRACKET": 219,
  "RBRACKET": 221,
  "MINUS": 189,
  "EQUAL": 187,
  "ONE": 49,
  "TWO": 50,
  "THREE": 51,
  "FOUR": 52,
  "FIVE": 53,
  "SIX": 54,
  "SEVEN": 55,
  "EIGHT": 56,
  "NINE": 57,
  "ZERO": 48,
}

function reverseBinds(): Record<number, string> {
  let out: Record<number, string> = {};
  for (let k of Object.keys(binds)) {
    out[binds[k]] = k;
  }
  return out;
}

const bindsRev = reverseBinds();

export function getKeyCode(key: string): number {
  return binds[key];
}

export function getKeyForCode(code: number): string {
  return bindsRev[code];
}