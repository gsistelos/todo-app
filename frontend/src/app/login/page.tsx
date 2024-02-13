"use client";

import { useState, ChangeEvent, FormEvent } from "react";
import { useRouter } from "next/navigation";
import styles from "./page.module.css";

const Login: React.FC = () => {
    const router = useRouter();

    const [formData, setFormData] = useState({
        email: "",
        password: "",
    });

    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setFormData((prevData) => ({
            ...prevData,
            [name]: value,
        }));
    };

    const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        const postData = {
            email: formData.email,
            password: formData.password,
        };

        try {
            const response = await fetch("http://localhost:8080/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(postData),
            })

            if (response.status !== 200) {
                const data = await response.json();
                throw new Error(data.error);
            } else {
                router.push("/home");
            }
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
                </div>
                <div className={styles.formGroup}>
                    <label>Password</label>
                    <input
                        type="password"
                        name="password"
                        value={formData.password}
                        onChange={handleChange}
                    />
                </div>
                <button className={styles.button} type="submit">Login</button>
            </form>
        </main>
    );
}

export default Login;
