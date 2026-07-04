export function format(name: string, number: number) {
  let suffix = "th";

  const lastTwo = number % 100;
  const last = number % 10;

  if (lastTwo < 11 || lastTwo > 13) {
    if (last === 1) suffix = "st";
    else if (last === 2) suffix = "nd";
    else if (last === 3) suffix = "rd";
  }

  return `${name}, you are the ${number}${suffix} customer we serve today. Thank you!`;
}
