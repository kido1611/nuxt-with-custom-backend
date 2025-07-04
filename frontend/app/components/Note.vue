<script setup lang="ts">
import type { NoteResponse } from "~/types";
const { note } = defineProps<{
  note: NoteResponse;
}>();

const client = useNuxtApp().$apiFetch;
const toast = useToast();

const emit = defineEmits<{
  askReload: [];
}>();

async function deleteNote() {
  try {
    await client(`/api/api/notes/${note.id}`, {
      method: "DELETE",
    });
    toast.add({
      title: "Success",
      color: "success",
    });

    emit("askReload");
  } catch (error: any) {
    toast.add({
      title: "Error",
      description: error.data.message,
      color: "error",
    });
  }
}
</script>

<template>
  <UCard variant="soft">
    <p>{{ note.title }}</p>
    <p class="whitespace-pre-line">{{ note.description ?? "empty" }}</p>

    <UButton @click.prevent="deleteNote" type="button">Delete</UButton>
  </UCard>
</template>
