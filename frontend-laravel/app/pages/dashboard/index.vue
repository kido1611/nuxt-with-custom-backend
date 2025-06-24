<script setup lang="ts">
definePageMeta({
  layout: "dashboard",
  middleware: ["sanctum:auth"],
});
</script>

<template>
  <div>
    <h1>Dashboard</h1>

    <div>
      <NuxtLink to="/dashboard/notes" class="font-semibold px-4 py-3"
        >See Notes</NuxtLink
      >
    </div>
  </div>
</template>
