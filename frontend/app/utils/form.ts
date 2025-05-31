import type { FormError } from "@nuxt/ui";

export const parseError = (errorData?: any | null): FormError[] => {
  const errors: FormError[] = [];

  Object.keys(errorData).forEach((key: string) => {
    const err: string[] = errorData[key];
    err.forEach((msg: string) => {
      errors.push({
        name: key,
        message: msg,
      });
    });
  });

  return errors;
};
