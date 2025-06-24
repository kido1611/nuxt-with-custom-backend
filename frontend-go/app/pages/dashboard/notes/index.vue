<script setup lang="ts">
import { useApiFetch } from "~/compossables/useApiFetch";
import type { ApiCollectionResponse, NoteResponse } from "~/types";

definePageMeta({
  layout: "dashboard",
  middleware: ["laravel-auth"],
});

const { data, status, error, refresh } =
  await useApiFetch<ApiCollectionResponse<NoteResponse>>("/api/api/notes");

async function refreshNotes() {
  await refresh();
}
</script>

<template>
  <div>
    <h1>Notes</h1>

    <UButton to="/dashboard/notes/create">Create Note</UButton>

    <div
      v-if="data"
      class="grid grid-cols-[repeat(auto-fill,minmax(300px,1fr))] gap-5"
    >
      <Note
        v-for="note in data.data"
        :key="note.id"
        :note
        @ask-reload="refreshNotes"
      />
    </div>
  </div>
</template>
