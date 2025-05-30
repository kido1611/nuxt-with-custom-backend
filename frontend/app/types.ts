import * as v from "valibot";

export const LoginSchema = v.object({
  email: v.pipe(v.string(), v.trim(), v.email(), v.maxLength(100)),
  password: v.pipe(v.string(), v.trim(), v.minLength(8), v.maxLength(32)),
});

export const RegisterSchema = v.intersect([
  LoginSchema,
  v.object({
    name: v.pipe(v.string(), v.trim(), v.maxLength(100)),
  }),
]);

export type LoginType = v.InferOutput<typeof LoginSchema>;
export type RegisterType = v.InferOutput<typeof RegisterSchema>;
