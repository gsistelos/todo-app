"use client";

import styles from "./LoginForm.module.css";
import { useRouter } from "next/navigation";
import { useState } from "react";
import FormGroup from "../FormGroup";

type FormData = {
  email: string;
  password: string;
};

const Login = () => {
  const router = useRouter();

  const [formData, setFormData] = useState<FormData>({
    email: "",
    password: "",
  });

  const [formError, setFormError] = useState<FormData>({
    email: "",
    password: "",
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const postData: FormData = {
      email: formData.email,
      password: formData.password,
    };

    try {
      const response = await fetch("http://localhost:8080/api/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(postData),
      });

      const data = await response.json();

      if (response.status !== 200) {
        setFormError(data);
        return;
      }

      localStorage.setItem("token", data.token);
      router.push("/");
    } catch (error) {
      alert(error);
    }
  };

  return (
    <form className={styles.form} onSubmit={handleSubmit}>
      <FormGroup
        label="Email"
        type="email"
        name="email"
        value={formData.email}
        onChange={handleChange}
        error={formError.email}
      />
      <FormGroup
        label="Password"
        type="password"
        name="password"
        value={formData.password}
        onChange={handleChange}
        error={formError.password}
      />
      <button className={styles.button} type="submit">
        Login
      </button>
    </form>
  );
};

export default Login;
