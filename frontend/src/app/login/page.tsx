"use client";

import styles from "./page.module.css";
import { useRouter } from "next/navigation";
import { useState, ChangeEvent, FormEvent } from "react";

interface FormData {
  email: string;
  password: string;
}

const Login: React.FC = () => {
  const router = useRouter();

  const [formData, setFormData] = useState<FormData>({
    email: "",
    password: "",
  });

  const [formError, setFormError] = useState<FormData>({
    email: "",
    password: "",
  });

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
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

      if (response.status !== 200) {
        const data = await response.json();
        setFormError(data);
        return;
      }

      router.push("/home");
    } catch (error) {
      alert(error);
    }
  };

  return (
    <main className={styles.main}>
      <h1 className={styles.title}>Login</h1>
      <p className={styles.paragraph}>This is the Login page</p>
      <form className={styles.form} onSubmit={handleSubmit}>
        <div className={styles.formGroup}>
          <label>Email</label>
          <input
            type="email"
            name="email"
            value={formData.email}
            onChange={handleChange}
          />
          {formError.email && <p className={styles.error}>{formError.email}</p>}
        </div>
        <div className={styles.formGroup}>
          <label>Password</label>
          <input
            type="password"
            name="password"
            value={formData.password}
            onChange={handleChange}
          />
          {formError.password && (
            <p className={styles.error}>{formError.password}</p>
          )}
        </div>
        <button className={styles.button} type="submit">
          Login
        </button>
      </form>
    </main>
  );
};

export default Login;
