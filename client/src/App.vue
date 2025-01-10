<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";

type RouteLink = { name: string; path: string };

const routes: RouteLink[] = [
  { name: "Home", path: "/" },
  { name: "About", path: "/about" },
];
</script>

<template>
  <header class="flex items-center justify-between bg-gray-800 p-4 text-white">
    <img alt="Vue logo" class="w-10" src="@/assets/logo.svg" />
    <nav class="space-x-4">
      <RouterLink v-for="route in routes" :to="route.path">{{ route.name }}</RouterLink>
    </nav>
  </header>

  <RouterView />
</template>
