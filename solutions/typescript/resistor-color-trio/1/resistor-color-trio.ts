const COLORS = {
  black: 0, brown: 1, red: 2, 
  orange: 3, yellow: 4, green: 5,
  blue: 6, violet: 7, grey: 8, white: 9
} as const;

type Color = keyof typeof COLORS;

export function decodedResistorValue(colors: Color[]) {
  const [first, second, third] = colors;
  
  const base: number = COLORS[first] * 10 + COLORS[second];
  const powered: number = base * (10 ** COLORS[third]);
  
  let word: string = "";
  let finalValue = powered;
  
  if (powered >= 1_000_000_000) {
    word = "giga";
    finalValue = powered / 1_000_000_000;
  } else if (powered >= 1_000_000) {
    word = "mega";
    finalValue = powered / 1_000_000;
  } else if (powered >= 1_000) {
    word = "kilo";
    finalValue = powered / 1_000;
  };
  return (`${finalValue} ${word}ohms`)
}
