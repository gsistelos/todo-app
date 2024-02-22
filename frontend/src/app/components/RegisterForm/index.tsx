"use client";

import styles from "./RegisterForm.module.css";
import { useRouter } from "next/navigation";
import { useState } from "react";

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
      <div className={styles.formGroup}>
        <label>Username</label>
        <input
          type="text"
          name="username"
          value={formData.username}
          onChange={handleChange}
        />
        {formError.username && (
          <p className={styles.error}>{formError.username}</p>
        )}
      </div>
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
      <div className={styles.formGroup}>
        <label>Confirm password</label>
        <input
          type="password"
          value={confirmPassword}
          onChange={(e) => setConfirmPassword(e.target.value)}
        />
        {confirmPasswordError && (
          <p className={styles.error}>{confirmPasswordError}</p>
        )}
      </div>
      <button className={styles.button} type="submit">
        Register
      </button>
    </form>
  );
};

export default Register;
