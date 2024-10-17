<script setup lang="ts">
import { resolve } from "path";
import { readdir } from "fs/promises";

const images = await readdir(resolve("public", "lilou"));
const image = images[Math.floor(Math.random() * images.length)];
const imagePath = `/lilou/${image}`;
</script>

<template>
  <main
    class="h-[100vh] items-center justify-center p-5 flex flex-col space-y-4 bg-[#fa86c4]"
  >
    <NuxtRouteAnnouncer />

    <h1
      class="text-4xl md:text-5xl lg:text-6xl font-serif text-white dark:text-white"
    >
      ✨ Lilou ✨
    </h1>
    <img class="w-full max-w-[50vh]" :src="imagePath" alt="Lilou" />
    <h1
      class="text-4xl md:text-5xl lg:text-6xl font-serif text-white dark:text-white"
    >
      ✨ Lilou ✨
    </h1>
  </main>
</template>
