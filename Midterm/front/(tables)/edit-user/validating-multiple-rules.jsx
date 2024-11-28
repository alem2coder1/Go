import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { useEffect } from 'react';
import * as z from "zod";
import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";


const FormSchema = z.object({
    username: z.string().min(5, {
        message: "Username must be at least 5 characters.",
    }),
    email: z.string().email({
        message: "Email is not valid.",
    }),
    password: z
        .string()
        .min(6, {
            message: "Password must be at least 6 characters.",
        })
        .refine((password) => /[A-Z]/.test(password), {
            message: "Password must contain at least one uppercase letter",
        })
        .refine((password) => /[a-z]/.test(password), {
            message: "Password must contain at least one lowercase letter",
        })
        .refine((password) => /[0-9]/.test(password), {
            message: "Password must contain at least one number",
        })
        .refine((password) => /[^A-Za-z0-9]/.test(password), {
            message: "Password must contain at least one special character",
        }),
});

const ValidatingEditRules = () => {

    const form = useForm({
        resolver: zodResolver(FormSchema),
        defaultValues: {
            name: "Alem",
            surname: "Kakan",
            email: "alemkaken@gmail.com",
            password: "",
            role: "editor",
            job: "Back-end Developer",
            birthday: "2003-03-01"
        },
    });


    return (
        <Form {...form}>
            <form >
                <div className="flex flex-col gap-4">
                    <FormField
                        control={form.control}
                        name="name"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>name</FormLabel>
                                <FormControl>
                                    <Input

                                        {...field}
                                        className={cn("", {
                                            "border-destructive focus:border-destructive":
                                            form.formState.errors.name,
                                        })}
                                    />
                                </FormControl>
                                <FormMessage className="bg-destructive/90 text-primary-foreground text-[10px] inline-flex justify-center items-center font-base h-[22px] px-2 rounded-sm " />
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="surname"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Surname</FormLabel>
                                <FormControl>
                                    <Input

                                        {...field}
                                        className={cn("", {
                                            "border-destructive focus:border-destructive":
                                            form.formState.errors.surname,
                                        })}
                                    />
                                </FormControl>
                                <FormMessage className="bg-destructive/90 text-primary-foreground text-[10px] inline-flex justify-center items-center font-base h-[22px] px-2 rounded-sm " />
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="email"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Email</FormLabel>
                                <FormControl>
                                    <Input
                                        placeholder="Enter your email"
                                        {...field}
                                        className={cn("", {
                                            "border-destructive focus:border-destructive":
                                            form.formState.errors.email,
                                        })}
                                    />
                                </FormControl>
                                <FormMessage className="bg-destructive/90 text-primary-foreground text-[10px] inline-flex justify-center items-center font-base h-[22px] px-2 rounded-sm " />
                            </FormItem>
                        )}
                    />

                    <FormField
                        control={form.control}
                        name="password"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Password</FormLabel>
                                <FormControl>
                                    <Input

                                        {...field}
                                        className={cn("", {
                                            "border-destructive focus:border-destructive":
                                            form.formState.errors.password,
                                        })}
                                    />
                                </FormControl>
                                <FormMessage className="bg-destructive/90 text-primary-foreground text-[10px] inline-flex justify-center items-center font-base h-[22px] px-2 rounded-sm " />
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="role"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Role</FormLabel>
                                <FormControl>
                                    <Input

                                        {...field}
                                        className={cn("", {
                                            "border-destructive focus:border-destructive":
                                            form.formState.errors.role,
                                        })}
                                    />
                                </FormControl>
                                <FormMessage className="bg-destructive/90 text-primary-foreground text-[10px] inline-flex justify-center items-center font-base h-[22px] px-2 rounded-sm " />
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="job"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Job</FormLabel>
                                <FormControl>
                                    <Input
                                        placeholder="Enter Job"
                                        {...field}
                                        className={cn("", {
                                            "border-destructive focus:border-destructive":
                                            form.formState.errors.job,
                                        })}
                                    />
                                </FormControl>
                                <FormMessage className="bg-destructive/90 text-primary-foreground text-[10px] inline-flex justify-center items-center font-base h-[22px] px-2 rounded-sm " />
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="birthday"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Birthday</FormLabel>
                                <FormControl>
                                    <Input
                                        placeholder="Enter Birthday"
                                        {...field}
                                        className={cn("", {
                                            "border-destructive focus:border-destructive":
                                            form.formState.errors.birthday,
                                        })}
                                    />
                                </FormControl>
                                <FormMessage className="bg-destructive/90 text-primary-foreground text-[10px] inline-flex justify-center items-center font-base h-[22px] px-2 rounded-sm " />
                            </FormItem>
                        )}
                    />
                </div>
                <div className="mt-4 flex ltr:justify-end rtl:justify-start">
                    <Button type="submit">Submit</Button>
                </div>
            </form>
        </Form>
    );
};

export default ValidatingEditRules;
