<template>
  <main
    class="h-full items-center relative p-5 justify-center flex flex-col space-y-4 bg-[#fa86c4] z-10"
  >
    <NuxtRouteAnnouncer />
    <h1
      class="text-4xl md:text-5xl lg:text-6xl font-serif text-white dark:text-white"
    >
      ✨ Lilou ✨
    </h1>
    <img
      class="w-full max-w-[50vh] rotate-90"
      src="https://lilou.cat/api"
      alt="Lilou"
    />
    <h1
      class="text-4xl md:text-5xl lg:text-6xl font-serif text-white dark:text-white"
    >
      ✨ Lilou ✨
    </h1>

    <div class="absolute inset-0">
      <StarshineGlitter />
    </div>
  </main>
</template>
