"use client";

import styles from "./RegisterForm.module.css";
import { useRouter } from "next/navigation";
import { useState } from "react";
import FormGroup from "../FormGroup";

type FormData = {
  username: string;
  email: string;
  password: string;
};

const Register = () => {
  const router = useRouter();

  const [formData, setFormData] = useState<FormData>({
    username: "",
    email: "",
    password: "",
  });

  const [formError, setFormError] = useState<FormData>({
    username: "",
    email: "",
    password: "",
  });

  const [confirmPassword, setConfirmPassword] = useState<string>("");
  const [confirmPasswordError, setConfirmPasswordError] = useState<string>("");

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (formData.password !== confirmPassword) {
      setFormError({ username: "", email: "", password: "" });
      setConfirmPasswordError("Passwords do not match");
      return;
    }

    const postData: FormData = {
      username: formData.username,
      email: formData.email,
      password: formData.password,
    };

    try {
      const response = await fetch("http://localhost:8080/api/users", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(postData),
      });

      if (response.status !== 201) {
        const data = await response.json();
        setFormError(data);
        return;
      }

      alert("User created");
      router.push("/login");
    } catch (error) {
      alert(error);
    }
  };

  return (
    <form className={styles.form} onSubmit={handleSubmit}>
      <FormGroup
        label="Username"
        type="text"
        name="username"
        value={formData.username}
        onChange={handleChange}
        error={formError.username}
      />
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
      <FormGroup
        label="Confirm Password"
        type="password"
        name="confirmPassword"
        value={confirmPassword}
        onChange={(e) => setConfirmPassword(e.target.value)}
        error={confirmPasswordError}
      />
      <button className={styles.button} type="submit">
        Register
      </button>
    </form>
  );
};

export default Register;
