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

export const NoteSchema = v.object({
  title: v.pipe(v.string(), v.minLength(1), v.maxLength(100)),
  description: v.pipe(v.nullish(v.string(), ""), v.maxLength(1000)),
});
export type NoteType = v.InferOutput<typeof NoteSchema>;

export type ApiCollectionResponse<T> = {
  data: T[];
};

export type NoteResponse = {
  id: number;
  title: string;
  description?: string | null;
  is_public: boolean;
  visible_at: string;
  created_at: string;
  updated_at: string;
};

export type UserResponse = {
  id: number;
  name: string;
  email: string;
  email_verified_at: string | null;
  created_at: string | null;
  updated_at: string | null;
};
