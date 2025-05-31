<script setup lang="ts">
import type { UserResponse } from "~/types";

const { logout } = useSanctumAuth();
const user = useSanctumUser<UserResponse>();
</script>

<template>
  <div>
    <header class="bg-neutral-700 text-white h-14">
      <div class="container mx-auto px-4 flex flex-row justify-between gap-x-4">
        header

        <div>Hello, {{ user?.name }}</div>
        <UButton @click.prevent="logout">Logout</UButton>
      </div>
    </header>
    <main class="container mx-auto px-6 py-8">
      <slot />
    </main>
  </div>
</template>
