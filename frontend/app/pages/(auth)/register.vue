<script setup lang="ts">
import type { FormSubmitEvent } from "@nuxt/ui";
import { RegisterSchema, type RegisterType } from "~/types";

definePageMeta({
  layout: "auth",
  middleware: ["laravel-guest"],
});

const client = useNuxtApp().$apiFetch;
const toast = useToast();
const form = useTemplateRef("form");

const state = reactive({
  name: "ahmad",
  email: "ahmad@local.host",
  password: "password",
});

async function onSubmit(event: FormSubmitEvent<RegisterType>) {
  form.value?.clear();

  try {
    await client("/api/api/auth/register", {
      method: "POST",
      body: event.data,
    });

    await navigateTo("/login");

    toast.add({
      title: "Success",
      color: "success",
    });
    // TODO: fix type
  } catch (error: any) {
    toast.add({
      title: "Error",
      color: "error",
    });

    // const errors: FormError[] = parseError(error.data.errors);
    //
    // form.value?.setErrors(errors);
  }
}
</script>

<template>
  <div class="flex flex-col gap-y-8">
    <h1 class="text-4xl">Register</h1>

    <UForm
      ref="form"
      :schema="RegisterSchema"
      :state
      @submit="onSubmit"
      class="flex flex-col gap-y-5"
    >
      <UFormField label="Name" required name="name">
        <UInput
          type="text"
          v-model="state.name"
          required
          :ui="{
            root: 'flex',
          }"
        />
      </UFormField>

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

      <UButton type="submit" class="self-start">Register</UButton>
    </UForm>

    <p class="text-sm">
      Already have an account?
      <UButton to="/login" variant="ghost" color="neutral">click here!</UButton>
    </p>
  </div>
</template>
