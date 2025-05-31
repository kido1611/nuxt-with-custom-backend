<script setup lang="ts">
const { logout } = useSanctumAuth();
</script>

<template>
  <div>
    <header class="bg-neutral-700 text-white h-14">
      <div>
        header

        <UButton @click.prevent="logout">Logout</UButton>
      </div>
    </header>
    <main>
      <slot />
    </main>
  </div>
</template>
