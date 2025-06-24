<script setup lang="ts">
import type { FormError, FormSubmitEvent } from "@nuxt/ui";
import type { NoteType } from "~/types";
import { NoteSchema } from "~/types";

definePageMeta({
  layout: "dashboard",
  middleware: ["laravel-auth"],
});

const toast = useToast();
const client = useNuxtApp().$apiFetch;
const form = useTemplateRef("form");

const state = reactive({
  title: "",
  description: "",
});

async function onSubmit(event: FormSubmitEvent<NoteType>) {
  form.value?.clear();
  try {
    await client("/api/api/notes", {
      method: "POST",
      body: event.data,
    });

    await navigateTo("/dashboard/notes");

    toast.add({
      title: "Success",
      color: "success",
    });
    // TODO: fix type
  } catch (error: any) {
    toast.add({
      title: "Error",
      description: error.data.message,
      color: "error",
    });

    const errors: FormError[] = parseError(error.data.errors);

    form.value?.setErrors(errors);
  }
}
</script>

<template>
  <div>
    <h1>Create Notes</h1>

    <UForm
      ref="form"
      :schema="NoteSchema"
      :state
      @submit="onSubmit"
      class="flex flex-col gap-y-5"
    >
      <UFormField label="title" required name="title">
        <UInput
          type="text"
          v-model="state.title"
          required
          :ui="{
            root: 'flex',
          }"
        />
      </UFormField>
      <UFormField label="description" name="Description">
        <UTextarea
          v-model="state.description"
          :rows="4"
          :ui="{
            root: 'flex',
          }"
        />
      </UFormField>

      <UButton type="submit" class="self-start">Add</UButton>
    </UForm>
  </div>
</template>
