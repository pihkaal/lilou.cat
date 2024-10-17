<script setup lang="ts">
import { resolve } from "path";
import { readdir } from "fs/promises";

const images = await readdir(resolve("public", "lilou"));
const image = images[Math.floor(Math.random() * images.length)];
const imagePath = `/lilou/${image}`;
</script>

<template>
  <img class="w-full max-w-[50vh]" :src="imagePath" alt="Lilou" >
</template>
