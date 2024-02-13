"use client";

import styles from "./page.module.css";
import { useRouter } from "next/navigation";

const Menu: React.FC = () => {
    const router = useRouter();

    return (
        <main className={styles.main}>
            <h1 className={styles.title}>Menu</h1>
            <p className={styles.paragraph}>This is the Menu page</p>
            <div className={styles.container}>
                <button className={styles.button} onClick={() => router.push("/register")}>
                    Register
                </button>
                <button className={styles.button} onClick={() => router.push("/login")}>
                    Login
                </button>
            </div>
        </main>
    );
}

export default Menu;
