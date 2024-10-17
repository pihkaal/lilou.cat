import { resolve } from "path";
import { readdir, readFile } from "fs/promises";

export default defineEventHandler(async () => {
  const images = await readdir(resolve("public", "lilou"));
  const imageName = images[Math.floor(Math.random() * images.length)];
  const image = await readFile(resolve("public", "lilou", imageName));

  return image;
});
