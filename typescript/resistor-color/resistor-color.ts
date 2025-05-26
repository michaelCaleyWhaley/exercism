export const COLORS = [
  "black",
  "brown",
  "red",
  "orange",
  "yellow",
  "green",
  "blue",
  "violet",
  "grey",
  "white",
] as const;

type ColorCode = (typeof COLORS)[number];

export const colorCode = (color: ColorCode): number => {
  return COLORS.indexOf(color);
};
