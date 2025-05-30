<script setup lang="ts">
import type { FormSubmitEvent } from "@nuxt/ui";
import { LoginSchema, type LoginType } from "~/types";

definePageMeta({
  layout: "auth",
});

const state = reactive({
  email: "",
  password: "",
});

async function onSubmit(event: FormSubmitEvent<LoginType>) {
  console.info(event.data);
}
</script>

<template>
  <div class="flex flex-col gap-y-8">
    <h1 class="text-4xl">Login</h1>
    <UForm
      :schema="LoginSchema"
      :state
      @submit="onSubmit"
      class="flex flex-col gap-y-5"
    >
      <UFormField label="Email" required name="email">
        <UInput
          type="email"
          v-model="state.email"
          required
          :ui="{
            root: 'flex',
          }"
        />
      </UFormField>

      <UFormField label="Password" required name="password">
        <UInput
          type="password"
          v-model="state.password"
          required
          :ui="{
            root: 'flex',
          }"
        />
      </UFormField>

      <UButton type="submit" class="self-start">Login</UButton>
    </UForm>
    <p class="text-sm">
      Doesn't have an account?
      <UButton to="/register" variant="ghost" color="neutral"
        >Click here!</UButton
      >
    </p>
  </div>
</template>
