const COLORS = {
  black: 0, brown: 1, red: 2, 
  orange: 3, yellow: 4, green: 5,
  blue: 6, violet: 7, grey: 8, white: 9
} as const;

type Color = keyof typeof COLORS;

export function decodedValue(colors: Color[]): number {
  const [first, second] = colors;
  return(COLORS[first] * 10 + COLORS[second]);
}
