<script setup lang="ts">
import type { FormSubmitEvent, FormError } from "@nuxt/ui";
import { useAuth } from "~/compossables/useAuth";
import { LoginSchema, type LoginType } from "~/types";

definePageMeta({
  layout: "auth",
  middleware: ["laravel-guest"],
});

const toast = useToast();
const form = useTemplateRef("form");

const { login } = useAuth();

const state = reactive({
  email: "ahmad@local.host",
  password: "password",
});

async function onSubmit(event: FormSubmitEvent<LoginType>) {
  form.value?.clear();

  try {
    await login(event.data);

    await navigateTo("/dashboard");
    // TODO: fix type
  } catch (error: any) {
    console.log(error, error.data);
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
  <div class="flex flex-col gap-y-8">
    <h1 class="text-4xl">Login</h1>
    <UForm
      ref="form"
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
