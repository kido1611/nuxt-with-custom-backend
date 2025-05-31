<script setup lang="ts">
definePageMeta({
  layout: "dashboard",
  middleware: ["sanctum:auth"],
});
</script>

<template>
  <div>
    <h1>Dashboard</h1>
  </div>
</template>
